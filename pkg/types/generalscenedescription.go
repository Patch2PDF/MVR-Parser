package MVRTypes

type GeneralSceneDescription struct {
	VersionMajor    int
	VersionMinor    int
	Provider        string
	ProviderVersion string
	UserData        *UserData
	Scene           *Scene
}
