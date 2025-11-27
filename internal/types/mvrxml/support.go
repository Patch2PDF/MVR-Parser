package MVRXML

import MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"

type Support struct {
	UUID             string           `xml:"uuid,attr"`
	Name             string           `xml:"name,attr"`
	Multipatch       string           `xml:"multipatch,attr,omitempty"`
	Matrix           *Matrix          `xml:"Matrix,omitempty"`
	Class            *string          `xml:"Classing,omitempty"`
	Position         *string          `xml:"Position,omitempty"`
	Geometries       Geometries       `xml:"Geometries"`
	Function         *string          `xml:"Function,omitempty"`
	ChainLength      float32          `xml:"ChainLength"`
	GDTFSpec         fileName         `xml:"GDTFSpec"`
	GDTFMode         string           `xml:"GDTFMode"`
	CastShadow       bool             `xml:"CastShadow"`
	Addresses        *Addresses       `xml:"Addresses"`
	Alignments       []*Alignment     `xml:"Alignments>Alignment"`
	CustomCommands   []*CustomCommand `xml:"CustomCommands>CustomCommand"`
	Overwrites       []*Overwrite     `xml:"Overwrites>Overwrite"`
	Connections      []*Connection    `xml:"Connections>Connection"`
	FixtureID        string           `xml:"FixtureID"`
	FixtureIDNumeric int              `xml:"FixtureIDNumeric"` // can be 0 e.g. in MA export
	UnitNumber       int              `xml:"UnitNumber"`
	CustomId         int              `xml:"CustomId"`
	CustomIdType     int              `xml:"CustomIdType"`
	ChildList
}

func (a *Support) Parse() *MVRTypes.Support {
	return &MVRTypes.Support{
		UUID:             a.UUID,
		Name:             a.Name,
		Multipatch:       a.Multipatch,
		Matrix:           a.Matrix.ToMeshMatrix(),
		Class:            a.Class,
		GDTFSpec:         a.GDTFSpec,
		GDTFMode:         a.GDTFMode,
		CastShadow:       a.CastShadow,
		Position:         a.Position,
		Function:         a.Function,
		FixtureID:        a.FixtureID,
		FixtureIDNumeric: a.FixtureIDNumeric,
		UnitNumber:       a.UnitNumber,
		Addresses:        a.Addresses.Parse(),
		Alignments:       ParseList(&a.Alignments),
		CustomCommands:   ParseList(&a.CustomCommands),
		Overwrites:       ParseList(&a.Overwrites),
		Connections:      ParseList(&a.Connections),
		CustomId:         a.CustomId,
		CustomIdType:     a.CustomIdType,
		ChildList:        a.ChildList.Parse(),
		Geometries:       a.Geometries.Parse(),
		ChainLength:      a.ChainLength,
	}
}
