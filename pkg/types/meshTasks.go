package MVRTypes

import (
	"sync"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/v2/pkg/MeshTypes"
	GDTFTypes "github.com/Patch2PDF/GDTF-Parser/pkg/types"
)

type MeshTasks = []MeshTransformationTask

type MeshTransformationTask struct {
	Matrix MeshTypes.Matrix
	Mesh   *MeshTypes.Mesh
}

type ParentMeshConfig struct {
	Transformation MeshTypes.Matrix
	ModelConfig    ModelNodeConfig
}

type MeshTaskCreator interface {
	GenerateMesh(meshTasks *MeshTasks, stageModel *StageModel, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig)
}

func getConfigOverrides(modelConfig ModelConfig, parentMeshConfig ParentMeshConfig, uuid string) ModelNodeConfig {
	configOverrides := parentMeshConfig.ModelConfig
	if _, found := modelConfig.Individual[uuid]; found {
		temp := modelConfig.Individual[uuid]
		if temp.Exclude != nil {
			configOverrides.Exclude = temp.Exclude
		}
		if temp.RenderOnlyAddressedFixture != nil {
			configOverrides.RenderOnlyAddressedFixture = temp.RenderOnlyAddressedFixture
		}
	}
	return configOverrides
}

func GenerateMeshes[T MeshTaskCreator](objects []T, meshTasks *MeshTasks, stageModel *StageModel, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	for _, obj := range objects {
		obj.GenerateMesh(meshTasks, stageModel, modelConfig, parentMeshConfig)
	}
}

func (a *GeneralSceneDescription) GenerateMeshes(meshTasks *MeshTasks, modelConfig ModelConfig) {
	if a.StageModel == nil {
		a.StageModel = &StageModel{}
	}
	for _, layer := range a.Scene.Layers {
		layer.GenerateMesh(meshTasks, a.StageModel, modelConfig, ParentMeshConfig{
			Transformation: MeshTypes.IdentityMatrix(),
			ModelConfig:    modelConfig.Global.asNodeConfig(),
		})
	}
}

func (a *ChildList) GenerateMesh(meshTasks *MeshTasks, stageModel *StageModel, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	GenerateMeshes(a.SceneObjects, meshTasks, stageModel, modelConfig, parentMeshConfig)
	GenerateMeshes(a.GroupObjects, meshTasks, stageModel, modelConfig, parentMeshConfig)
	GenerateMeshes(a.FocusPoints, meshTasks, stageModel, modelConfig, parentMeshConfig)
	GenerateMeshes(a.Fixtures, meshTasks, stageModel, modelConfig, parentMeshConfig)
	GenerateMeshes(a.Supports, meshTasks, stageModel, modelConfig, parentMeshConfig)
	GenerateMeshes(a.Trusses, meshTasks, stageModel, modelConfig, parentMeshConfig)
	GenerateMeshes(a.VideoScreens, meshTasks, stageModel, modelConfig, parentMeshConfig)
	GenerateMeshes(a.Projectors, meshTasks, stageModel, modelConfig, parentMeshConfig)
}

func (obj *GroupObject) GenerateMesh(meshTasks *MeshTasks, stageModel *StageModel, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)

	obj.ChildList.GenerateMesh(meshTasks, stageModel, modelConfig, ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	})
}

func (obj *SceneObject) GenerateMesh(meshTasks *MeshTasks, stageModel *StageModel, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	if obj.GDTFSpec.Ptr != nil {
		model := SceneObjectModel{
			SceneObject:          obj,
			TransformationMatrix: matrix,
			MeshModel:            make([]GDTFTypes.MeshModel, 0, len(obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels)),
		}
		for _, part := range obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels {
			copy := part.Copy()
			copy.Mesh.RotateAndTranslate(matrix)
			model.MeshModel = append(model.MeshModel, copy)
		}
		stageModel.SceneObjectModels = append(stageModel.SceneObjectModels, model)
	}

	parentConf := ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	}

	obj.Geometries.GenerateMesh(meshTasks, stageModel, modelConfig, parentConf)

	obj.ChildList.GenerateMesh(meshTasks, stageModel, modelConfig, ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	})
}

func (obj *FocusPoint) GenerateMesh(meshTasks *MeshTasks, stageModel *StageModel, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	obj.Geometries.GenerateMesh(meshTasks, stageModel, modelConfig, ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	})
}

func (obj *Fixture) GenerateMesh(meshTasks *MeshTasks, stageModel *StageModel, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}
	if (config.RenderOnlyAddressedFixture != nil && *config.RenderOnlyAddressedFixture) && (obj.Addresses == nil || len(obj.Addresses.Addresses) == 0) {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	if obj.GDTFSpec.Ptr != nil {
		model := FixtureModel{
			Fixture:              obj,
			TransformationMatrix: matrix,
			MeshModel:            make([]GDTFTypes.MeshModel, 0, len(obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels)),
		}
		for _, part := range obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels {
			copy := part.Copy()
			copy.Mesh.RotateAndTranslate(matrix)
			model.MeshModel = append(model.MeshModel, copy)
		}
		stageModel.FixtureModels = append(stageModel.FixtureModels, model)
	}

	obj.ChildList.GenerateMesh(meshTasks, stageModel, modelConfig, ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	})
}

