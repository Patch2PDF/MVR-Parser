package MVRXML

import MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"

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

func (a *ChildList) Parse() MVRTypes.ChildList {
	return MVRTypes.ChildList{
		SceneObjects: ParseList(&a.SceneObjects),
		GroupObjects: ParseList(&a.GroupObjects),
		FocusPoints:  ParseList(&a.FocusPoints),
		Fixtures:     ParseList(&a.Fixtures),
		Supports:     ParseList(&a.Supports),
		Trusses:      ParseList(&a.Trusses),
		VideoScreens: ParseList(&a.VideoScreens),
		Projectors:   ParseList(&a.Projectors),
	}
}
