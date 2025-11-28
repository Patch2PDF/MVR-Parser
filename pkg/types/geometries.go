package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type Geometries struct {
	Geometry3D []*Geometry3D
	Symbol     []*Symbol
}

func (a *Geometries) ResolveReference() {
	ResolveReferences(&a.Symbol)
}

type Geometry3D struct {
	FileName fileName
	Matrix   MeshTypes.Matrix
}

type Symbol struct {
	UUID   string
	SymDef NodeReference[SymDef]
	Matrix MeshTypes.Matrix
}

func (a *Symbol) ResolveReference() {
	if a.SymDef.String != nil {
		a.SymDef.Ptr = refPointers.SymDefs[*a.SymDef.String]
	}
}
