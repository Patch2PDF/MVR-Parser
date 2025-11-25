package MVRXML

type GroupObject struct {
	UUID   string  `xml:"uuid,attr"`
	Name   string  `xml:"name,attr"`
	Matrix *Matrix `xml:"Matrix"`
	Class  *string `xml:"Classing"`
	ChildList
}
