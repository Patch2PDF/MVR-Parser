package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type GroupObject struct {
	UUID   string
	Name   string
	Matrix MeshTypes.Matrix
	Class  NodeReference[Class] // TODO: Node reference
	ChildList
}
