package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type SceneObject struct {
	UUID             string
	Name             string
	Multipatch       string
	Matrix           MeshTypes.Matrix
	Class            NodeReference[Class]
	Geometries       *Geometries
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

func (a *SceneObject) CreateReferencePointer() {
	a.ChildList.CreateReferencePointer()
}

func (a *SceneObject) ResolveReference() {
	if a.Class.String != nil {
		a.Class.Ptr = refPointers.Classes[*a.Class.String]
	}
	if a.GDTFSpec.String != nil {
		a.GDTFSpec.Ptr = refPointers.GDTFSpecs[*a.GDTFSpec.String]
	}
	a.Geometries.ResolveReference()
	a.ChildList.ResolveReference()
}

type Alignment struct {
	Geometry  string // Defines the Beam Geometry that gets aligned.
	Up        Vector // default: 0,0,1
	Direction Vector // default: 0,0,-1
}

type CustomCommand struct {
	Object string
	Value  string
}

// This node defines an overwrite with the Universal.gdtt GDTF template inside the MVR to overwrite Wheel Slots, Emitters and Filters for the fixture
type Overwrite struct {
	Universal string // Node Link to the Wheel, Emitter or Filter. Starting point is the the collect of the Universal GDTF.
	Target    string // Node Link to the Wheel, Emitter or Filter. Starting point is the the collect of the linked GDTF of the fixture. When no target is given, it will be like a static gobo or filter that you attach in front of all beams
}

type Connection struct {
	Own      string             // Node Link to the Geometry with DIN SPEC 15800 Type Wiring Object . Starting point is the Geometry Collect of the linked GDTF.
	Other    string             // Node Link to the Geometry with DIN SPEC 15800 Type Wiring Object . Starting point is the Geometry Collect of the linked GDTF of the object defined in toObject.
	ToObject NodeReference[any] // UUID of an other object in the scene.
}
