package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type Support struct {
	UUID             string
	Name             string
	Multipatch       string
	Matrix           MeshTypes.Matrix
	Class            NodeReference[Class]    // TODO: Node reference
	Position         NodeReference[Position] // TODO: Node reference
	Geometries       *Geometries
	Function         *string
	ChainLength      float32
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

func (a *Support) CreateReferencePointer() {
	a.ChildList.CreateReferencePointer()
}

func (a *Support) ResolveReference() {
	if a.Class.String != nil {
		a.Class.Ptr = refPointers.Classes[*a.Class.String]
	}
	// a.GDTFSpec.Ptr = refPointers.Classes[*a.Class.String] // TODO:
	if a.Position.String != nil {
		a.Position.Ptr = refPointers.Positions[*a.Position.String]
	}
	a.Geometries.ResolveReference()
	a.ChildList.ResolveReference()
}
