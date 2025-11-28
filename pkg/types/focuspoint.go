package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type FocusPoint struct {
	UUID       string
	Name       string
	Matrix     MeshTypes.Matrix
	Class      NodeReference[Class] // TODO: Node reference
	Geometries *Geometries
}

func (a *FocusPoint) CreateReferencePointer() {
	refPointers.FoucsPoints[a.UUID] = a
}

func (a *FocusPoint) ResolveReference() {
	if a.Class.String != nil {
		a.Class.Ptr = refPointers.Classes[*a.Class.String]
	}
	a.Geometries.ResolveReference()
}
