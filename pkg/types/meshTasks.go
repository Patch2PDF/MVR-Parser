package MVRTypes

import (
	"sync"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
)

type MeshTasks = []MeshTransformationTask

type MeshTransformationTask struct {
	Matrix       MeshTypes.Matrix
	OriginalMesh *MeshTypes.Mesh
}

type ParentMeshConfig struct {
	Transformation MeshTypes.Matrix
	ModelConfig    ModelNodeConfig
}

type MeshTaskCreator interface {
	CreateMeshTask(meshTasks *MeshTasks, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig)
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

func CreateMeshTasks[T MeshTaskCreator](objects []T, meshTasks *MeshTasks, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	for _, obj := range objects {
		obj.CreateMeshTask(meshTasks, modelConfig, parentMeshConfig)
	}
}

func (a *GeneralSceneDescription) CreateMeshTasks(meshTasks *MeshTasks, modelConfig ModelConfig) {
	for _, layer := range a.Scene.Layers {
		layer.CreateMeshTask(meshTasks, modelConfig, ParentMeshConfig{
			Transformation: MeshTypes.IdentityMatrix(),
			ModelConfig:    modelConfig.Global.asNodeConfig(),
		})
	}
}

func (a *ChildList) CreateMeshTask(meshTasks *MeshTasks, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	CreateMeshTasks(a.SceneObjects, meshTasks, modelConfig, parentMeshConfig)
	CreateMeshTasks(a.GroupObjects, meshTasks, modelConfig, parentMeshConfig)
	CreateMeshTasks(a.FocusPoints, meshTasks, modelConfig, parentMeshConfig)
	CreateMeshTasks(a.Fixtures, meshTasks, modelConfig, parentMeshConfig)
	CreateMeshTasks(a.Supports, meshTasks, modelConfig, parentMeshConfig)
	CreateMeshTasks(a.Trusses, meshTasks, modelConfig, parentMeshConfig)
	CreateMeshTasks(a.VideoScreens, meshTasks, modelConfig, parentMeshConfig)
	CreateMeshTasks(a.Projectors, meshTasks, modelConfig, parentMeshConfig)
}

func (obj *GroupObject) CreateMeshTask(meshTasks *MeshTasks, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)

	obj.ChildList.CreateMeshTask(meshTasks, modelConfig, ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	})
}

func (obj *SceneObject) CreateMeshTask(meshTasks *MeshTasks, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	*meshTasks = append(*meshTasks, MeshTransformationTask{
		Matrix:       matrix,
		OriginalMesh: obj.GDTFSpec.Ptr.Meshes[obj.GDTFMode],
	})

	parentConf := ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	}

	obj.Geometries.CreateMeshTask(meshTasks, modelConfig, parentConf)

	obj.ChildList.CreateMeshTask(meshTasks, modelConfig, ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	})
}

func (obj *FocusPoint) CreateMeshTask(meshTasks *MeshTasks, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	obj.Geometries.CreateMeshTask(meshTasks, modelConfig, ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	})
}

func (obj *Fixture) CreateMeshTask(meshTasks *MeshTasks, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}
	if (config.RenderOnlyAddressedFixture != nil && *config.RenderOnlyAddressedFixture) && (obj.Addresses == nil || len(obj.Addresses.Addresses) == 0) {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	*meshTasks = append(*meshTasks, MeshTransformationTask{
		Matrix:       matrix,
		OriginalMesh: obj.GDTFSpec.Ptr.Meshes[obj.GDTFMode],
	})

	obj.ChildList.CreateMeshTask(meshTasks, modelConfig, ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	})
}

func (obj *Support) CreateMeshTask(meshTasks *MeshTasks, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	*meshTasks = append(*meshTasks, MeshTransformationTask{
		Matrix:       matrix,
		OriginalMesh: obj.GDTFSpec.Ptr.Meshes[obj.GDTFMode],
	})

	parentConf := ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	}

	obj.Geometries.CreateMeshTask(meshTasks, modelConfig, parentConf)

	obj.ChildList.CreateMeshTask(meshTasks, modelConfig, parentConf)
}

