package GDTFReader

import (
	"archive/zip"
	"strings"

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
		for gdtfMode := range task.GDTFModes {
			if config.MeshHandling >= MVRTypes.BuildFixtureModels && gdtf.FixtureType.DMXModes[gdtfMode].MeshModels == nil {
				_, err := gdtf.BuildMesh(gdtfMode)
				if err != nil {
					return err
				}
			}
		}
		results <- &MVRTypes.GDTF{
			Name: task.GDTFSpec,
			Data: gdtf,
		}
	}
	return nil
}

func GetGDTFs(gdtfTaskMap *map[string]*GDTFTask, refPointers *MVRTypes.ReferencePointers, fileMap map[string]*zip.File, config MVRTypes.MVRParserConfig) error {
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
		refPointers.GDTFSpecs[gdtf.Name] = gdtf
	}

	return nil
}
