package MVRTypes

import (
	"archive/zip"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
)

type Scene struct {
	AuxData *AuxData
	Layers  []*Layer
}

func (a *Scene) CreateReferencePointer() {
	a.AuxData.CreateReferencePointer()
	CreateReferencePointers(&a.Layers)
}

func (a *Scene) ResolveReference() {
	ResolveReferences(&a.Layers)
}

// auxiliary data for the scene node
type AuxData struct {
	SymDefs            []*SymDef
	Positions          []*Position
	MappingDefinitions []*MappingDefinition
	Classes            []*Class
}

func (a *AuxData) CreateReferencePointer() {
	CreateReferencePointers(&a.SymDefs)
	CreateReferencePointers(&a.Positions)
	CreateReferencePointers(&a.MappingDefinitions)
	CreateReferencePointers(&a.Classes)
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

func (a *SymDef) CreateReferencePointer() {
	refPointers.SymDefs[a.UUID] = a
}

func (a *SymDef) ResolveReference() {
	a.Geometries.ResolveReference()
}

func (a *SymDef) ReadMesh(fileMap map[string]*zip.File) error {
	return a.Geometries.ReadMesh(fileMap)
}

// logical grouping of lighting devices and trusses
type Position struct {
	UUID string
	Name string
}

func (a *Position) CreateReferencePointer() {
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

func (a *MappingDefinition) CreateReferencePointer() {
	refPointers.MappingDefinitions[a.UUID] = a
}

// logical grouping across different layers
type Class struct {
	UUID string
	Name string
}

func (a *Class) CreateReferencePointer() {
	refPointers.Classes[a.UUID] = a
}

// spatial representation of a geometric container
type Layer struct {
	UUID   string
	Name   string
	Matrix MeshTypes.Matrix
	ChildList
}

func (a *Layer) CreateReferencePointer() {
	a.ChildList.CreateReferencePointer()
}

func (a *Layer) ResolveReference() {
	a.ChildList.ResolveReference()
}
