package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type Truss struct {
	UUID             string
	Name             string
	Multipatch       string
	Matrix           MeshTypes.Matrix
	Class            NodeReference[Class]    // TODO: Node reference
	Position         NodeReference[Position] // TODO: Node reference
	Geometries       *Geometries
	Function         *string
	GDTFSpec         NodeReference[GDTF] // TODO: Node reference
	GDTFMode         string
	CastShadow       bool
	Addresses        *Addresses
	Alignments       []*Alignment
	CustomCommands   []*CustomCommand
	Overwrites       []*Overwrite
	Connections      []*Connection
	ChildPosition    *string // Node link to the geometry. Starting point is the Geometry Collect of the linked parent GDTF of this object.
	FixtureID        string
	FixtureIDNumeric int // v1.6 only, MA exports 1.5
	UnitNumber       int
	CustomId         int
	CustomIdType     int
	ChildList
}
