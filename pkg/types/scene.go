package MVRTypes

import (
	"archive/zip"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
)

type Scene struct {
	AuxData *AuxData
	Layers  []*Layer
}

func (a *Scene) CreateReferencePointer(refPointers *ReferencePointers) {
	a.AuxData.CreateReferencePointer(refPointers)
	CreateReferencePointers(refPointers, &a.Layers)
}

func (a *Scene) ResolveReference(refPointers *ReferencePointers) {
	ResolveReferences(refPointers, &a.Layers)
}

// auxiliary data for the scene node
type AuxData struct {
	SymDefs            []*SymDef
	Positions          []*Position
	MappingDefinitions []*MappingDefinition
	Classes            []*Class
}

func (a *AuxData) CreateReferencePointer(refPointers *ReferencePointers) {
	CreateReferencePointers(refPointers, &a.SymDefs)
	CreateReferencePointers(refPointers, &a.Positions)
	CreateReferencePointers(refPointers, &a.MappingDefinitions)
	CreateReferencePointers(refPointers, &a.Classes)
}

func (a *AuxData) ReadMesh(fileMap map[string]*zip.File) error {
	return ReadMeshes(a.SymDefs, fileMap)
}

// contains the graphics so the scene can refer to this, thus optimizing repetition of the geometry
type SymDef struct {
	UUID       string
	Name       string
	Geometries *Geometries
}

func (a *SymDef) CreateReferencePointer(refPointers *ReferencePointers) {
	refPointers.SymDefs[a.UUID] = a
}

func (a *SymDef) ResolveReference(refPointers *ReferencePointers) {
	a.Geometries.ResolveReference(refPointers)
}

func (a *SymDef) ReadMesh(fileMap map[string]*zip.File) error {
	return a.Geometries.ReadMesh(fileMap)
}

// logical grouping of lighting devices and trusses
type Position struct {
	UUID string
	Name string
}

func (a *Position) CreateReferencePointer(refPointers *ReferencePointers) {
	refPointers.Positions[a.UUID] = a
}

// input source for fixture color mapping applications
type MappingDefinition struct {
	UUID           string
	Name           string
	SizeX          int
	SizeY          int
	Source         Source
	ScaleHandeling *string // ScaleKeepRatio or ScaleIgnoreRatio or KeepSizeCenter
}

func (a *MappingDefinition) CreateReferencePointer(refPointers *ReferencePointers) {
	refPointers.MappingDefinitions[a.UUID] = a
}

// logical grouping across different layers
type Class struct {
	UUID string
	Name string
}

func (a *Class) CreateReferencePointer(refPointers *ReferencePointers) {
	refPointers.Classes[a.UUID] = a
}

// spatial representation of a geometric container
type Layer struct {
	UUID   string
	Name   string
	Matrix MeshTypes.Matrix
	ChildList
}

func (a *Layer) CreateReferencePointer(refPointers *ReferencePointers) {
	a.ChildList.CreateReferencePointer(refPointers)
}

func (a *Layer) ResolveReference(refPointers *ReferencePointers) {
	a.ChildList.ResolveReference(refPointers)
}
