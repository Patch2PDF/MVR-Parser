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

func AddToTaskMap(gdtfTaskMap *map[string]*GDTFTask, gdtfSpec string, gdtfMode string) {
	if gdtfSpec == "" {
		return
	}
	if (*gdtfTaskMap)[gdtfSpec] == nil {
		(*gdtfTaskMap)[gdtfSpec] = &GDTFTask{
			GDTFSpec:  gdtfSpec,
			GDTFModes: map[string]struct{}{gdtfMode: {}},
		}
	} else if _, found := (*gdtfTaskMap)[gdtfSpec].GDTFModes[gdtfMode]; found {
		(*gdtfTaskMap)[gdtfSpec].GDTFModes[gdtfMode] = struct{}{}
	}
}

func getGDTF(jobs <-chan *GDTFTask, results chan<- *MVRTypes.GDTF, fileMap map[string]*zip.File, config MVRTypes.MVRParserConfig) error {
	for task := range jobs {
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
		results <- &MVRTypes.GDTF{
			Name:   task.GDTFSpec,
			Data:   gdtf,
			Meshes: meshes,
		}
	}
	return nil
}

func GetGDTFs(gdtfTaskMap *map[string]*GDTFTask, fileMap map[string]*zip.File, config MVRTypes.MVRParserConfig) error {
	eg := errgroup.Group{}

	var numWorkers = config.GDTFParserWorkers
	jobs := make(chan *GDTFTask, len(*gdtfTaskMap))
	results := make(chan *MVRTypes.GDTF, len(*gdtfTaskMap))

	for range numWorkers {
		eg.Go(func() error {
			return getGDTF(jobs, results, fileMap, config)
		})
	}

	for _, t := range *gdtfTaskMap {
		jobs <- t
	}
	close(jobs)

	err := eg.Wait()
	if err != nil {
		return err
	}

	close(results)

	for gdtf := range results {
		MVRTypes.AddGDTFPointer(gdtf.Name, gdtf)
	}

	return nil
}
