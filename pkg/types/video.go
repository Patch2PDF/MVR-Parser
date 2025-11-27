package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type VideoScreen struct {
	UUID             string
	Name             string
	Multipatch       string
	Matrix           MeshTypes.Matrix
	Class            NodeReference[Class] // TODO: Node reference
	Geometries       *Geometries
	Sources          []*Source
	Function         *string
	GDTFSpec         NodeReference[GDTF] // TODO: Node reference
	GDTFMode         string
	CastShadow       bool
	Addresses        *Addresses
	Alignments       []*Alignment
	CustomCommands   []*CustomCommand
	Overwrites       []*Overwrite
	Connections      []*Connection
	FixtureID        string
	FixtureIDNumeric int
	UnitNumber       int
	CustomId         int
	CustomIdType     int
	ChildList
}

type Projector struct {
	UUID             string
	Name             string
	Multipatch       string
	Matrix           MeshTypes.Matrix
	Class            NodeReference[Class] // TODO: Node reference
	Geometries       *Geometries
	Projections      []*Projection
	Function         *string
	GDTFSpec         NodeReference[GDTF] // TODO: Node reference
	GDTFMode         string
	CastShadow       bool
	Addresses        *Addresses
	Alignments       []*Alignment
	CustomCommands   []*CustomCommand
	Overwrites       []*Overwrite
	Connections      []*Connection
	FixtureID        string
	FixtureIDNumeric int
	UnitNumber       int
	CustomId         int
	CustomIdType     int
	ChildList
}

type Source struct {
	LinkedGeometry string // reference to a geometry in the gdtf
	Type           string
	Value          string
}

type Projection struct {
	Source         Source
	ScaleHandeling string
}
