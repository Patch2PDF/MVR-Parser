package GDTFReader

import (
	"archive/zip"
	"strings"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
	GDTFParser "github.com/Patch2PDF/GDTF-Parser"
	MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"
)

func GetGDTF(fileMap map[string]*zip.File, gdtfMap *map[string]*MVRTypes.GDTF, gdtfSpec string, gdtfMode string, modelLevel int, readThumbnail bool) error {
	if gdtfSpec == "" {
		return nil
	}
	if !strings.HasSuffix(gdtfSpec, ".gdtf") {
		gdtfSpec = gdtfSpec + ".gdtf"
	}
	if (*gdtfMap)[gdtfSpec] == nil {
		file, err := fileMap[gdtfSpec].Open()
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
		(*gdtfMap)[gdtfSpec] = &MVRTypes.GDTF{
			Data:   gdtf,
			Meshes: meshes,
		}
	} else if modelLevel >= MVRTypes.BuildFixtureModels && (*gdtfMap)[gdtfSpec].Meshes[gdtfMode] == nil {
		mesh, err := (*gdtfMap)[gdtfSpec].Data.BuildMesh(gdtfMode)
		if err != nil {
			return err
		}
		(*gdtfMap)[gdtfSpec].Meshes[gdtfMode] = mesh
	}
	return nil
}
