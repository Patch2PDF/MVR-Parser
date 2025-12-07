package MVRTypes

import "archive/zip"

type ChildList struct {
	SceneObjects []*SceneObject
	GroupObjects []*GroupObject
	FocusPoints  []*FocusPoint
	Fixtures     []*Fixture
	Supports     []*Support
	Trusses      []*Truss
	VideoScreens []*VideoScreen
	Projectors   []*Projector
}

func (c *ChildList) CreateReferencePointer(refPointers *ReferencePointers) {
	CreateReferencePointers(refPointers, &c.SceneObjects)
	CreateReferencePointers(refPointers, &c.GroupObjects)
	CreateReferencePointers(refPointers, &c.FocusPoints)
	CreateReferencePointers(refPointers, &c.Fixtures)
	CreateReferencePointers(refPointers, &c.Supports)
	CreateReferencePointers(refPointers, &c.Trusses)
	CreateReferencePointers(refPointers, &c.VideoScreens)
	CreateReferencePointers(refPointers, &c.Projectors)
}

func (c *ChildList) ResolveReference(refPointers *ReferencePointers) {
	ResolveReferences(refPointers, &c.SceneObjects)
	ResolveReferences(refPointers, &c.GroupObjects)
	ResolveReferences(refPointers, &c.FocusPoints)
	ResolveReferences(refPointers, &c.Fixtures)
	ResolveReferences(refPointers, &c.Supports)
	ResolveReferences(refPointers, &c.Trusses)
	ResolveReferences(refPointers, &c.VideoScreens)
	ResolveReferences(refPointers, &c.Projectors)
}

func (c *ChildList) ReadMesh(fileMap map[string]*zip.File) error {
	err := ReadMeshes(c.SceneObjects, fileMap)
	if err != nil {
		return err
	}
	err = ReadMeshes(c.GroupObjects, fileMap)
	if err != nil {
		return err
	}
	err = ReadMeshes(c.FocusPoints, fileMap)
	if err != nil {
		return err
	}
	err = ReadMeshes(c.Fixtures, fileMap)
	if err != nil {
		return err
	}
	err = ReadMeshes(c.Supports, fileMap)
	if err != nil {
		return err
	}
	err = ReadMeshes(c.Trusses, fileMap)
	if err != nil {
		return err
	}
	err = ReadMeshes(c.VideoScreens, fileMap)
	if err != nil {
		return err
	}
	err = ReadMeshes(c.Projectors, fileMap)
	if err != nil {
		return err
	}
	return nil
}

type MeshReader interface {
	ReadMesh(fileMap map[string]*zip.File) error
}

func ReadMeshes[T MeshReader](src []T, fileMap map[string]*zip.File) error {
	for _, element := range src {
		err := element.ReadMesh(fileMap)
		if err != nil {
			return err
		}
	}
	return nil
}
