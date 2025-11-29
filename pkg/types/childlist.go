package MVRTypes

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
