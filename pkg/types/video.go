package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type VideoScreen struct {
	UUID             string
	Name             string
	Multipatch       string
	Matrix           MeshTypes.Matrix
	Class            NodeReference[Class]
	Geometries       *Geometries
	Sources          []*Source
	Function         *string
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

func (a *VideoScreen) CreateReferencePointer() {
	a.ChildList.CreateReferencePointer()
}

func (a *VideoScreen) ResolveReference() {
	if a.Class.String != nil {
		a.Class.Ptr = refPointers.Classes[*a.Class.String]
	}
	if a.GDTFSpec.String != nil {
		a.GDTFSpec.Ptr = refPointers.GDTFSpecs[*a.GDTFSpec.String]
	}
	a.Geometries.ResolveReference()
	a.ChildList.ResolveReference()
}

type Projector struct {
	UUID             string
	Name             string
	Multipatch       string
	Matrix           MeshTypes.Matrix
	Class            NodeReference[Class]
	Geometries       *Geometries
	Projections      []*Projection
	Function         *string
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

func (a *Projector) CreateReferencePointer() {
	a.ChildList.CreateReferencePointer()
}

func (a *Projector) ResolveReference() {
	if a.Class.String != nil {
		a.Class.Ptr = refPointers.Classes[*a.Class.String]
	}
	if a.GDTFSpec.String != nil {
		a.GDTFSpec.Ptr = refPointers.GDTFSpecs[*a.GDTFSpec.String]
	}
	a.Geometries.ResolveReference()
	a.ChildList.ResolveReference()
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
