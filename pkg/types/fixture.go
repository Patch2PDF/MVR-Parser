package MVRTypes

import (
	"archive/zip"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
)

type Fixture struct {
	UUID             string
	Name             string
	Multipatch       *string
	Matrix           MeshTypes.Matrix
	Class            NodeReference[Class]
	GDTFSpec         NodeReference[GDTF]
	GDTFMode         string
	Focus            NodeReference[FocusPoint]
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

func (a *Fixture) CreateReferencePointer(refPointers *ReferencePointers) {
	a.ChildList.CreateReferencePointer(refPointers)
}

func (a *Fixture) ResolveReference(refPointers *ReferencePointers) {
	if a.Class.String != nil {
		a.Class.Ptr = refPointers.Classes[*a.Class.String]
	}
	if a.GDTFSpec.String != nil {
		a.GDTFSpec.Ptr = refPointers.GDTFSpecs[*a.GDTFSpec.String]
	}
	if a.Position.String != nil {
		a.Position.Ptr = refPointers.Positions[*a.Position.String]
	}
	if a.Focus.String != nil {
		a.Focus.Ptr = refPointers.FoucsPoints[*a.Focus.String]
	}
	ResolveReferences(refPointers, &a.Mappings)
}

func (a *Fixture) ReadMesh(fileMap map[string]*zip.File) error {
	return a.ChildList.ReadMesh(fileMap)
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

func (a *Mapping) ResolveReference(refPointers *ReferencePointers) {
	a.LinkedDef.Ptr = refPointers.MappingDefinitions[*a.LinkedDef.String]
}
