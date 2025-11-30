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

func (c *ChildList) CreateReferencePointer() {
	CreateReferencePointers(&c.SceneObjects)
	CreateReferencePointers(&c.GroupObjects)
	CreateReferencePointers(&c.FocusPoints)
	CreateReferencePointers(&c.Fixtures)
	CreateReferencePointers(&c.Supports)
	CreateReferencePointers(&c.Trusses)
	CreateReferencePointers(&c.VideoScreens)
	CreateReferencePointers(&c.Projectors)
}

func (c *ChildList) ResolveReference() {
	ResolveReferences(&c.SceneObjects)
	ResolveReferences(&c.GroupObjects)
	ResolveReferences(&c.FocusPoints)
	ResolveReferences(&c.Fixtures)
	ResolveReferences(&c.Supports)
	ResolveReferences(&c.Trusses)
	ResolveReferences(&c.VideoScreens)
	ResolveReferences(&c.Projectors)
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
