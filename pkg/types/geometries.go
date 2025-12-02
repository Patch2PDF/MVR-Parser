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

func (a *Geometries) ResolveReference() {
	ResolveReferences(&a.Symbol)
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

func (a *Symbol) ResolveReference() {
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

func (a *Geometries) GenerateMesh(parentTransformation MeshTypes.Matrix) *MeshTypes.Mesh {
	newMesh := &MeshTypes.Mesh{}
	for _, element := range a.Geometry3D {
		temp := GenerateMesh(parentTransformation, element.Matrix, element.Mesh)
		// temp, err := GenerateMesh(parentTransformation, element.Matrix, element.Mesh)
		// if err != nil {
		// 	return err
		// }
		newMesh.Add(temp)
	}
	for _, element := range a.Symbol {
		matrix := parentTransformation.Mul(element.Matrix)
		temp := element.GenerateMesh(matrix)
		newMesh.Add(temp)
	}
	return newMesh
}

func (a *Symbol) GenerateMesh(parentTransformation MeshTypes.Matrix) *MeshTypes.Mesh {
	if a.SymDef.Ptr != nil {
		matrix := parentTransformation.Mul(a.Matrix)
		return a.SymDef.Ptr.Geometries.GenerateMesh(matrix)
	}
	return &MeshTypes.Mesh{}
}
