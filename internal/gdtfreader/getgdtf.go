package GDTFReader

import (
	"archive/zip"
	"strings"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
	GDTFParser "github.com/Patch2PDF/GDTF-Parser"
	MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"
)

func GetGDTF(fileMap map[string]*zip.File, gdtfSpec string, gdtfMode string, modelLevel int, readThumbnail bool) error {
	if gdtfSpec == "" {
		return nil
	}
	gdtfFile := gdtfSpec
	if !strings.HasSuffix(gdtfSpec, ".gdtf") {
		gdtfFile = gdtfFile + ".gdtf"
	}
	ptr := MVRTypes.GetGDTFPointer(gdtfSpec)
	if ptr == nil {
		file, err := fileMap[gdtfFile].Open()
		if err != nil {
			return err
		}
		gdtf, err := GDTFParser.ParseGDTFByFile(file, modelLevel >= MVRTypes.ReadMeshesIntoModels, readThumbnail)
		if err != nil {
			return err
		}
		meshes := map[string]*MeshTypes.Mesh{}
		if modelLevel >= MVRTypes.BuildFixtureModels {
			mesh, err := gdtf.BuildMesh(gdtfMode)
			if err != nil {
				return err
			}
			meshes[gdtfMode] = mesh
		}
		MVRTypes.AddGDTFPointer(gdtfSpec, &MVRTypes.GDTF{
			Data:   gdtf,
			Meshes: meshes,
		})
	} else if modelLevel >= MVRTypes.BuildFixtureModels && ptr.Meshes[gdtfMode] == nil {
		mesh, err := ptr.Data.BuildMesh(gdtfMode)
		if err != nil {
			return err
		}
		ptr.Meshes[gdtfMode] = mesh
	}
	return nil
}
