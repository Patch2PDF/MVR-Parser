package MVRXML

import MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"

type FocusPoint struct {
	UUID       string     `xml:"uuid,attr"`
	Name       string     `xml:"name,attr"`
	Matrix     *Matrix    `xml:"Matrix,omitempty"`
	Class      *string    `xml:"Classing,omitempty"`
	Geometries Geometries `xml:"Geometries"`
}

func (a *FocusPoint) Parse() *MVRTypes.FocusPoint {
	return &MVRTypes.FocusPoint{
		UUID:       a.UUID,
		Name:       a.Name,
		Matrix:     a.Matrix.ToMeshMatrix(),
		Class:      MVRTypes.NodeReference[MVRTypes.Class]{String: a.Class},
		Geometries: a.Geometries.Parse(),
	}
}
