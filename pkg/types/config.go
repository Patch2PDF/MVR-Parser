package MVRTypes

type ModelConfig struct {
	Global     GlobalModelConfig
	Individual map[string]ModelNodeConfig // configure by Node (fixture, ...) UUID (also applies to children)
}

type GlobalModelConfig struct {
	RenderOnlyAddressedFixture bool
}

func GetBoolPtr(value bool) *bool {
	return &value
}

func (a GlobalModelConfig) asNodeConfig() ModelNodeConfig {
	return ModelNodeConfig{
		RenderOnlyAddressedFixture: &a.RenderOnlyAddressedFixture,
		Exclude:                    GetBoolPtr(false),
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

func getConfigOverrides(modelConfig ModelConfig, parentModelConfig ModelNodeConfig, uuid string) ModelNodeConfig {
	if _, found := modelConfig.Individual[uuid]; found {
		temp := modelConfig.Individual[uuid]
		if temp.Exclude != nil {
			parentModelConfig.Exclude = temp.Exclude
		}
		if temp.RenderOnlyAddressedFixture != nil {
			parentModelConfig.RenderOnlyAddressedFixture = temp.RenderOnlyAddressedFixture
		}
	}
	return parentModelConfig
}
