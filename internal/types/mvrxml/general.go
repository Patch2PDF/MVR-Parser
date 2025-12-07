package MVRXML

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
	GDTFReader "github.com/Patch2PDF/MVR-Parser/internal/gdtfreader"
)

type Matrix [4][3]float64

func (dest *Matrix) UnmarshalXMLAttr(attr xml.Attr) error {
	rows := strings.Split(strings.Trim(attr.Value, "{}"), "}{")
	if len(rows) != 4 {
		return fmt.Errorf("invalid structure for Matrix")
	}
	for index, row := range rows {
		columns := strings.Split(row, ",")
		if len(columns) != 3 {
			return fmt.Errorf("invalid structure for Matrix")
		}
		for column_index, column_value := range columns {
			value, err := strconv.ParseFloat(column_value, 64)
			if err != nil {
				return err
			}
			dest[index][column_index] = value
		}
	}
	return nil
}

func (m *Matrix) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}

	// reuse your parsing logic on content (trim spaces, braces)
	content = strings.TrimSpace(content)
	content = strings.Trim(content, "{}")
	rows := strings.Split(content, "}{")
	if len(rows) != 4 {
		return fmt.Errorf("invalid structure for Matrix: %q", content)
	}
	for i, row := range rows {
		cols := strings.Split(row, ",")
		if len(cols) != 3 {
			return fmt.Errorf("invalid structure for Matrix: %q", content)
		}
		for j, cv := range cols {
			v, err := strconv.ParseFloat(strings.TrimSpace(cv), 64)
			if err != nil {
				return err
			}
			m[i][j] = v
		}
	}
	return nil
}

func (m *Matrix) ToMeshMatrix() MeshTypes.Matrix {
	if m == nil {
		return MeshTypes.IdentityMatrix()
	}
	return MeshTypes.Matrix{
		X00: m[0][0], X01: m[1][0], X02: m[2][0], X03: m[3][0] / 1000,
		X10: m[0][1], X11: m[1][1], X12: m[2][1], X13: m[3][1] / 1000,
		X20: m[0][2], X21: m[1][2], X22: m[2][2], X23: m[3][2] / 1000,
		X30: 0, X31: 0, X32: 0, X33: 0,
	}
}

type fileName = string

// format xyY (X,Y,Y2)
type ColorCIE struct {
	X  float32
	Y  float32
	Y2 float32
}

func (dest *ColorCIE) UnmarshalXMLAttr(attr xml.Attr) error {
	frags := strings.Split(attr.Value, ",")
	if len(frags) != 3 {
		return fmt.Errorf("invalid structure for ColorCIE")
	}
	value, err := strconv.ParseFloat(frags[0], 32)
	if err != nil {
		return err
	}
	dest.X = float32(value)
	value, err = strconv.ParseFloat(frags[1], 32)
	if err != nil {
		return err
	}
	dest.Y = float32(value)
	value, err = strconv.ParseFloat(frags[2], 32)
	if err != nil {
		return err
	}
	dest.Y2 = float32(value)
	return nil
}

func (c *ColorCIE) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}

	frags := strings.Split(strings.TrimSpace(content), ",")
	if len(frags) != 3 {
		return fmt.Errorf("invalid structure for ColorCIE: %q", content)
	}

	// Parse X
	v, err := strconv.ParseFloat(strings.TrimSpace(frags[0]), 32)
	if err != nil {
		return err
	}
	c.X = float32(v)

	// Parse Y
	v, err = strconv.ParseFloat(strings.TrimSpace(frags[1]), 32)
	if err != nil {
		return err
	}
	c.Y = float32(v)

	// Parse Y2
	v, err = strconv.ParseFloat(strings.TrimSpace(frags[2]), 32)
	if err != nil {
		return err
	}
	c.Y2 = float32(v)

	return nil
}

type IPv4 = string

type IPv6 = string

type Vector = string

type ConvertToDestinationStruct[T any] interface {
	Parse(config ParseConfigData) T
}

type ConvertToDestinationMapStruct[T any] interface {
	ConvertToDestinationStruct[T]
	ParseKey() string
}

func ParseList[Source ConvertToDestinationStruct[Destination], Destination any](config ParseConfigData, source *[]Source) []Destination {
	if source == nil {
		return nil
	}
	var destination []Destination = make([]Destination, len(*source))
	for index, element := range *source {
		parsedElement := element.Parse(config)
		destination[index] = parsedElement
	}
	return destination
}

func ParseMap[Source ConvertToDestinationMapStruct[Destination], Destination any](config ParseConfigData, source *[]Source) map[string]*Destination {
	destination := make(map[string]*Destination)
	for _, element := range *source {
		parsedElement := element.Parse(config)
		destination[element.ParseKey()] = &parsedElement
	}
	return destination
}

type ParseConfigData struct {
	GDTFTaskMap *map[string]*GDTFReader.GDTFTask
}
