package MVRXML

type FocusPoint struct {
	UUID       string     `xml:"uuid,attr"`
	Name       string     `xml:"name,attr"`
	Matrix     *Matrix    `xml:"Matrix,omitempty"`
	Class      *string    `xml:"Classing,omitempty"`
	Geometries Geometries `xml:"Geometries"`
}
