package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type Geometries struct {
	Geometry3D []*Geometry3D
	Symbol     []*Symbol
}

type Geometry3D struct {
	FileName fileName
	Matrix   MeshTypes.Matrix
}

type Symbol struct {
	UUID   string
	SymDef NodeReference[SymDef] // TODO: Node reference
	Matrix MeshTypes.Matrix
}