func (obj *Support) GenerateMesh(meshTasks *MeshTasks, stageModel *StageModel, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	if obj.GDTFSpec.Ptr != nil {
		model := SupportModel{
			Support:              obj,
			TransformationMatrix: matrix,
			MeshModel:            make([]GDTFTypes.MeshModel, 0, len(obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels)),
		}
		for _, part := range obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels {
			copy := part.Copy()
			copy.Mesh.RotateAndTranslate(matrix)
			model.MeshModel = append(model.MeshModel, copy)
		}
		stageModel.SupportModels = append(stageModel.SupportModels, model)
	}

	parentConf := ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	}

	obj.Geometries.GenerateMesh(meshTasks, stageModel, modelConfig, parentConf)

	obj.ChildList.GenerateMesh(meshTasks, stageModel, modelConfig, parentConf)
}

func (obj *Truss) GenerateMesh(meshTasks *MeshTasks, stageModel *StageModel, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	if obj.GDTFSpec.Ptr != nil {
		model := TrussModel{
			Truss:                obj,
			TransformationMatrix: matrix,
			MeshModel:            make([]GDTFTypes.MeshModel, 0, len(obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels)),
		}
		for _, part := range obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels {
			copy := part.Copy()
			copy.Mesh.RotateAndTranslate(matrix)
			model.MeshModel = append(model.MeshModel, copy)
		}
		stageModel.TrussModels = append(stageModel.TrussModels, model)
	}

	parentConf := ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	}

	obj.Geometries.GenerateMesh(meshTasks, stageModel, modelConfig, parentConf)

	obj.ChildList.GenerateMesh(meshTasks, stageModel, modelConfig, parentConf)
}

func (obj *VideoScreen) GenerateMesh(meshTasks *MeshTasks, stageModel *StageModel, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	if obj.GDTFSpec.Ptr != nil {
		model := VideoScreenModel{
			VideoScreen:          obj,
			TransformationMatrix: matrix,
			MeshModel:            make([]GDTFTypes.MeshModel, 0, len(obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels)),
		}
		for _, part := range obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels {
			copy := part.Copy()
			copy.Mesh.RotateAndTranslate(matrix)
			model.MeshModel = append(model.MeshModel, copy)
		}
		stageModel.VideoScreenModels = append(stageModel.VideoScreenModels, model)
	}

	parentConf := ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	}

	obj.Geometries.GenerateMesh(meshTasks, stageModel, modelConfig, parentConf)

	obj.ChildList.GenerateMesh(meshTasks, stageModel, modelConfig, parentConf)
}

func (obj *Projector) GenerateMesh(meshTasks *MeshTasks, stageModel *StageModel, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	if obj.GDTFSpec.Ptr != nil {
		model := ProjectorModel{
			Projector:            obj,
			TransformationMatrix: matrix,
			MeshModel:            make([]GDTFTypes.MeshModel, 0, len(obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels)),
		}
		for _, part := range obj.GDTFSpec.Ptr.Data.FixtureType.DMXModes[obj.GDTFMode].MeshModels {
			copy := part.Copy()
			copy.Mesh.RotateAndTranslate(matrix)
			model.MeshModel = append(model.MeshModel, copy)
		}
		stageModel.ProjectorModels = append(stageModel.ProjectorModels, model)
	}

	parentConf := ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	}

	obj.Geometries.GenerateMesh(meshTasks, stageModel, modelConfig, parentConf)

	obj.ChildList.GenerateMesh(meshTasks, stageModel, modelConfig, parentConf)
}

func (obj *Geometries) GenerateMesh(meshTasks *MeshTasks, stageModel *StageModel, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	for _, element := range obj.Geometry3D {
		matrix := parentMeshConfig.Transformation.Mul(element.Matrix)
		*meshTasks = append(*meshTasks, MeshTransformationTask{
			Matrix: matrix,
			Mesh:   element.Mesh,
		})
	}
	for _, element := range obj.Symbol {
		matrix := parentMeshConfig.Transformation.Mul(element.Matrix)
		element.GenerateMesh(meshTasks, stageModel, modelConfig, ParentMeshConfig{
			Transformation: matrix,
			ModelConfig:    parentMeshConfig.ModelConfig,
		})
	}
}

func (a *Symbol) GenerateMesh(meshTasks *MeshTasks, stageModel *StageModel, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, a.UUID)
	if config.Exclude != nil && *config.Exclude {
		return
	}
	if a.SymDef.Ptr != nil {
		matrix := parentMeshConfig.Transformation.Mul(a.Matrix)
		a.SymDef.Ptr.Geometries.GenerateMesh(meshTasks, stageModel, modelConfig, ParentMeshConfig{
			Transformation: matrix,
			ModelConfig:    config,
		})
	}
}

func meshTaskWorker(jobs <-chan MeshTransformationTask, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		j.Mesh.RotateAndTranslate(j.Matrix)
		// mesh := j.Mesh.Copy()
		// mesh.RotateAndTranslate(j.Matrix)
		// *j.Mesh = mesh
	}
}

func CompleteMeshTasks(meshTasks *MeshTasks, config MVRParserConfig) {

	var numWorkers = config.StageMeshWorkers
	jobs := make(chan MeshTransformationTask, len(*meshTasks))

	var wg sync.WaitGroup

	for range numWorkers {
		wg.Add(1)
		go meshTaskWorker(jobs, &wg)
	}

	for _, t := range *meshTasks {
		jobs <- t
	}
	close(jobs)

	wg.Wait()

	// return
}
