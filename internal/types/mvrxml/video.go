package MVRXML

import (
	"strconv"

	GDTFReader "github.com/Patch2PDF/MVR-Parser/internal/gdtfreader"
	MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"
)

type VideoScreen struct {
	UUID             string           `xml:"uuid,attr"`
	Name             string           `xml:"name,attr"`
	Multipatch       string           `xml:"multipatch,attr,omitempty"`
	Matrix           *Matrix          `xml:"Matrix,omitempty"`
	Class            *string          `xml:"Classing,omitempty"`
	Geometries       Geometries       `xml:"Geometries"`
	Sources          []*Source        `xml:"Sources>Source"`
	Function         *string          `xml:"Function,omitempty"`
	GDTFSpec         fileName         `xml:"GDTFSpec"`
	GDTFMode         string           `xml:"GDTFMode"`
	CastShadow       bool             `xml:"CastShadow"`
	Addresses        *Addresses       `xml:"Addresses"`
	Alignments       []*Alignment     `xml:"Alignments>Alignment"`
	CustomCommands   []*CustomCommand `xml:"CustomCommands>CustomCommand"`
	Overwrites       []*Overwrite     `xml:"Overwrites>Overwrite"`
	Connections      []*Connection    `xml:"Connections>Connection"`
	FixtureID        string           `xml:"FixtureID"`
	FixtureIDNumeric int              `xml:"FixtureIDNumeric"`
	UnitNumber       int              `xml:"UnitNumber"`
	CustomId         int              `xml:"CustomId"`
	CustomIdType     int              `xml:"CustomIdType"`
	ChildList
}

func (a *VideoScreen) Parse(config ParseConfigData) *MVRTypes.VideoScreen {
	fixtureIDNumeric := a.FixtureIDNumeric
	if a.FixtureIDNumeric == 0 {
		value, err := strconv.ParseInt(a.FixtureID, 10, 0)
		if err != nil {
			// TODO: return err
		}
		fixtureIDNumeric = int(value)
	}
	GDTFReader.AddToTaskMap(config.GDTFTaskMap, a.GDTFSpec, a.GDTFMode)
	return &MVRTypes.VideoScreen{
		UUID:             a.UUID,
		Name:             a.Name,
		Multipatch:       a.Multipatch,
		Matrix:           a.Matrix.ToMeshMatrix(),
		Class:            MVRTypes.NodeReference[MVRTypes.Class]{String: a.Class},
		GDTFSpec:         MVRTypes.NodeReference[MVRTypes.GDTF]{String: &a.GDTFSpec},
		GDTFMode:         a.GDTFMode,
		CastShadow:       a.CastShadow,
		Function:         a.Function,
		FixtureID:        a.FixtureID,
		FixtureIDNumeric: fixtureIDNumeric,
		UnitNumber:       a.UnitNumber,
		Addresses:        a.Addresses.Parse(config),
		Alignments:       ParseList(config, &a.Alignments),
		CustomCommands:   ParseList(config, &a.CustomCommands),
		Overwrites:       ParseList(config, &a.Overwrites),
		Connections:      ParseList(config, &a.Connections),
		CustomId:         a.CustomId,
		CustomIdType:     a.CustomIdType,
		ChildList:        a.ChildList.Parse(config),
		Geometries:       a.Geometries.Parse(config),
		Sources:          ParseList(config, &a.Sources),
	}
}

type Projector struct {
	UUID             string           `xml:"uuid,attr"`
	Name             string           `xml:"name,attr"`
	Multipatch       string           `xml:"multipatch,attr,omitempty"`
	Matrix           *Matrix          `xml:"Matrix,omitempty"`
	Class            *string          `xml:"Classing,omitempty"`
	Geometries       Geometries       `xml:"Geometries"`
	Projections      []*Projection    `xml:"Projections>Projection"`
	Function         *string          `xml:"Function,omitempty"`
	GDTFSpec         fileName         `xml:"GDTFSpec"`
	GDTFMode         string           `xml:"GDTFMode"`
	CastShadow       bool             `xml:"CastShadow"`
	Addresses        *Addresses       `xml:"Addresses"`
	Alignments       []*Alignment     `xml:"Alignments>Alignment"`
	CustomCommands   []*CustomCommand `xml:"CustomCommands>CustomCommand"`
	Overwrites       []*Overwrite     `xml:"Overwrites>Overwrite"`
	Connections      []*Connection    `xml:"Connections>Connection"`
	FixtureID        string           `xml:"FixtureID"`
	FixtureIDNumeric int              `xml:"FixtureIDNumeric"`
	UnitNumber       int              `xml:"UnitNumber"`
	CustomId         int              `xml:"CustomId"`
	CustomIdType     int              `xml:"CustomIdType"`
	ChildList
}

func (a *Projector) Parse(config ParseConfigData) *MVRTypes.Projector {
	fixtureIDNumeric := a.FixtureIDNumeric
	if a.FixtureIDNumeric == 0 {
		value, err := strconv.ParseInt(a.FixtureID, 10, 0)
		if err != nil {
			// TODO: return err
		}
		fixtureIDNumeric = int(value)
	}
	GDTFReader.AddToTaskMap(config.GDTFTaskMap, a.GDTFSpec, a.GDTFMode)
	return &MVRTypes.Projector{
		UUID:             a.UUID,
		Name:             a.Name,
		Multipatch:       a.Multipatch,
		Matrix:           a.Matrix.ToMeshMatrix(),
		Class:            MVRTypes.NodeReference[MVRTypes.Class]{String: a.Class},
		GDTFSpec:         MVRTypes.NodeReference[MVRTypes.GDTF]{String: &a.GDTFSpec},
		GDTFMode:         a.GDTFMode,
		CastShadow:       a.CastShadow,
		Function:         a.Function,
		FixtureID:        a.FixtureID,
		FixtureIDNumeric: fixtureIDNumeric,
		UnitNumber:       a.UnitNumber,
		Addresses:        a.Addresses.Parse(config),
		Alignments:       ParseList(config, &a.Alignments),
		CustomCommands:   ParseList(config, &a.CustomCommands),
		Overwrites:       ParseList(config, &a.Overwrites),
		Connections:      ParseList(config, &a.Connections),
		CustomId:         a.CustomId,
		CustomIdType:     a.CustomIdType,
		ChildList:        a.ChildList.Parse(config),
		Geometries:       a.Geometries.Parse(config),
		Projections:      ParseList(config, &a.Projections),
	}
}

type Source struct {
	LinkedGeometry string `xml:"linkedGeometry,attr"`
	Type           string `xml:"type,attr"`
	Value          string `xml:",innerxml"`
}

func (a *Source) Parse(config ParseConfigData) *MVRTypes.Source {
	return &MVRTypes.Source{
		LinkedGeometry: a.LinkedGeometry,
		Type:           a.Type,
		Value:          a.Value,
	}
}

type Projection struct {
	Source         Source `xml:"Source"`
	ScaleHandeling string `xml:"ScaleHandeling"`
}

func (a *Projection) Parse(config ParseConfigData) *MVRTypes.Projection {
	return &MVRTypes.Projection{
		Source:         *a.Source.Parse(config),
		ScaleHandeling: a.ScaleHandeling,
	}
}
