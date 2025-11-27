package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type Truss struct {
	UUID             string
	Name             string
	Multipatch       string
	Matrix           MeshTypes.Matrix
	Class            *string
	Position         *string
	Geometries       *Geometries
	Function         *string
	GDTFSpec         fileName
	GDTFMode         string
	CastShadow       bool
	Addresses        *Addresses
	Alignments       []*Alignment
	CustomCommands   []*CustomCommand
	Overwrites       []*Overwrite
	Connections      []*Connection
	ChildPosition    *string // TODO: check what this is for
	FixtureID        string
	FixtureIDNumeric int // v1.6 only, MA exports 1.5
	UnitNumber       int
	CustomId         int
	CustomIdType     int
	ChildList
}
