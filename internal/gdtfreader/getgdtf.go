package GDTFReader

import (
	"archive/zip"
	"strings"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
	GDTFParser "github.com/Patch2PDF/GDTF-Parser"
	MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"
	"golang.org/x/sync/errgroup"
)

type GDTFTask struct {
	GDTFSpec  string
	GDTFModes map[string]struct{}
}

var gdtfTaskMap = map[string]*GDTFTask{}

func AddToTaskMap(gdtfSpec string, gdtfMode string) {
	if gdtfSpec == "" {
		return
	}
	if gdtfTaskMap[gdtfSpec] == nil {
		gdtfTaskMap[gdtfSpec] = &GDTFTask{
			GDTFSpec:  gdtfSpec,
			GDTFModes: map[string]struct{}{gdtfMode: {}},
		}
	} else if _, found := gdtfTaskMap[gdtfSpec].GDTFModes[gdtfMode]; found {
		gdtfTaskMap[gdtfSpec].GDTFModes[gdtfMode] = struct{}{}
	}
}

func getGDTF(task *GDTFTask, fileMap map[string]*zip.File, config MVRTypes.MVRParserConfig) error {
	gdtfFile := task.GDTFSpec
	if !strings.HasSuffix(gdtfFile, ".gdtf") {
		gdtfFile = gdtfFile + ".gdtf"
	}
	file, err := fileMap[gdtfFile].Open()
	if err != nil {
		return err
	}
	gdtf, err := GDTFParser.ParseGDTFByFile(file, config.MeshHandling >= MVRTypes.ReadMeshesIntoModels, config.ReadThumbnail)
	if err != nil {
		return err
	}
	meshes := map[string]*MeshTypes.Mesh{}
	for gdtfMode := range task.GDTFModes {
		if config.MeshHandling >= MVRTypes.BuildFixtureModels {
			mesh, err := gdtf.BuildMesh(gdtfMode)
			if err != nil {
				return err
			}
			meshes[gdtfMode] = mesh
		}
	}
	MVRTypes.AddGDTFPointer(task.GDTFSpec, &MVRTypes.GDTF{
		Data:   gdtf,
		Meshes: meshes,
	})
	return nil
}

func GetGDTFs(fileMap map[string]*zip.File, config MVRTypes.MVRParserConfig) error {
	eg := errgroup.Group{}

	var numWorkers = config.GDTFParserWorkers
	jobs := make(chan *GDTFTask, len(gdtfTaskMap))

	for i := 0; i < numWorkers; i++ {
		eg.Go(func() error {
			for j := range jobs {
				err := getGDTF(j, fileMap, config)
				if err != nil {
					return err
				}
			}
			return nil
		})
	}

	for _, t := range gdtfTaskMap {
		jobs <- t
	}
	close(jobs)

	return eg.Wait()
}

// func GetGDTFs(fileMap map[string]*zip.File, modelLevel int, readThumbnail bool) error {
// 	eg := errgroup.Group{}
// 	// TODO: switch to worker pool
// 	for _, task := range gdtfTaskMap {
// 		eg.Go(func() error {
// 			return func(task *GDTFTask) error {
// 				gdtfFile := task.GDTFSpec
// 				if !strings.HasSuffix(gdtfFile, ".gdtf") {
// 					gdtfFile = gdtfFile + ".gdtf"
// 				}
// 				file, err := fileMap[gdtfFile].Open()
// 				if err != nil {
// 					return err
// 				}
// 				gdtf, err := GDTFParser.ParseGDTFByFile(file, modelLevel >= MVRTypes.ReadMeshesIntoModels, readThumbnail)
// 				if err != nil {
// 					return err
// 				}
// 				meshes := map[string]*MeshTypes.Mesh{}
// 				for gdtfMode := range task.GDTFModes {
// 					if modelLevel >= MVRTypes.BuildFixtureModels {
// 						mesh, err := gdtf.BuildMesh(gdtfMode)
// 						if err != nil {
// 							return err
// 						}
// 						meshes[gdtfMode] = mesh
// 					}
// 				}
// 				MVRTypes.AddGDTFPointer(task.GDTFSpec, &MVRTypes.GDTF{
// 					Data:   gdtf,
// 					Meshes: meshes,
// 				})
// 				return nil
// 			}(task)
// 		})
// 	}
// 	return eg.Wait()
// }
