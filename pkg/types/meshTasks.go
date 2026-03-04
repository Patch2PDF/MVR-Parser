package MVRTypes

import (
	"github.com/Patch2PDF/GDTF-Mesh-Reader/v2/pkg/MeshTypes"
	GDTFTypes "github.com/Patch2PDF/GDTF-Parser/pkg/types"
)

type ParentMeshConfig struct {
	Transformation MeshTypes.Matrix
}

type MeshTaskCreator interface {
	GenerateMesh(parentMeshConfig ParentMeshConfig)
}

func GenerateMeshes[T MeshTaskCreator](objects []T, parentMeshConfig ParentMeshConfig) {
	for _, obj := range objects {
		obj.GenerateMesh(parentMeshConfig)
	}
}

func (a *GeneralSceneDescription) GenerateMeshes() {
	for _, layer := range a.Scene.Layers {
		layer.GenerateMesh(ParentMeshConfig{
			Transformation: MeshTypes.IdentityMatrix(),
		})
	}
}

func (a *ChildList) GenerateMesh(parentMeshConfig ParentMeshConfig) {
	GenerateMeshes(a.SceneObjects, parentMeshConfig)
	GenerateMeshes(a.GroupObjects, parentMeshConfig)
	GenerateMeshes(a.FocusPoints, parentMeshConfig)
	GenerateMeshes(a.Fixtures, parentMeshConfig)
	GenerateMeshes(a.Supports, parentMeshConfig)
	GenerateMeshes(a.Trusses, parentMeshConfig)
	GenerateMeshes(a.VideoScreens, parentMeshConfig)
	GenerateMeshes(a.Projectors, parentMeshConfig)
}

func (obj *GroupObject) GenerateMesh(parentMeshConfig ParentMeshConfig) {
	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)

	obj.ChildList.GenerateMesh(ParentMeshConfig{
		Transformation: matrix,
	})
}

func (obj *SceneObject) GenerateMesh(parentMeshConfig ParentMeshConfig) {

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)

	model := SceneObjectModel{
		SceneObject:          obj,
		TransformationMatrix: matrix,
	}

	if obj.GDTFSpec.Ptr != nil {
		model.MeshModel = make([]GDTFTypes.MeshModel, 0, len(obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels))

		for _, part := range obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels {
			copy := part.Copy()
			copy.Mesh.RotateAndTranslate(matrix)
			model.MeshModel = append(model.MeshModel, copy)
		}
	}

	parentConf := ParentMeshConfig{
		Transformation: matrix,
	}

	model.Geometries = obj.Geometries.GenerateMeshes(parentConf)

	obj.Model = model

	obj.ChildList.GenerateMesh(ParentMeshConfig{
		Transformation: matrix,
	})
}

func (obj *FocusPoint) GenerateMesh(parentMeshConfig ParentMeshConfig) {

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)

	model := FocusPointModel{
		FocusPoint:           obj,
		TransformationMatrix: matrix,
	}

	model.Geometries = obj.Geometries.GenerateMeshes(ParentMeshConfig{
		Transformation: matrix,
	})

	obj.Model = model
}

func (obj *Fixture) GenerateMesh(parentMeshConfig ParentMeshConfig) {
	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)

	model := FixtureModel{
		Fixture:              obj,
		TransformationMatrix: matrix,
	}

	if obj.GDTFSpec.Ptr != nil {
		model.MeshModel = make([]GDTFTypes.MeshModel, 0, len(obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels))

		for _, part := range obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels {
			copy := part.Copy()
			copy.Mesh.RotateAndTranslate(matrix)
			model.MeshModel = append(model.MeshModel, copy)
		}
	}

	obj.Model = model

	obj.ChildList.GenerateMesh(ParentMeshConfig{
		Transformation: matrix,
	})
}

