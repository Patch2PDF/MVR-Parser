package MVRXML

type UserData struct {
	Data []Data `xml:"Data"`
}

type Data struct {
	Provider string `xml:"provider,attr"`
	Version  string `xml:"ver,attr"`
}
