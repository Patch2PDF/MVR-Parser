package MVRTypes

import (
	"archive/zip"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/v2/pkg/MeshTypes"
)

type Support struct {
	UUID             string
	Name             string
	Multipatch       string
	Matrix           MeshTypes.Matrix
	Class            NodeReference[Class]
	Position         NodeReference[Position]
	Geometries       *Geometries
	Function         *string
	ChainLength      float32
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
	Model            SupportModel
	ChildList
}

func (a *Support) CreateReferencePointer(refPointers *ReferencePointers) {
	a.ChildList.CreateReferencePointer(refPointers)
}

func (a *Support) ResolveReference(refPointers *ReferencePointers) {
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

func (a *Support) ReadMesh(fileMap map[string]*zip.File) error {
	err := a.Geometries.ReadMesh(fileMap)
	if err != nil {
		return err
	}
	return a.ChildList.ReadMesh(fileMap)
}

func (a *Support) addNodeModelsToStageModel(stageModel *StageModel, modelConfig ModelConfig, parentConfig ModelNodeConfig, parentParameters parentNodeParameters) {
	config := getConfigOverrides(modelConfig, parentConfig, a.UUID)

	validClass, classID := checkShouldIncludeClassInModel(modelConfig.ClassConfig, a.Class.String, parentParameters.classID)

	if (config.Exclude == nil || !(*config.Exclude)) && validClass {
		stageModel.SupportModels = append(stageModel.SupportModels, a.Model.Copy())
	}

	childParameters := parentNodeParameters{
		classID: classID,
	}

	a.ChildList.addNodeModelsToStageModel(stageModel, modelConfig, config, childParameters)
}
