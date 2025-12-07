package MVRXML

import (
	"strconv"

	GDTFReader "github.com/Patch2PDF/MVR-Parser/internal/gdtfreader"
	MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"
)

type Truss struct {
	UUID             string           `xml:"uuid,attr"`
	Name             string           `xml:"name,attr"`
	Multipatch       string           `xml:"multipatch,attr,omitempty"`
	Matrix           *Matrix          `xml:"Matrix,omitempty"`
	Class            *string          `xml:"Classing,omitempty"`
	Position         *string          `xml:"Position,omitempty"`
	Geometries       Geometries       `xml:"Geometries"`
	Function         *string          `xml:"Function,omitempty"`
	GDTFSpec         fileName         `xml:"GDTFSpec"`
	GDTFMode         string           `xml:"GDTFMode"`
	CastShadow       bool             `xml:"CastShadow"`
	Addresses        *Addresses       `xml:"Addresses"`
	Alignments       []*Alignment     `xml:"Alignments>Alignment"`
	CustomCommands   []*CustomCommand `xml:"CustomCommands>CustomCommand"`
	Overwrites       []*Overwrite     `xml:"Overwrites>Overwrite"`
	Connections      []*Connection    `xml:"Connections>Connection"`
	ChildPosition    *string          `xml:"ChildPosition,omitempty"` // Node link to the geometry. Starting point is the Geometry Collect of the linked parent GDTF of this object.
	FixtureID        string           `xml:"FixtureID"`
	FixtureIDNumeric int              `xml:"FixtureIDNumeric"` // v1.6 only, MA exports 1.5
	UnitNumber       int              `xml:"UnitNumber"`
	CustomId         int              `xml:"CustomId"`
	CustomIdType     int              `xml:"CustomIdType"`
	ChildList
}

func (a *Truss) Parse(config ParseConfigData) *MVRTypes.Truss {
	fixtureIDNumeric := a.FixtureIDNumeric
	if a.FixtureIDNumeric == 0 {
		value, err := strconv.ParseInt(a.FixtureID, 10, 0)
		if err != nil {
			// TODO: return err
		}
		fixtureIDNumeric = int(value)
	}
	GDTFReader.AddToTaskMap(config.GDTFTaskMap, a.GDTFSpec, a.GDTFMode)
	return &MVRTypes.Truss{
		UUID:             a.UUID,
		Name:             a.Name,
		Multipatch:       a.Multipatch,
		Matrix:           a.Matrix.ToMeshMatrix(),
		Class:            MVRTypes.NodeReference[MVRTypes.Class]{String: a.Class},
		GDTFSpec:         MVRTypes.NodeReference[MVRTypes.GDTF]{String: &a.GDTFSpec},
		GDTFMode:         a.GDTFMode,
		CastShadow:       a.CastShadow,
		Position:         MVRTypes.NodeReference[MVRTypes.Position]{String: a.Position},
		Function:         a.Function,
		FixtureID:        a.FixtureID,
		FixtureIDNumeric: fixtureIDNumeric,
		UnitNumber:       a.UnitNumber,
		ChildPosition:    a.ChildPosition,
		Addresses:        a.Addresses.Parse(config),
		Alignments:       ParseList(config, &a.Alignments),
		CustomCommands:   ParseList(config, &a.CustomCommands),
		Overwrites:       ParseList(config, &a.Overwrites),
		Connections:      ParseList(config, &a.Connections),
		CustomId:         a.CustomId,
		CustomIdType:     a.CustomIdType,
		ChildList:        a.ChildList.Parse(config),
		Geometries:       a.Geometries.Parse(config),
	}
}
