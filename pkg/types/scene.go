package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type Scene struct {
	AuxData *AuxData
	Layers  []*Layer
}

// auxiliary data for the scene node
type AuxData struct {
	SymDefs            []*SymDef
	Positions          []*Position
	MappingDefinitions []*MappingDefinition
	Classes            []*Class
}

// contains the graphics so the scene can refer to this, thus optimizing repetition of the geometry
type SymDef struct {
	UUID       string
	Name       string
	Geometries *Geometries
}

// logical grouping of lighting devices and trusses
type Position struct {
	UUID string
	Name string
}

// input source for fixture color mapping applications
type MappingDefinition struct {
	UUID           string
	Name           string
	SizeX          int
	SizeY          int
	Source         Source
	ScaleHandeling *string // ScaleKeepRatio or ScaleIgnoreRatio or KeepSizeCenter
}

// logical grouping across different layers
type Class struct {
	UUID string
	Name string
}

// spatial representation of a geometric container
type Layer struct {
	UUID   string
	Name   string
	Matrix MeshTypes.Matrix
	ChildList
}
