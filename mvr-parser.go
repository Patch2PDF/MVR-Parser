package MVRParser

import (
	"archive/zip"
	"encoding/xml"
	"io"
	"os"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
	STL "github.com/Patch2PDF/GDTF-Parser/examples/stl"
	"github.com/Patch2PDF/MVR-Parser/internal/geometry3d"
	MVRXML "github.com/Patch2PDF/MVR-Parser/internal/types/mvrxml"
	MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"
)

func ParseMVRZipReader(zipfile *zip.Reader, meshHandling int, readThumbnail bool, modelConfig *MVRTypes.ModelConfig) (*MVRTypes.GeneralSceneDescription, error) {
	var mvrData MVRXML.GeneralSceneDescription

	// TODO: build stage model only if requested + allow toggle for only including addressed fixtures in model, possibly excluding groups
	// TODO: Leverage GoRoutines
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

	// read required gdtf files and generate models as desired
	for _, layer := range mvrData.Scene.Layers {
		err := layer.ChildList.GetGDTFs(fileMap, meshHandling, readThumbnail)
		if err != nil {
			return nil, err
		}
	}

	parsedData := mvrData.Parse()

	parsedData.CreateReferencePointer()

	parsedData.ResolveReference()

	if meshHandling >= MVRTypes.ReadMeshesIntoModels {
		geometry3d.ReadMeshes(fileMap, parsedData)
	}

	if meshHandling >= MVRTypes.BuildStageModel {
		mesh := parsedData.Scene.Layers[0].ChildList.GenerateMesh(MeshTypes.IdentityMatrix(), modelConfig)

		// write mesh as STL
		f, _ := os.Create("Test.stl")
		STL.WriteBinary(f, mesh)
		f.Close()
	}

	return parsedData, nil
}
