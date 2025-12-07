package MVRXML

import MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"

type Geometries struct {
	Geometry3D []*Geometry3D `xml:"Geometry3D"`
	Symbol     []*Symbol     `xml:"Symbol"`
}

func (a *Geometries) Parse(config ParseConfigData) *MVRTypes.Geometries {
	return &MVRTypes.Geometries{
		Geometry3D: ParseList(config, &a.Geometry3D),
		Symbol:     ParseList(config, &a.Symbol),
	}
}

type Geometry3D struct {
	FileName fileName `xml:"fileName,attr"`
	Matrix   *Matrix  `xml:"Matrix,omitempty"`
}

func (a *Geometry3D) Parse(config ParseConfigData) *MVRTypes.Geometry3D {
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

func (a *Symbol) Parse(config ParseConfigData) *MVRTypes.Symbol {
	return &MVRTypes.Symbol{
		UUID:   a.UUID,
		SymDef: MVRTypes.NodeReference[MVRTypes.SymDef]{String: &a.SymDef},
		Matrix: a.Matrix.ToMeshMatrix(),
	}
}
