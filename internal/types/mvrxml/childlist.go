package MVRXML

type ChildList struct {
	SceneObjects []*SceneObject `xml:"ChildList>SceneObject,omitempty"`
	GroupObjects []*SceneObject `xml:"ChildList>GroupObject,omitempty"`
	FocusPoints  []*FocusPoint  `xml:"ChildList>FocusPoint,omitempty"`
	Fixtures     []*Fixture     `xml:"ChildList>Fixture,omitempty"`
	Supports     []*Support     `xml:"ChildList>Support,omitempty"`
	Trusses      []*Truss       `xml:"ChildList>Truss,omitempty"`
	VideoScreens []*VideoScreen `xml:"ChildList>VideoScreen,omitempty"`
	Projectors   []*Projector   `xml:"ChildList>Projector,omitempty"`
}
