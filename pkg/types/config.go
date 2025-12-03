package MVRTypes

type ModelConfig struct {
	RenderOnlyAddressedFixture bool
	ExcludeUUIDs               map[string]struct{} // meshes (including childs) of these uuids will be excluded
}
