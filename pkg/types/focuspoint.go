package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type FocusPoint struct {
	UUID       string
	Name       string
	Matrix     MeshTypes.Matrix
	Class      *string
	Geometries *Geometries
}
