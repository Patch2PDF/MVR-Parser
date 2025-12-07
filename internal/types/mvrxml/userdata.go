package MVRXML

import MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"

type UserData struct {
	Data []Data `xml:"Data"`
}

func (u *UserData) Parse(config ParseConfigData) *MVRTypes.UserData {
	return &MVRTypes.UserData{
		Data: ParseList(config, &u.Data),
	}
}

type Data struct {
	Provider string `xml:"provider,attr"`
	Version  string `xml:"ver,attr"`
}

func (u Data) Parse(config ParseConfigData) MVRTypes.Data {
	return MVRTypes.Data(u)
}
