package MVRTypes

import (
	"archive/zip"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
)

type GroupObject struct {
	UUID   string
	Name   string
	Matrix MeshTypes.Matrix
	Class  NodeReference[Class]
	ChildList
}

func (a *GroupObject) CreateReferencePointer() {
	a.ChildList.CreateReferencePointer()
}

func (a *GroupObject) ResolveReference() {
	if a.Class.String != nil {
		a.Class.Ptr = refPointers.Classes[*a.Class.String]
	}
	a.ChildList.ResolveReference()
}

func (a *GroupObject) ReadMesh(fileMap map[string]*zip.File) error {
	return a.ChildList.ReadMesh(fileMap)
}
