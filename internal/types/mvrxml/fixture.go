package MVRXML

import (
	"strconv"

	GDTFReader "github.com/Patch2PDF/MVR-Parser/internal/gdtfreader"
	MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"
)

type Fixture struct {
	UUID             string           `xml:"uuid,attr"`
	Name             string           `xml:"name,attr"`
	Multipatch       *string          `xml:"multipatch,attr,omitempty"`
	Matrix           *Matrix          `xml:"Matrix,omitempty"`
	Class            *string          `xml:"Classing,omitempty"`
	GDTFSpec         fileName         `xml:"GDTFSpec"`
	GDTFMode         string           `xml:"GDTFMode"`
	Focus            string           `xml:"Foces"`
	CastShadow       bool             `xml:"CastShadow"`
	DMXInvertPan     bool             `xml:"DMXInvertPan"`
	DMXInvertTilt    bool             `xml:"DMXInvertTilt"`
	Position         *string          `xml:"Position,omitempty"`
	Function         *string          `xml:"Function,omitempty"`
	FixtureID        string           `xml:"FixtureID"`
	FixtureIDNumeric int              `xml:"FixtureIDNumeric"`
	UnitNumber       int              `xml:"UnitNumber"`
	ChildPosition    string           `xml:"ChildPosition"`
	Addresses        *Addresses       `xml:"Addresses"`
	Protocols        []*Protocol      `xml:"Protocols>Protocol"`
	Alignments       []*Alignment     `xml:"Alignments>Alignment"`
	CustomCommands   []*CustomCommand `xml:"CustomCommands>CustomCommand"`
	Overwrites       []*Overwrite     `xml:"Overwrites>Overwrite"`
	Connections      []*Connection    `xml:"Connections>Connection"`
	Color            *ColorCIE        `xml:"Color"`
	CustomId         int              `xml:"CustomId"`
	CustomIdType     int              `xml:"CustomIdType"`
	Mappings         []*Mapping       `xml:"Mappings>Mapping"`
	Gobo             *Gobo            `xml:"Gobo,omitempty"`
	ChildList
}

func (a *Fixture) Parse(config ParseConfigData) *MVRTypes.Fixture {
	// helper as e.g. MA3 does not export this (MVR 1.5)
	fixtureIDNumeric := a.FixtureIDNumeric
	if a.FixtureIDNumeric == 0 {
		value, err := strconv.ParseInt(a.FixtureID, 10, 0)
		if err != nil {
			// TODO: return err
		}
		fixtureIDNumeric = int(value)
	}
	GDTFReader.AddToTaskMap(config.GDTFTaskMap, a.GDTFSpec, a.GDTFMode)
	return &MVRTypes.Fixture{
		UUID:             a.UUID,
		Name:             a.Name,
		Multipatch:       a.Multipatch,
		Matrix:           a.Matrix.ToMeshMatrix(),
		Class:            MVRTypes.NodeReference[MVRTypes.Class]{String: a.Class},
		GDTFSpec:         MVRTypes.NodeReference[MVRTypes.GDTF]{String: &a.GDTFSpec},
		GDTFMode:         a.GDTFMode,
		Focus:            MVRTypes.NodeReference[MVRTypes.FocusPoint]{String: &a.Focus},
		CastShadow:       a.CastShadow,
		DMXInvertPan:     a.DMXInvertPan,
		DMXInvertTilt:    a.DMXInvertTilt,
		Position:         MVRTypes.NodeReference[MVRTypes.Position]{String: a.Position},
		Function:         a.Function,
		FixtureID:        a.FixtureID,
		FixtureIDNumeric: fixtureIDNumeric,
		UnitNumber:       a.UnitNumber,
		ChildPosition:    a.ChildPosition,
		Addresses:        a.Addresses.Parse(config),
		Protocols:        ParseList(config, &a.Protocols),
		Alignments:       ParseList(config, &a.Alignments),
		CustomCommands:   ParseList(config, &a.CustomCommands),
		Overwrites:       ParseList(config, &a.Overwrites),
		Connections:      ParseList(config, &a.Connections),
		Color:            (*MVRTypes.ColorCIE)(a.Color),
		CustomId:         a.CustomId,
		CustomIdType:     a.CustomIdType,
		Mappings:         ParseList(config, &a.Mappings),
		Gobo:             (*MVRTypes.Gobo)(a.Gobo),
		ChildList:        a.ChildList.Parse(config),
	}
}

type Gobo struct {
	Rotation float32
}

type Protocol struct {
	Geometry     string // defaults to NetworkInOut_1
	Name         string // Custom Name of the protocol to identify the protocol. Needs to be unique for this instance of object.
	Type         string // Name of the protocol.
	Version      string // This is the protocol version if available.
	Transmission string // Unicast, Multicast, Broadcast, Anycast
}

func (a *Protocol) Parse(config ParseConfigData) *MVRTypes.Protocol {
	return &MVRTypes.Protocol{
		Geometry:     a.Geometry,
		Name:         a.Name,
		Type:         a.Type,
		Version:      a.Version,
		Transmission: a.Transmission,
	}
}

type Mapping struct {
	LinkedDef string
	Ux        int
	Uy        int
	Ox        int
	Oy        int
	Rz        float32
}

func (a *Mapping) Parse(config ParseConfigData) *MVRTypes.Mapping {
	return &MVRTypes.Mapping{
		LinkedDef: MVRTypes.NodeReference[MVRTypes.MappingDefinition]{String: &a.LinkedDef},
		Ux:        a.Ux,
		Uy:        a.Uy,
		Ox:        a.Ox,
		Oy:        a.Oy,
		Rz:        a.Rz,
	}
}
