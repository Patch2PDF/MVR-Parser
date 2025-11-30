package geometry3d

import (
	"archive/zip"

	MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"
)

func ReadMeshes(fileMap map[string]*zip.File, mvr *MVRTypes.GeneralSceneDescription) error {
	err := mvr.Scene.AuxData.ReadMesh(fileMap)
	if err != nil {
		return err
	}
	for _, layer := range mvr.Scene.Layers {
		err := layer.ChildList.ReadMesh(fileMap)
		if err != nil {
			return err
		}
	}
	return nil
}
