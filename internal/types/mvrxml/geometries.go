package MVRXML

type Geometries struct {
	Geometry3D []*Geometry3D `xml:"Geometry3D"`
	Symbol     []*Symbol     `xml:"Symbol"`
}

type Geometry3D struct {
	FileName fileName `xml:"fileName,attr"`
	Matrix   *Matrix  `xml:"Matrix,omitempty"`
}

type Symbol struct {
	UUID   string  `xml:"uuid,attr"`
	SymDef string  `xml:"symdef,attr"`
	Matrix *Matrix `xml:"Matrix,omitempty"`
}
