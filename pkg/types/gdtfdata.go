package MVRTypes

import (
	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
	GDTFTypes "github.com/Patch2PDF/GDTF-Parser/pkg/types"
)

type GDTF struct {
	Data   *GDTFTypes.GDTF
	Meshes map[string]*MeshTypes.Mesh
}
