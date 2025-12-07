package MVRXML

import MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"

type GroupObject struct {
	UUID   string  `xml:"uuid,attr"`
	Name   string  `xml:"name,attr"`
	Matrix *Matrix `xml:"Matrix"`
	Class  *string `xml:"Classing"`
	ChildList
}

func (a *GroupObject) Parse(config ParseConfigData) *MVRTypes.GroupObject {
	return &MVRTypes.GroupObject{
		UUID:      a.UUID,
		Name:      a.Name,
		Matrix:    a.Matrix.ToMeshMatrix(),
		Class:     MVRTypes.NodeReference[MVRTypes.Class]{String: a.Class},
		ChildList: a.ChildList.Parse(config),
	}
}
