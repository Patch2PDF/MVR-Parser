package MVRXML

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
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
