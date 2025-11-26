package MVRParser

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"

	MVRXML "github.com/Patch2PDF/MVR-Parser/internal/types/mvrxml"
	MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"
)

func ParseMVRZipReader(zipfile *zip.Reader, meshHandling int, readThumbnail bool) (*MVRXML.GeneralSceneDescription, error) {
	var mvrData MVRXML.GeneralSceneDescription

	// TODO: create file map, extract MVR xml data, parse GDTFs with meshes and thumbnails in Go routines, build models in go routines

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
	var gdtfMap map[string]*MVRTypes.GDTF = make(map[string]*MVRTypes.GDTF)

	for _, layer := range mvrData.Scene.Layers {
		err := layer.ChildList.GetGDTFs(fileMap, &gdtfMap, meshHandling, readThumbnail)
		if err != nil {
			return nil, err
		}
	}

	fmt.Printf("%+v\n", gdtfMap)

	// for _, file := range zipfile.File {
	// 	if filepath.Ext(file.Name) == ".gdtf" {
	// 		fmt.Printf("%s\n", file.Name)
	// 		// gdtf, err := file.Open()
	// 		// if err != nil {
	// 		// 	return nil, err
	// 		// }
	// 		// data, err := GDTFParser.ParseGDTFByFile(gdtf, true, false)
	// 		// if err != nil {
	// 		// 	return nil, err
	// 		// }
	// 		// fmt.Printf("%+v\n", *data)
	// 		// keys := make([]string, 0, len(data.FixtureType.DMXModes))
	// 		// for k := range data.FixtureType.DMXModes {
	// 		// 	keys = append(keys, k)
	// 		// }
	// 		// mesh, err := data.BuildMesh(keys[0])
	// 		// if err != nil {
	// 		// 	return nil, err
	// 		// }
	// 		// // write mesh as STL
	// 		// f, _ := os.Create(file.Name + ".stl")
	// 		// STL.WriteBinary(f, mesh)
	// 		// f.Close()
	// 	} else if file.Name == "GeneralSceneDescription.xml" {
	// 		xmlFile, err := file.Open()
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		data, err := io.ReadAll(xmlFile)
	// 		if err != nil {
	// 			return nil, err
	// 		}

	// 		err = xml.Unmarshal(data, &mvrData)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 	}
	// }

	return &mvrData, nil
}
