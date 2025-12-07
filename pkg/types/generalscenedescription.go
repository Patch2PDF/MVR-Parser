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

func (a *GeneralSceneDescription) CreateReferencePointer(refPointers *ReferencePointers) {
	a.Scene.CreateReferencePointer(refPointers)
}

func (a *GeneralSceneDescription) ResolveReference(refPointers *ReferencePointers) {
	a.Scene.ResolveReference(refPointers)
}
