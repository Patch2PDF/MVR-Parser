package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type Fixture struct {
	UUID             string
	Name             string
	Multipatch       *string
	Matrix           MeshTypes.Matrix
	Class            NodeReference[Class]
	GDTFSpec         NodeReference[GDTF]
	GDTFMode         string
	Focus            NodeReference[FocusPoint] // TODO:
	CastShadow       bool
	DMXInvertPan     bool
	DMXInvertTilt    bool
	Position         NodeReference[Position]
	Function         *string
	FixtureID        string
	FixtureIDNumeric int
	UnitNumber       int
	ChildPosition    string // Node link to the geometry. Starting point is the Geometry Collect of the linked parent GDTF of this object.
	Addresses        *Addresses
	Protocols        []*Protocol
	Alignments       []*Alignment
	CustomCommands   []*CustomCommand
	Overwrites       []*Overwrite
	Connections      []*Connection
	Color            *ColorCIE
	CustomId         int
	CustomIdType     int
	Mappings         []*Mapping
	Gobo             *Gobo
	ChildList
}

func (a *Fixture) CreateReferencePointer() {
	a.ChildList.CreateReferencePointer()
}

func (a *Fixture) ResolveReference() {
	if a.Class.String != nil {
		a.Class.Ptr = refPointers.Classes[*a.Class.String]
	}
	// a.GDTFSpec.Ptr = refPointers.Classes[*a.Class.String] // TODO:
	if a.Position.String != nil {
		a.Position.Ptr = refPointers.Positions[*a.Position.String]
	}
	ResolveReferences(&a.Mappings)
}

type Gobo struct {
	Rotation float32
}

type Protocol struct {
	Geometry     string // defaults to NetworkInOut_1
	Name         string // Custom Name of the protocol to identify the protocol. Needs to be unique for this instance of object.
	Type         string // Name of the protocol.
	Version      string // This is the protocol version if available.
	Transmission string // Unicast, Multicast, Broadcast, Anycast
}

type Mapping struct {
	LinkedDef NodeReference[MappingDefinition]
	Ux        int
	Uy        int
	Ox        int
	Oy        int
	Rz        float32
}

func (a *Mapping) ResolveReference() {
	a.LinkedDef.Ptr = refPointers.MappingDefinitions[*a.LinkedDef.String]
}
