package MVRXML

import MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"

type Geometries struct {
	Geometry3D []*Geometry3D `xml:"Geometry3D"`
	Symbol     []*Symbol     `xml:"Symbol"`
}

func (a *Geometries) Parse() *MVRTypes.Geometries {
	return &MVRTypes.Geometries{
		Geometry3D: ParseList(&a.Geometry3D),
		Symbol:     ParseList(&a.Symbol),
	}
}

type Geometry3D struct {
	FileName fileName `xml:"fileName,attr"`
	Matrix   *Matrix  `xml:"Matrix,omitempty"`
}

func (a *Geometry3D) Parse() *MVRTypes.Geometry3D {
	return &MVRTypes.Geometry3D{
		FileName: a.FileName,
		Matrix:   a.Matrix.ToMeshMatrix(),
	}
}

type Symbol struct {
	UUID   string  `xml:"uuid,attr"`
	SymDef string  `xml:"symdef,attr"`
	Matrix *Matrix `xml:"Matrix,omitempty"`
}

func (a *Symbol) Parse() *MVRTypes.Symbol {
	return &MVRTypes.Symbol{
		UUID:   a.UUID,
		SymDef: MVRTypes.NodeReference[MVRTypes.SymDef]{String: &a.SymDef},
		Matrix: a.Matrix.ToMeshMatrix(),
	}
}
