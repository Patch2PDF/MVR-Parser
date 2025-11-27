package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type Fixture struct {
	UUID             string
	Name             string
	Multipatch       *string
	Matrix           MeshTypes.Matrix
	Class            *string
	GDTFSpec         fileName
	GDTFMode         string
	Focus            string
	CastShadow       bool
	DMXInvertPan     bool
	DMXInvertTilt    bool
	Position         *string
	Function         *string
	FixtureID        string
	FixtureIDNumeric int // can be 0 e.g. in MA export
	UnitNumber       int
	ChildPosition    string // TODO: check what this is for
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
	LinkedDef string
	Ux        int
	Uy        int
	Ox        int
	Oy        int
	Rz        float32
}
