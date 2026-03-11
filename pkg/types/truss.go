package MVRTypes

import (
	"archive/zip"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/v2/pkg/MeshTypes"
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
	Model            TrussModel
	ChildList
}

func (a *Truss) CreateReferencePointer(refPointers *ReferencePointers) {
	a.ChildList.CreateReferencePointer(refPointers)
}

func (a *Truss) ResolveReference(refPointers *ReferencePointers) {
	if a.Class.String != nil {
		a.Class.Ptr = refPointers.Classes[*a.Class.String]
	}
	if a.GDTFSpec.String != nil {
		a.GDTFSpec.Ptr = refPointers.GDTFSpecs[*a.GDTFSpec.String]
	}
	if a.Position.String != nil {
		a.Position.Ptr = refPointers.Positions[*a.Position.String]
	}
	a.Geometries.ResolveReference(refPointers)
	a.ChildList.ResolveReference(refPointers)
}

func (a *Truss) ReadMesh(fileMap map[string]*zip.File) error {
	err := a.Geometries.ReadMesh(fileMap)
	if err != nil {
		return err
	}
	return a.ChildList.ReadMesh(fileMap)
}

func (a *Truss) addNodeModelsToStageModel(stageModel *StageModel, modelConfig ModelConfig, parentConfig ModelNodeConfig) {
	config := getConfigOverrides(modelConfig, parentConfig, a.UUID)

	if config.Exclude == nil || !(*config.Exclude) {
		stageModel.TrussModels = append(stageModel.TrussModels, a.Model.Copy())
	}

	a.ChildList.addNodeModelsToStageModel(stageModel, modelConfig, config)
}
