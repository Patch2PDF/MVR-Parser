package MVRXML

type Fixture struct {
	UUID             string           `xml:"uuid,attr"`
	Name             string           `xml:"name,attr"`
	Multipatch       *string          `xml:"multipatch,attr,omitempty"`
	Matrix           *Matrix          `xml:"Matrix,omitempty"`
	Class            *string          `xml:"Classing,omitempty"`
	GDTFSpec         fileName         `xml:"GDTFSpec"`
	GDTFMode         string           `xml:"GDTFMode"`
	Focus            string           `xml:"Foces"`
	CastShadow       bool             `xml:"CastShadow"`
	DMXInvertPan     bool             `xml:"DMXInvertPan"`
	DMXInvertTilt    bool             `xml:"DMXInvertTilt"`
	Position         *string          `xml:"Position,omitempty"`
	Function         *string          `xml:"Function,omitempty"`
	FixtureID        string           `xml:"FixtureID"`
	FixtureIDNumeric int              `xml:"FixtureIDNumeric"` // can be 0 e.g. in MA export
	UnitNumber       int              `xml:"UnitNumber"`
	ChildPosition    string           `xml:"ChildPosition"` // TODO: check what this is for
	Addresses        *Addresses       `xml:"Addresses"`
	Protocols        []*Protocol      `xml:"Protocols>Protocol"`
	Alignments       []*Alignment     `xml:"Alignments>Alignment"`
	CustomCommands   []*CustomCommand `xml:"CustomCommands>CustomCommand"`
	Overwrites       []*Overwrite     `xml:"Overwrites>Overwrite"`
	Connections      []*Connection    `xml:"Connections>Connection"`
	Color            *ColorCIE        `xml:"Color"`
	CustomId         int              `xml:"CustomId"`
	CustomIdType     int              `xml:"CustomIdType"`
	Mappings         []*Mapping       `xml:"Mappings>Mapping"`
	Gobo             *Gobo            `xml:"Gobo,omitempty"`
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
