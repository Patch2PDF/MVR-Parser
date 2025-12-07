package MVRTypes

import (
	"archive/zip"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
)

type Truss struct {
	UUID             string
	Name             string
	Multipatch       string
	Matrix           MeshTypes.Matrix
	Class            NodeReference[Class]
	Position         NodeReference[Position]
	Geometries       *Geometries
	Function         *string
	GDTFSpec         NodeReference[GDTF]
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

func (a *Truss) CreateReferencePointer() {
	a.ChildList.CreateReferencePointer()
}

func (a *Truss) ResolveReference() {
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

func (a *Truss) ReadMesh(fileMap map[string]*zip.File) error {
	err := a.Geometries.ReadMesh(fileMap)
	if err != nil {
		return err
	}
	return a.ChildList.ReadMesh(fileMap)
}
