package MVRTypes

type GeneralSceneDescription struct {
	VersionMajor    int
	VersionMinor    int
	Provider        string
	ProviderVersion string
	UserData        *UserData
	Scene           *Scene
	StageModel      *StageModel
}

func (a *GeneralSceneDescription) CreateReferencePointer(refPointers *ReferencePointers) {
	a.Scene.CreateReferencePointer(refPointers)
}

func (a *GeneralSceneDescription) ResolveReference(refPointers *ReferencePointers) {
	a.Scene.ResolveReference(refPointers)
}
