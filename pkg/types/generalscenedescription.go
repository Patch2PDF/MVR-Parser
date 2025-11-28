package MVRTypes

type GeneralSceneDescription struct {
	VersionMajor    int
	VersionMinor    int
	Provider        string
	ProviderVersion string
	UserData        *UserData
	Scene           *Scene
}

func (a *GeneralSceneDescription) CreateReferencePointer() {
	a.Scene.CreateReferencePointer()
}

func (a *GeneralSceneDescription) ResolveReference() {
	a.Scene.ResolveReference()
}
