package MVRTypes

const (
	IgnoreMeshes         = 0 // ignore meshes entirely (leave nil)
	ReadMeshesIntoModels = 1 // generate individual model part meshes
	BuildFixtureModels   = 2 // assemble single mesh with correct rotations per fixture type
	BuildStageModel      = 3 // assemble
)