func (obj *Truss) CreateMeshTask(meshTasks *MeshTasks, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	*meshTasks = append(*meshTasks, MeshTransformationTask{
		Matrix:       matrix,
		OriginalMesh: obj.GDTFSpec.Ptr.Meshes[obj.GDTFMode],
	})

	parentConf := ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	}

	obj.Geometries.CreateMeshTask(meshTasks, modelConfig, parentConf)

	obj.ChildList.CreateMeshTask(meshTasks, modelConfig, parentConf)
}

func (obj *VideoScreen) CreateMeshTask(meshTasks *MeshTasks, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	*meshTasks = append(*meshTasks, MeshTransformationTask{
		Matrix:       matrix,
		OriginalMesh: obj.GDTFSpec.Ptr.Meshes[obj.GDTFMode],
	})

	parentConf := ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	}

	obj.Geometries.CreateMeshTask(meshTasks, modelConfig, parentConf)

	obj.ChildList.CreateMeshTask(meshTasks, modelConfig, parentConf)
}

func (obj *Projector) CreateMeshTask(meshTasks *MeshTasks, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, obj.UUID)

	if config.Exclude != nil && *config.Exclude {
		return
	}

	matrix := parentMeshConfig.Transformation.Mul(obj.Matrix)
	*meshTasks = append(*meshTasks, MeshTransformationTask{
		Matrix:       matrix,
		OriginalMesh: obj.GDTFSpec.Ptr.Meshes[obj.GDTFMode],
	})

	parentConf := ParentMeshConfig{
		Transformation: matrix,
		ModelConfig:    config,
	}

	obj.Geometries.CreateMeshTask(meshTasks, modelConfig, parentConf)

	obj.ChildList.CreateMeshTask(meshTasks, modelConfig, parentConf)
}

func (obj *Geometries) CreateMeshTask(meshTasks *MeshTasks, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	for _, element := range obj.Geometry3D {
		matrix := parentMeshConfig.Transformation.Mul(element.Matrix)
		*meshTasks = append(*meshTasks, MeshTransformationTask{
			Matrix:       matrix,
			OriginalMesh: element.Mesh,
		})
	}
	for _, element := range obj.Symbol {
		matrix := parentMeshConfig.Transformation.Mul(element.Matrix)
		element.CreateMeshTask(meshTasks, modelConfig, ParentMeshConfig{
			Transformation: matrix,
			ModelConfig:    parentMeshConfig.ModelConfig,
		})
	}
}

func (a *Symbol) CreateMeshTask(meshTasks *MeshTasks, modelConfig ModelConfig, parentMeshConfig ParentMeshConfig) {
	config := getConfigOverrides(modelConfig, parentMeshConfig, a.UUID)
	if config.Exclude != nil && *config.Exclude {
		return
	}
	if a.SymDef.Ptr != nil {
		matrix := parentMeshConfig.Transformation.Mul(a.Matrix)
		a.SymDef.Ptr.Geometries.CreateMeshTask(meshTasks, modelConfig, ParentMeshConfig{
			Transformation: matrix,
			ModelConfig:    config,
		})
	}
}

func meshTaskWorker(jobs <-chan MeshTransformationTask, results chan<- *MeshTypes.Mesh, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		mesh := j.OriginalMesh.Copy()
		mesh.RotateAndTranslate(j.Matrix)
		results <- &mesh
	}
}

func CompleteMeshTasks(meshTasks *MeshTasks, config MVRParserConfig) *MeshTypes.Mesh {

	var numWorkers = config.StageMeshWorkers
	jobs := make(chan MeshTransformationTask, len(*meshTasks))
	results := make(chan *MeshTypes.Mesh, len(*meshTasks))

	var wg sync.WaitGroup

	for range numWorkers {
		wg.Add(1)
		go meshTaskWorker(jobs, results, &wg)
	}

	for _, t := range *meshTasks {
		jobs <- t
	}
	close(jobs)

	wg.Wait()
	close(results)

	// pre allocate length to reduce array resizings
	meshes := make([]*MeshTypes.Mesh, 0, len(results))
	var totalTriangles int = 0

	for mesh := range results {
		totalTriangles += len(mesh.Triangles)
		meshes = append(meshes, mesh)
	}

	result := MeshTypes.Mesh{
		Triangles: make([]*MeshTypes.Triangle, 0, totalTriangles),
	}

	result.Add(meshes...)

	return &result
}
