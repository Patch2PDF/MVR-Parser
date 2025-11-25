package MVRXML

type Addresses struct {
	Addresses []*Address `xml:"Address"`
	Networks  []*Network `xml:"Network"`
}

type Address struct {
	Break int    `xml:"break,attr"`
	Value string `xml:",innerxml"` // needs to be converted into numeric, with seperate universe field at another time
}

type Network struct {
	Geometry   string  `xml:"geometry,attr"`
	IPv4       *IPv4   `xml:"ipv4,attr"`
	SubNetMask *IPv4   `xml:"subnetmask,attr"`
	IPv6       *IPv6   `xml:"ipv6,attr"`
	DHCP       bool    `xml:"dhcp,attr"`
	Hostname   *string `xml:"hostname,attr"`
}
