package MVRTypes

type ModelConfig struct {
	Global      GlobalModelConfig
	Individual  map[string]ModelNodeConfig // configure by Node (fixture, ...) UUID (also applies to children)
	ClassConfig ModelClassConfig
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

// Note: include has precedence over exclude
type ModelClassConfig struct {
	Excludes map[string]struct{} // specify if you want to exclude these Class UUIDs from the model
	Includes map[string]struct{} // specify if only want to include these Class UUIDs in the model
}

func checkShouldIncludeClassInModel(classConfig ModelClassConfig, classID *string, parentClassID *string) (result bool, newParentClassID *string) {
	if classID == nil {
		if parentClassID == nil {
			return len(classConfig.Includes) <= 0, nil
		} else {
			classID = parentClassID
		}
	}
	if len(classConfig.Includes) > 0 {
		if _, ok := classConfig.Includes[*classID]; ok {
			return true, classID
		}
	} else {
		if _, ok := classConfig.Excludes[*classID]; !ok {
			return true, classID
		}
	}
	return false, classID
}

type MVRParserConfig struct {
	MeshHandling      int
	ReadThumbnail     bool
	GDTFParserWorkers int
	StageMeshWorkers  int
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

type parentNodeParameters struct {
	classID *string
}
