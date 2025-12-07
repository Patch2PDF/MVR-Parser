package MVRXML

import MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"

type Addresses struct {
	Addresses []*Address `xml:"Address"`
	Networks  []*Network `xml:"Network"`
}

func (a *Addresses) Parse(config ParseConfigData) *MVRTypes.Addresses {
	if a == nil {
		return nil
	}
	return &MVRTypes.Addresses{
		Addresses: ParseList(config, &a.Addresses),
		Networks:  ParseList(config, &a.Networks),
	}
}

type Address struct {
	Break int    `xml:"break,attr"`
	Value string `xml:",innerxml"` // needs to be converted into numeric, with seperate universe field at another time
}

func (a *Address) Parse(config ParseConfigData) *MVRTypes.Address {
	address, _ := MVRTypes.GetDMXAddress(a.Value) // TODO: return err
	return &MVRTypes.Address{
		Break: a.Break,
		Value: *address,
	}
}

type Network struct {
	Geometry   string  `xml:"geometry,attr"`
	IPv4       *IPv4   `xml:"ipv4,attr"`
	SubNetMask *IPv4   `xml:"subnetmask,attr"`
	IPv6       *IPv6   `xml:"ipv6,attr"`
	DHCP       bool    `xml:"dhcp,attr"`
	Hostname   *string `xml:"hostname,attr"`
}

func (a *Network) Parse(config ParseConfigData) *MVRTypes.Network {
	return &MVRTypes.Network{
		Geometry:   a.Geometry,
		IPv4:       a.IPv4,
		SubNetMask: a.SubNetMask,
		IPv6:       a.IPv6,
		DHCP:       a.DHCP,
		Hostname:   a.Hostname,
	}
}
