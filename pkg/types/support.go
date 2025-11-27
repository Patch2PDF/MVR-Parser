package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type Support struct {
	UUID             string
	Name             string
	Multipatch       string
	Matrix           MeshTypes.Matrix
	Class            *string
	Position         *string
	Geometries       *Geometries
	Function         *string
	ChainLength      float32
	GDTFSpec         fileName
	GDTFMode         string
	CastShadow       bool
	Addresses        *Addresses
	Alignments       []*Alignment
	CustomCommands   []*CustomCommand
	Overwrites       []*Overwrite
	Connections      []*Connection
	FixtureID        string
	FixtureIDNumeric int // can be 0 e.g. in MA export
	UnitNumber       int
	CustomId         int
	CustomIdType     int
	ChildList
}
