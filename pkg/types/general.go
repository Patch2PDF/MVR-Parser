package MVRTypes

import (
	"strconv"
	"strings"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
)

type Matrix = MeshTypes.Matrix

type fileName = string

// format xyY (X,Y,Y2)
type ColorCIE struct {
	X  float32
	Y  float32
	Y2 float32
}

type IPv4 = string

type IPv6 = string

type Vector = string

// NOTE: values are 0 indexed
type DMXAddress struct {
	Address  int16 // 0 indexed
	Universe int   // 0 indexed
}

func GetDMXAddress(value string) (*DMXAddress, error) {
	b := DMXAddress{}
	if strings.Contains(value, ".") {
		frags := strings.Split(value, ".")
		value, err := strconv.ParseInt(frags[0], 10, 0)
		if err != nil {
			return nil, err
		}
		b.Universe = int(value)

		value, err = strconv.ParseInt(frags[1], 10, 16)
		if err != nil {
			return nil, err
		}
		b.Address = int16(value)
	} else {
		absolute, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		b.Address = int16((absolute - 1) % 512)
		b.Universe = int((absolute - 1) / 512)
	}
	return &b, nil
}
