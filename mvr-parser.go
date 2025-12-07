package MVRParser

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"

	GDTFReader "github.com/Patch2PDF/MVR-Parser/internal/gdtfreader"
	"github.com/Patch2PDF/MVR-Parser/internal/geometry3d"
	MVRXML "github.com/Patch2PDF/MVR-Parser/internal/types/mvrxml"
	MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"
)

func ParseMVRZipReader(zipfile *zip.Reader, config MVRTypes.MVRParserConfig) (*MVRTypes.GeneralSceneDescription, error) {
	var mvrData MVRXML.GeneralSceneDescription

	// TODO: docs + tests + readme

	// put all files in zip into the filemap
	var fileMap map[string]*zip.File = make(map[string]*zip.File)
	for _, file := range zipfile.File {
		fileMap[file.Name] = file
	}

	// parse mvr xml
	xmlFile, err := fileMap["GeneralSceneDescription.xml"].Open()
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(data, &mvrData)
	if err != nil {
		return nil, err
	}

	parseConfig := MVRXML.ParseConfigData{
		GDTFTaskMap: &map[string]*GDTFReader.GDTFTask{},
	}

	parsedData := mvrData.Parse(parseConfig)

	refPointers := MVRTypes.CreateRefPointersMap()

	GDTFReader.GetGDTFs(parseConfig.GDTFTaskMap, refPointers, fileMap, config)

	parsedData.CreateReferencePointer(refPointers)

	parsedData.ResolveReference(refPointers)

	if config.MeshHandling >= MVRTypes.ReadMeshesIntoModels {
		geometry3d.ReadMeshes(fileMap, parsedData)
	}

	if config.MeshHandling >= MVRTypes.BuildStageModel {
		meshTasks := MVRTypes.MeshTasks{}
		parsedData.CreateMeshTasks(&meshTasks, config.ModelConfig)
		fmt.Println(len(meshTasks))

		parsedData.StageModel = MVRTypes.CompleteMeshTasks(&meshTasks, config)
	}

	return parsedData, nil
}
