package MVRTypes

type GeneralSceneDescription struct {
	VersionMajor    int
	VersionMinor    int
	Provider        string
	ProviderVersion string
	UserData        *UserData
	Scene           *Scene
}

func (a *GeneralSceneDescription) CreateReferencePointer(refPointers *ReferencePointers) {
	a.Scene.CreateReferencePointer(refPointers)
}

func (a *GeneralSceneDescription) ResolveReference(refPointers *ReferencePointers) {
	a.Scene.ResolveReference(refPointers)
}

func (a *GeneralSceneDescription) GetStageModel(config ModelConfig) StageModel {
	model := StageModel{}
	for _, layer := range a.Scene.Layers {
		layer.addNodeModelsToStageModel(&model, config, ModelNodeConfig{}) // TODO: respect layer config
	}
	return model
}
