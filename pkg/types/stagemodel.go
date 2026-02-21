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

type StageModel struct {
	SceneObjectModels []SceneObjectModel
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
}

func (obj SceneObjectModel) Copy() SceneObjectModel {
	return SceneObjectModel{
		SceneObject:          obj.SceneObject,
		TransformationMatrix: obj.TransformationMatrix,
		MeshModel:            CopySlice(obj.MeshModel),
	}
}

type FixtureModel struct {
	Fixture              *Fixture
	TransformationMatrix MeshTypes.Matrix
	MeshModel            []GDTFTypes.MeshModel
}

func (obj FixtureModel) Copy() FixtureModel {
	return FixtureModel{
		Fixture:              obj.Fixture,
		TransformationMatrix: obj.TransformationMatrix,
		MeshModel:            CopySlice(obj.MeshModel),
	}
}

type SupportModel struct {
	Support              *Support
	TransformationMatrix MeshTypes.Matrix
	MeshModel            []GDTFTypes.MeshModel
}

func (obj SupportModel) Copy() SupportModel {
	return SupportModel{
		Support:              obj.Support,
		TransformationMatrix: obj.TransformationMatrix,
		MeshModel:            CopySlice(obj.MeshModel),
	}
}

type TrussModel struct {
	Truss                *Truss
	TransformationMatrix MeshTypes.Matrix
	MeshModel            []GDTFTypes.MeshModel
}

func (obj TrussModel) Copy() TrussModel {
	return TrussModel{
		Truss:                obj.Truss,
		TransformationMatrix: obj.TransformationMatrix,
		MeshModel:            CopySlice(obj.MeshModel),
	}
}

type VideoScreenModel struct {
	VideoScreen          *VideoScreen
	TransformationMatrix MeshTypes.Matrix
	MeshModel            []GDTFTypes.MeshModel
}

func (obj VideoScreenModel) Copy() VideoScreenModel {
	return VideoScreenModel{
		VideoScreen:          obj.VideoScreen,
		TransformationMatrix: obj.TransformationMatrix,
		MeshModel:            CopySlice(obj.MeshModel),
	}
}

type ProjectorModel struct {
	Projector            *Projector
	TransformationMatrix MeshTypes.Matrix
	MeshModel            []GDTFTypes.MeshModel
}

func (obj ProjectorModel) Copy() ProjectorModel {
	return ProjectorModel{
		Projector:            obj.Projector,
		TransformationMatrix: obj.TransformationMatrix,
		MeshModel:            CopySlice(obj.MeshModel),
	}
}
