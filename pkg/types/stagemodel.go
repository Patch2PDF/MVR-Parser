package MVRTypes

import (
	"github.com/Patch2PDF/GDTF-Mesh-Reader/v2/pkg/MeshTypes"
	GDTFTypes "github.com/Patch2PDF/GDTF-Parser/pkg/types"
)

type DeepCopy[T any] interface {
	Copy() T
}

func CopySlice[Source DeepCopy[Destination], Destination any](source []Source) []Destination {
	if source == nil {
		return []Destination{}
	}
	var destination []Destination = make([]Destination, len(source))
	for index, element := range source {
		destination[index] = element.Copy()
	}
	return destination
}

func CopyMeshSlice(source []MeshTypes.Mesh) []MeshTypes.Mesh {
	slice := make([]MeshTypes.Mesh, 0, len(source))
	for _, mesh := range source {
		slice = append(slice, mesh.Copy())
	}
	return slice
}

type NodeModel interface {
	addNodeModelsToStageModel(stageModel *StageModel, modelConfig ModelConfig, parentConfig ModelNodeConfig, parentParameters parentNodeParameters)
}

func addNodeModelsToStageModel[T NodeModel](source []T, stageModel *StageModel, modelConfig ModelConfig, parentConfig ModelNodeConfig, parentParameters parentNodeParameters) {
	if source == nil {
		return
	}
	for i := range source {
		source[i].addNodeModelsToStageModel(stageModel, modelConfig, parentConfig, parentParameters)
	}
}

type StageModel struct {
	SceneObjectModels []SceneObjectModel
	FocusPointModels  []FocusPointModel
	FixtureModels     []FixtureModel
	SupportModels     []SupportModel
	TrussModels       []TrussModel
	VideoScreenModels []VideoScreenModel
	ProjectorModels   []ProjectorModel
}

func (obj *StageModel) Copy() StageModel {
	return StageModel{
		SceneObjectModels: CopySlice(obj.SceneObjectModels),
		FixtureModels:     CopySlice(obj.FixtureModels),
		SupportModels:     CopySlice(obj.SupportModels),
		TrussModels:       CopySlice(obj.TrussModels),
		VideoScreenModels: CopySlice(obj.VideoScreenModels),
		ProjectorModels:   CopySlice(obj.ProjectorModels),
	}
}

type SceneObjectModel struct {
	SceneObject          *SceneObject
	TransformationMatrix MeshTypes.Matrix
	MeshModel            []GDTFTypes.MeshModel
	Geometries           []MeshTypes.Mesh
}

func (obj SceneObjectModel) Copy() SceneObjectModel {
	return SceneObjectModel{
		SceneObject:          obj.SceneObject,
		TransformationMatrix: obj.TransformationMatrix,
		MeshModel:            CopySlice(obj.MeshModel),
		Geometries:           CopyMeshSlice(obj.Geometries),
	}
}

type FocusPointModel struct {
	FocusPoint           *FocusPoint
	TransformationMatrix MeshTypes.Matrix
	Geometries           []MeshTypes.Mesh
}

func (obj FocusPointModel) Copy() FocusPointModel {
	return FocusPointModel{
		FocusPoint:           obj.FocusPoint,
		TransformationMatrix: obj.TransformationMatrix,
		Geometries:           CopyMeshSlice(obj.Geometries),
	}
}

type FixtureModel struct {
	Fixture              *Fixture
	TransformationMatrix MeshTypes.Matrix
	MeshModel            []GDTFTypes.MeshModel
	Geometries           []MeshTypes.Mesh
}

func (obj FixtureModel) Copy() FixtureModel {
	return FixtureModel{
		Fixture:              obj.Fixture,
		TransformationMatrix: obj.TransformationMatrix,
		MeshModel:            CopySlice(obj.MeshModel),
		Geometries:           CopyMeshSlice(obj.Geometries),
	}
}

type SupportModel struct {
	Support              *Support
	TransformationMatrix MeshTypes.Matrix
	MeshModel            []GDTFTypes.MeshModel
	Geometries           []MeshTypes.Mesh
}

func (obj SupportModel) Copy() SupportModel {
	return SupportModel{
		Support:              obj.Support,
		TransformationMatrix: obj.TransformationMatrix,
		MeshModel:            CopySlice(obj.MeshModel),
		Geometries:           CopyMeshSlice(obj.Geometries),
	}
}

type TrussModel struct {
	Truss                *Truss
	TransformationMatrix MeshTypes.Matrix
	MeshModel            []GDTFTypes.MeshModel
	Geometries           []MeshTypes.Mesh
}

func (obj TrussModel) Copy() TrussModel {
	return TrussModel{
		Truss:                obj.Truss,
		TransformationMatrix: obj.TransformationMatrix,
		MeshModel:            CopySlice(obj.MeshModel),
		Geometries:           CopyMeshSlice(obj.Geometries),
	}
}

type VideoScreenModel struct {
	VideoScreen          *VideoScreen
	TransformationMatrix MeshTypes.Matrix
	MeshModel            []GDTFTypes.MeshModel
	Geometries           []MeshTypes.Mesh
}

func (obj VideoScreenModel) Copy() VideoScreenModel {
	return VideoScreenModel{
		VideoScreen:          obj.VideoScreen,
		TransformationMatrix: obj.TransformationMatrix,
		MeshModel:            CopySlice(obj.MeshModel),
		Geometries:           CopyMeshSlice(obj.Geometries),
	}
}

type ProjectorModel struct {
	Projector            *Projector
	TransformationMatrix MeshTypes.Matrix
	MeshModel            []GDTFTypes.MeshModel
	Geometries           []MeshTypes.Mesh
}

func (obj ProjectorModel) Copy() ProjectorModel {
	return ProjectorModel{
		Projector:            obj.Projector,
		TransformationMatrix: obj.TransformationMatrix,
		MeshModel:            CopySlice(obj.MeshModel),
		Geometries:           CopyMeshSlice(obj.Geometries),
	}
}
