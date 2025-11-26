package MVRXML

type Scene struct {
	AuxData *AuxData `xml:"AUXData"`
	Layers  []*Layer `xml:"Layers>Layer"`
}

type AuxData struct {
	SymDefs            []*SymDef            `xml:"Symdef"`
	Positions          []*Position          `xml:"Position"`
	MappingDefinitions []*MappingDefinition `xml:"MappingDefinition"`
	Classes            []*Class             `xml:"Class"`
}

type SymDef struct {
	UUID       string     `xml:"uuid,attr"`
	Name       string     `xml:"name,attr"`
	Geometries Geometries `xml:"ChildList"`
}

type Position struct {
	UUID string `xml:"uuid,attr"`
	Name string `xml:"name,attr"`
}

type MappingDefinition struct {
	UUID           string  `xml:"uuid,attr"`
	Name           string  `xml:"name,attr"`
	SizeX          int     `xml:"SizeX"`
	SizeY          int     `xml:"SizeY"`
	Source         Source  `xml:"Source"`
	ScaleHandeling *string `xml:"ScaleHandeling"` // ScaleKeepRatio or ScaleIgnoreRatio or KeepSizeCenter
}

type Class struct {
	UUID string `xml:"uuid,attr"`
	Name string `xml:"name,attr"`
}

type Layer struct {
	UUID   string  `xml:"uuid,attr"`
	Name   string  `xml:"name,attr"`
	Matrix *Matrix `xml:"Matrix"`
	ChildList
}
