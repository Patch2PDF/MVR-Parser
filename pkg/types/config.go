package MVRTypes

type ModelConfig struct {
	Global     GlobalModelConfig
	Individual map[string]ModelNodeConfig // configure by Node (fixture, ...) UUID (also applies to children)
}

type GlobalModelConfig struct {
	RenderOnlyAddressedFixture bool
}

var FalsePtr = false
var TruePtr = true

func (a GlobalModelConfig) asNodeConfig() ModelNodeConfig {
	return ModelNodeConfig{
		RenderOnlyAddressedFixture: &a.RenderOnlyAddressedFixture,
		Exclude:                    &FalsePtr,
	}
}

type ModelNodeConfig struct {
	RenderOnlyAddressedFixture *bool
	Exclude                    *bool
}

type MVRParserConfig struct {
	MeshHandling      int
	ReadThumbnail     bool
	GDTFParserWorkers int
	StageMeshWorkers  int
	ModelConfig       ModelConfig
}
