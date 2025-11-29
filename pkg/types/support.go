package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type Support struct {
	UUID             string
	Name             string
	Multipatch       string
	Matrix           MeshTypes.Matrix
	Class            NodeReference[Class]
	Position         NodeReference[Position]
	Geometries       *Geometries
	Function         *string
	ChainLength      float32
	GDTFSpec         NodeReference[GDTF]
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
	if a.GDTFSpec.String != nil {
		a.GDTFSpec.Ptr = refPointers.GDTFSpecs[*a.GDTFSpec.String]
	}
	if a.Position.String != nil {
		a.Position.Ptr = refPointers.Positions[*a.Position.String]
	}
	a.Geometries.ResolveReference()
	a.ChildList.ResolveReference()
}
