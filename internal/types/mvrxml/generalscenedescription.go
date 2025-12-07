package MVRXML

import MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"

type GeneralSceneDescription struct {
	VersionMajor    int       `xml:"verMajor,attr"`
	VersionMinor    int       `xml:"verMinor,attr"`
	Provider        string    `xml:"provider,attr"`
	ProviderVersion string    `xml:"providerVersion,attr"`
	UserData        *UserData `xml:"UserData"`
	Scene           Scene     `xml:"Scene"`
}

func (obj *GeneralSceneDescription) Parse(config ParseConfigData) *MVRTypes.GeneralSceneDescription {
	return &MVRTypes.GeneralSceneDescription{
		VersionMajor:    obj.VersionMajor,
		VersionMinor:    obj.VersionMinor,
		Provider:        obj.Provider,
		ProviderVersion: obj.ProviderVersion,
		UserData:        obj.UserData.Parse(config),
		Scene:           obj.Scene.Parse(config),
	}
}
