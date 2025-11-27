package MVRTypes

type Addresses struct {
	Addresses []*Address
	Networks  []*Network
}

type Address struct {
	Break int
	Value DMXAddress
}

type Network struct {
	Geometry   string
	IPv4       *IPv4
	SubNetMask *IPv4
	IPv6       *IPv6
	DHCP       bool
	Hostname   *string
}
