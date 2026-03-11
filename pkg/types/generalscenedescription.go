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

	local_config := config.Global.asNodeConfig()

	addNodeModelsToStageModel(a.Scene.Layers, &model, config, local_config, parentNodeParameters{})

	return model
}