func (obj *Support) GenerateMesh(parentMeshConfig ParentMeshConfig) {

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)

	model := SupportModel{
		Support:              obj,
		TransformationMatrix: matrix,
	}

	if obj.GDTFSpec.Ptr != nil {
		model.MeshModel = make([]GDTFTypes.MeshModel, 0, len(obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels))

		for _, part := range obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels {
			copy := part.Copy()
			copy.Mesh.RotateAndTranslate(matrix)
			model.MeshModel = append(model.MeshModel, copy)
		}
	}

	parentConf := ParentMeshConfig{
		Transformation: matrix,
	}

	model.Geometries = obj.Geometries.GenerateMeshes(parentConf)

	obj.Model = model

	obj.ChildList.GenerateMesh(parentConf)
}

func (obj *Truss) GenerateMesh(parentMeshConfig ParentMeshConfig) {

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)

	model := TrussModel{
		Truss:                obj,
		TransformationMatrix: matrix,
	}

	if obj.GDTFSpec.Ptr != nil {
		model.MeshModel = make([]GDTFTypes.MeshModel, 0, len(obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels))

		for _, part := range obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels {
			copy := part.Copy()
			copy.Mesh.RotateAndTranslate(matrix)
			model.MeshModel = append(model.MeshModel, copy)
		}
	}

	parentConf := ParentMeshConfig{
		Transformation: matrix,
	}

	model.Geometries = obj.Geometries.GenerateMeshes(parentConf)

	obj.Model = model

	obj.ChildList.GenerateMesh(parentConf)
}

func (obj *VideoScreen) GenerateMesh(parentMeshConfig ParentMeshConfig) {

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)

	model := VideoScreenModel{
		VideoScreen:          obj,
		TransformationMatrix: matrix,
	}

	if obj.GDTFSpec.Ptr != nil {
		model.MeshModel = make([]GDTFTypes.MeshModel, 0, len(obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels))

		for _, part := range obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels {
			copy := part.Copy()
			copy.Mesh.RotateAndTranslate(matrix)
			model.MeshModel = append(model.MeshModel, copy)
		}
	}

	parentConf := ParentMeshConfig{
		Transformation: matrix,
	}

	model.Geometries = obj.Geometries.GenerateMeshes(parentConf)

	obj.Model = model

	obj.ChildList.GenerateMesh(parentConf)
}

func (obj *Projector) GenerateMesh(parentMeshConfig ParentMeshConfig) {

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)

	model := ProjectorModel{
		Projector:            obj,
		TransformationMatrix: matrix,
	}

	if obj.GDTFSpec.Ptr != nil {
		model.MeshModel = make([]GDTFTypes.MeshModel, 0, len(obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels))

		for _, part := range obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels {
			copy := part.Copy()
			copy.Mesh.RotateAndTranslate(matrix)
			model.MeshModel = append(model.MeshModel, copy)
		}
	}

	parentConf := ParentMeshConfig{
		Transformation: matrix,
	}

	model.Geometries = obj.Geometries.GenerateMeshes(parentConf)

	obj.Model = model

	obj.ChildList.GenerateMesh(parentConf)
}

func (obj *Geometries) GenerateMeshes(parentMeshConfig ParentMeshConfig) []MeshTypes.Mesh {
	meshes := make([]MeshTypes.Mesh, 0, len(obj.Geometry3D)) // allocate atleast amount of Geometry3D's, count of Symbol Meshes is unknown
	for _, element := range obj.Geometry3D {
		matrix := parentMeshConfig.Transformation.Mul(element.Matrix)
		temp := element.Mesh.Copy()
		temp.RotateAndTranslate(matrix)
		meshes = append(meshes, temp)
	}
	for _, element := range obj.Symbol {
		matrix := parentMeshConfig.Transformation.Mul(element.Matrix)
		meshes = append(
			meshes,
			element.GenerateMeshes(ParentMeshConfig{
				Transformation: matrix,
			})...,
		)
	}
	return meshes
}

func (a *Symbol) GenerateMeshes(parentMeshConfig ParentMeshConfig) []MeshTypes.Mesh {
	if a.SymDef.Ptr != nil {
		matrix := parentMeshConfig.Transformation.Mul(a.Matrix)
		return a.SymDef.Ptr.Geometries.GenerateMeshes(ParentMeshConfig{
			Transformation: matrix,
		})
	}
	return []MeshTypes.Mesh{}
}
