package MVRTypes

type ChildList struct {
	SceneObjects []*SceneObject
	GroupObjects []*SceneObject
	FocusPoints  []*FocusPoint
	Fixtures     []*Fixture
	Supports     []*Support
	Trusses      []*Truss
	VideoScreens []*VideoScreen
	Projectors   []*Projector
}
