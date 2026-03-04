package MVRTypes

import (
	"archive/zip"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/v2/pkg/MeshTypes"
)

type FocusPoint struct {
	UUID       string
	Name       string
	Matrix     MeshTypes.Matrix
	Class      NodeReference[Class]
	Model      FocusPointModel
	Geometries *Geometries
}

func (a *FocusPoint) CreateReferencePointer(refPointers *ReferencePointers) {
	refPointers.FoucsPoints[a.UUID] = a
}

func (a *FocusPoint) ResolveReference(refPointers *ReferencePointers) {
	if a.Class.String != nil {
		a.Class.Ptr = refPointers.Classes[*a.Class.String]
	}
	a.Geometries.ResolveReference(refPointers)
}

func (a *FocusPoint) ReadMesh(fileMap map[string]*zip.File) error {
	return a.Geometries.ReadMesh(fileMap)
}

func (a *FocusPoint) addNodeModelsToStageModel(stageModel *StageModel, modelConfig ModelConfig, parentConfig ModelNodeConfig) {
	config := getConfigOverrides(modelConfig, parentConfig, a.UUID)

	if config.Exclude == nil || !(*config.Exclude) {
		stageModel.FocusPointModels = append(stageModel.FocusPointModels, a.Model)
	}
}
