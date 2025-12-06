package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type GeneralSceneDescription struct {
	VersionMajor    int
	VersionMinor    int
	Provider        string
	ProviderVersion string
	UserData        *UserData
	Scene           *Scene
	StageModel      *MeshTypes.Mesh
}

func (a *GeneralSceneDescription) CreateReferencePointer() {
	a.Scene.CreateReferencePointer()
}

func (a *GeneralSceneDescription) ResolveReference() {
	a.Scene.ResolveReference()
}
