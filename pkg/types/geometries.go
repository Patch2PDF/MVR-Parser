package MVRTypes

import (
	"archive/zip"
	"path"

	GDTFMeshReader "github.com/Patch2PDF/GDTF-Mesh-Reader"
	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
)

type Geometries struct {
	Geometry3D []*Geometry3D
	Symbol     []*Symbol
}

func (a *Geometries) ResolveReference(refPointers *ReferencePointers) {
	ResolveReferences(refPointers, &a.Symbol)
}

type Geometry3D struct {
	FileName fileName
	Matrix   MeshTypes.Matrix
	Mesh     *MeshTypes.Mesh
}

type Symbol struct {
	UUID   string
	SymDef NodeReference[SymDef]
	Matrix MeshTypes.Matrix
}

func (a *Symbol) ResolveReference(refPointers *ReferencePointers) {
	if a.SymDef.String != nil {
		a.SymDef.Ptr = refPointers.SymDefs[*a.SymDef.String]
	}
}

func (a *Geometries) ReadMesh(fileMap map[string]*zip.File) error {
	return ReadMeshes(a.Geometry3D, fileMap)
}

func (a *Geometry3D) ReadMesh(fileMap map[string]*zip.File) error {
	if fileMap[a.FileName] != nil {
		file, err := fileMap[a.FileName].Open()
		if err != nil {
			return err
		}
		conf := GDTFMeshReader.ModelReaderConf{
			File:     file,
			Filename: &a.FileName,
		}
		mesh, err := GDTFMeshReader.GetModel(conf, nil) // keep original model size
		if err != nil {
			return err
		}
		// correct 3ds files being in mm, according to mvr spec
		if path.Ext(a.FileName) == ".3ds" {
			mesh.Scale(MeshTypes.Vector{X: 1.0 / 1000, Y: 1.0 / 1000, Z: 1.0 / 1000})
		}
		a.Mesh = mesh
	}
	return nil
}
