package MVRXML

type SceneObject struct {
	UUID             string           `xml:"uuid,attr"`
	Name             string           `xml:"name,attr"`
	Multipatch       string           `xml:"multipatch,attr,omitempty"`
	Matrix           *Matrix          `xml:"Matrix,omitempty"`
	Class            *string          `xml:"Classing,omitempty"`
	Geometries       Geometries       `xml:"Geometries"`
	GDTFSpec         fileName         `xml:"GDTFSpec"`
	GDTFMode         string           `xml:"GDTFMode"`
	CastShadow       bool             `xml:"CastShadow"`
	Addresses        *Addresses       `xml:"Addresses"`
	Alignments       []*Alignment     `xml:"Alignments>Alignment"`
	CustomCommands   []*CustomCommand `xml:"CustomCommands>CustomCommand"`
	Overwrites       []*Overwrite     `xml:"Overwrites>Overwrite"`
	Connections      []*Connection    `xml:"Connections>Connection"`
	FixtureID        string           `xml:"FixtureID"`
	FixtureIDNumeric int              `xml:"FixtureIDNumeric"` // can be 0 e.g. in MA export
	UnitNumber       int              `xml:"UnitNumber"`
	CustomId         int              `xml:"CustomId"`
	CustomIdType     int              `xml:"CustomIdType"`
	ChildList
}

type Alignment struct {
	Geometry  string `xml:"geometry,attr"`  // Defines the Beam Geometry that gets aligned.
	Up        Vector `xml:"up,attr"`        // default: 0,0,1
	Direction Vector `xml:"direction,attr"` // default: 0,0,-1
}

type CustomCommand struct {
	//TODO:
	// identical to GDTF ChannelFunction (copy paste?)
}

// This node defines an overwrite with the Universal.gdtt GDTF template inside the MVR to overwrite Wheel Slots, Emitters and Filters for the fixture
type Overwrite struct {
	Universal string `xml:"universal,attr"` // Node Link to the Wheel, Emitter or Filter. Starting point is the the collect of the Universal GDTF.
	Target    string `xml:"target,attr"`    // Node Link to the Wheel, Emitter or Filter. Starting point is the the collect of the linked GDTF of the fixture. When no target is given, it will be like a static gobo or filter that you attach in front of all beams
}

type Connection struct {
	Own      string `xml:"own,attr"`      // Node Link to the Geometry with DIN SPEC 15800 Type Wiring Object . Starting point is the Geometry Collect of the linked GDTF.
	Other    string `xml:"other,attr"`    // Node Link to the Geometry with DIN SPEC 15800 Type Wiring Object . Starting point is the Geometry Collect of the linked GDTF of the object defined in toObject.
	ToObject string `xml:"toObject,attr"` // UUID of an other object in the scene.
}
