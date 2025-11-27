package MVRTypes

type UserData struct {
	Data []Data
}

type Data struct {
	Provider string
	Version  string
}
