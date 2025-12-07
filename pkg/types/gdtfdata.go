package MVRTypes

import (
	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
	GDTFTypes "github.com/Patch2PDF/GDTF-Parser/pkg/types"
)

type GDTF struct {
	Name   string
	Data   *GDTFTypes.GDTF
	Meshes map[string]*MeshTypes.Mesh
}
