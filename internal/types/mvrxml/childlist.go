package MVRXML

import MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"

type ChildList struct {
	SceneObjects []*SceneObject `xml:"ChildList>SceneObject,omitempty"`
	GroupObjects []*GroupObject `xml:"ChildList>GroupObject,omitempty"`
	FocusPoints  []*FocusPoint  `xml:"ChildList>FocusPoint,omitempty"`
	Fixtures     []*Fixture     `xml:"ChildList>Fixture,omitempty"`
	Supports     []*Support     `xml:"ChildList>Support,omitempty"`
	Trusses      []*Truss       `xml:"ChildList>Truss,omitempty"`
	VideoScreens []*VideoScreen `xml:"ChildList>VideoScreen,omitempty"`
	Projectors   []*Projector   `xml:"ChildList>Projector,omitempty"`
}

func (a *ChildList) Parse(config ParseConfigData) MVRTypes.ChildList {
	return MVRTypes.ChildList{
		SceneObjects: ParseList(config, &a.SceneObjects),
		GroupObjects: ParseList(config, &a.GroupObjects),
		FocusPoints:  ParseList(config, &a.FocusPoints),
		Fixtures:     ParseList(config, &a.Fixtures),
		Supports:     ParseList(config, &a.Supports),
		Trusses:      ParseList(config, &a.Trusses),
		VideoScreens: ParseList(config, &a.VideoScreens),
		Projectors:   ParseList(config, &a.Projectors),
	}
}
