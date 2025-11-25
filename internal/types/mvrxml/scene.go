package MVRXML

type Scene struct {
	AuxData *AuxData `xml:"AuxData"`
	Layers  []*Layer `xml:"Layers>Layer"`
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
	Geometries Geometries
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
	UUID   string  `xml:"uuid,attr"`
	Name   string  `xml:"name,attr"`
	Matrix *Matrix `xml:"Matrix"`
	ChildList
}
