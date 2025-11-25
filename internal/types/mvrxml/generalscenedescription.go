package MVRXML

type GeneralSceneDescription struct {
	VersionMajor    int       `xml:"verMajor,attr"`
	VersionMinor    int       `xml:"verMinor,attr"`
	Provider        string    `xml:"provider,attr"`
	ProviderVersion string    `xml:"providerVersion,attr"`
	UserData        *UserData `xml:"UserData"`
	Scene           Scene     `xml:"Scene"`
}
