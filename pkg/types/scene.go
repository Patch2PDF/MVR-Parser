package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type Scene struct {
	AuxData *AuxData
	Layers  []*Layer
}

type AuxData struct {
	SymDefs            []*SymDef
	Positions          []*Position
	MappingDefinitions []*MappingDefinition
	Classes            []*Class
}

type SymDef struct {
	UUID       string
	Name       string
	Geometries *Geometries
}

type Position struct {
	UUID string
	Name string
}

type MappingDefinition struct {
	UUID           string
	Name           string
	SizeX          int
	SizeY          int
	Source         Source
	ScaleHandeling *string // ScaleKeepRatio or ScaleIgnoreRatio or KeepSizeCenter
}

type Class struct {
	UUID string
	Name string
}

type Layer struct {
	UUID   string
	Name   string
	Matrix MeshTypes.Matrix
	ChildList
}
