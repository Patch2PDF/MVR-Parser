package MVRTypes

import (
	"archive/zip"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
)

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

func (a *VideoScreen) CreateReferencePointer(refPointers *ReferencePointers) {
	a.ChildList.CreateReferencePointer(refPointers)
}

func (a *VideoScreen) ResolveReference(refPointers *ReferencePointers) {
	if a.Class.String != nil {
		a.Class.Ptr = refPointers.Classes[*a.Class.String]
	}
	if a.GDTFSpec.String != nil {
		a.GDTFSpec.Ptr = refPointers.GDTFSpecs[*a.GDTFSpec.String]
	}
	a.Geometries.ResolveReference(refPointers)
	a.ChildList.ResolveReference(refPointers)
}

func (a *VideoScreen) ReadMesh(fileMap map[string]*zip.File) error {
	err := a.Geometries.ReadMesh(fileMap)
	if err != nil {
		return err
	}
	return a.ChildList.ReadMesh(fileMap)
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

func (a *Projector) CreateReferencePointer(refPointers *ReferencePointers) {
	a.ChildList.CreateReferencePointer(refPointers)
}

func (a *Projector) ResolveReference(refPointers *ReferencePointers) {
	if a.Class.String != nil {
		a.Class.Ptr = refPointers.Classes[*a.Class.String]
	}
	if a.GDTFSpec.String != nil {
		a.GDTFSpec.Ptr = refPointers.GDTFSpecs[*a.GDTFSpec.String]
	}
	a.Geometries.ResolveReference(refPointers)
	a.ChildList.ResolveReference(refPointers)
}

func (a *Projector) ReadMesh(fileMap map[string]*zip.File) error {
	err := a.Geometries.ReadMesh(fileMap)
	if err != nil {
		return err
	}
	return a.ChildList.ReadMesh(fileMap)
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
