package MVRTypes

import (
	"archive/zip"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/v2/pkg/MeshTypes"
)

type GroupObject struct {
	UUID   string
	Name   string
	Matrix MeshTypes.Matrix
	Class  NodeReference[Class]
	ChildList
}

func (a *GroupObject) CreateReferencePointer(refPointers *ReferencePointers) {
	a.ChildList.CreateReferencePointer(refPointers)
}

func (a *GroupObject) ResolveReference(refPointers *ReferencePointers) {
	if a.Class.String != nil {
		a.Class.Ptr = refPointers.Classes[*a.Class.String]
	}
	a.ChildList.ResolveReference(refPointers)
}

func (a *GroupObject) ReadMesh(fileMap map[string]*zip.File) error {
	return a.ChildList.ReadMesh(fileMap)
}

func (a *GroupObject) addNodeModelsToStageModel(stageModel *StageModel, modelConfig ModelConfig, parentConfig ModelNodeConfig, parentParameters parentNodeParameters) {
	config := getConfigOverrides(modelConfig, parentConfig, a.UUID)

	classID := a.Class.String
	if classID == nil {
		classID = parentParameters.classID
	}
	childParameters := parentNodeParameters{
		classID: classID,
	}

	a.ChildList.addNodeModelsToStageModel(stageModel, modelConfig, config, childParameters)
}
