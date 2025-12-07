package MVRXML

import MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"

type Scene struct {
	AuxData *AuxData `xml:"AUXData"`
	Layers  []*Layer `xml:"Layers>Layer"`
}

func (s *Scene) Parse(config ParseConfigData) *MVRTypes.Scene {
	return &MVRTypes.Scene{
		AuxData: s.AuxData.Parse(config),
		Layers:  ParseList(config, &s.Layers),
	}
}

type AuxData struct {
	SymDefs            []*SymDef            `xml:"Symdef"`
	Positions          []*Position          `xml:"Position"`
	MappingDefinitions []*MappingDefinition `xml:"MappingDefinition"`
	Classes            []*Class             `xml:"Class"`
}

func (a *AuxData) Parse(config ParseConfigData) *MVRTypes.AuxData {
	return &MVRTypes.AuxData{
		SymDefs:            ParseList(config, &a.SymDefs),
		Positions:          ParseList(config, &a.Positions),
		MappingDefinitions: ParseList(config, &a.MappingDefinitions),
		Classes:            ParseList(config, &a.Classes),
	}
}

type SymDef struct {
	UUID       string     `xml:"uuid,attr"`
	Name       string     `xml:"name,attr"`
	Geometries Geometries `xml:"ChildList"`
}

func (a *SymDef) Parse(config ParseConfigData) *MVRTypes.SymDef {
	return &MVRTypes.SymDef{
		UUID:       a.UUID,
		Name:       a.Name,
		Geometries: a.Geometries.Parse(config),
	}
}

type Position struct {
	UUID string `xml:"uuid,attr"`
	Name string `xml:"name,attr"`
}

func (a *Position) Parse(config ParseConfigData) *MVRTypes.Position {
	return &MVRTypes.Position{
		UUID: a.UUID,
		Name: a.Name,
	}
}

type MappingDefinition struct {
	UUID           string  `xml:"uuid,attr"`
	Name           string  `xml:"name,attr"`
	SizeX          int     `xml:"SizeX"`
	SizeY          int     `xml:"SizeY"`
	Source         Source  `xml:"Source"`
	ScaleHandeling *string `xml:"ScaleHandeling"` // ScaleKeepRatio or ScaleIgnoreRatio or KeepSizeCenter
}

func (a *MappingDefinition) Parse(config ParseConfigData) *MVRTypes.MappingDefinition {
	return &MVRTypes.MappingDefinition{
		UUID:           a.UUID,
		Name:           a.Name,
		SizeX:          a.SizeX,
		SizeY:          a.SizeY,
		Source:         MVRTypes.Source(a.Source),
		ScaleHandeling: a.ScaleHandeling,
	}
}

type Class struct {
	UUID string `xml:"uuid,attr"`
	Name string `xml:"name,attr"`
}

func (a *Class) Parse(config ParseConfigData) *MVRTypes.Class {
	return &MVRTypes.Class{
		UUID: a.UUID,
		Name: a.Name,
	}
}

type Layer struct {
	UUID   string  `xml:"uuid,attr"`
	Name   string  `xml:"name,attr"`
	Matrix *Matrix `xml:"Matrix"`
	ChildList
}

func (a *Layer) Parse(config ParseConfigData) *MVRTypes.Layer {
	return &MVRTypes.Layer{
		UUID:      a.UUID,
		Name:      a.Name,
		Matrix:    a.Matrix.ToMeshMatrix(),
		ChildList: a.ChildList.Parse(config),
	}
}
