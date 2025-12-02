package MVRTypes

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

func GenerateMesh(parentTransformation MeshTypes.Matrix, ownTransformation MeshTypes.Matrix, mesh *MeshTypes.Mesh) *MeshTypes.Mesh {
	var mesh1 MeshTypes.Mesh
	transformation := parentTransformation.Mul(ownTransformation)
	if mesh != nil {
		mesh1 = mesh.Copy()
		mesh1.RotateAndTranslate(transformation)
	}
	return &mesh1
}

func (a *ChildList) GenerateMesh(parentTransformation MeshTypes.Matrix) *MeshTypes.Mesh {
	newMesh := &MeshTypes.Mesh{}
	for _, obj := range a.SceneObjects {
		// GDTF model
		if obj.GDTFSpec.Ptr != nil {
			temp := GenerateMesh(parentTransformation, obj.Matrix, obj.GDTFSpec.Ptr.Meshes[obj.GDTFMode])
			newMesh.Add(temp)
		}
		matrix := parentTransformation.Mul(obj.Matrix)

		// custom mvr geometries
		geometries := obj.Geometries.GenerateMesh(matrix)
		newMesh.Add(geometries)

		// recursive / render childs
		childs := obj.ChildList.GenerateMesh(matrix)
		newMesh.Add(childs)
	}
	for _, obj := range a.GroupObjects {
		matrix := parentTransformation.Mul(obj.Matrix)
		childs := obj.ChildList.GenerateMesh(matrix)
		newMesh.Add(childs)
	}
	for _, obj := range a.FocusPoints {
		matrix := parentTransformation.Mul(obj.Matrix)

		// custom mvr geometries
		geometries := obj.Geometries.GenerateMesh(matrix)
		newMesh.Add(geometries)
	}
	for _, obj := range a.Fixtures {
		if obj.GDTFSpec.Ptr != nil {
			temp := GenerateMesh(parentTransformation, obj.Matrix, obj.GDTFSpec.Ptr.Meshes[obj.GDTFMode])
			newMesh.Add(temp)
		}
		matrix := parentTransformation.Mul(obj.Matrix)
		childs := obj.ChildList.GenerateMesh(matrix)
		newMesh.Add(childs)
	}
	for _, obj := range a.Supports {
		if obj.GDTFSpec.Ptr != nil {
			temp := GenerateMesh(parentTransformation, obj.Matrix, obj.GDTFSpec.Ptr.Meshes[obj.GDTFMode])
			newMesh.Add(temp)
		}
		matrix := parentTransformation.Mul(obj.Matrix)

		// custom mvr geometries
		geometries := obj.Geometries.GenerateMesh(matrix)
		newMesh.Add(geometries)

		childs := obj.ChildList.GenerateMesh(matrix)
		newMesh.Add(childs)
	}
	for _, obj := range a.Trusses {
		if obj.GDTFSpec.Ptr != nil {
			temp := GenerateMesh(parentTransformation, obj.Matrix, obj.GDTFSpec.Ptr.Meshes[obj.GDTFMode])
			newMesh.Add(temp)
		}
		matrix := parentTransformation.Mul(obj.Matrix)

		// custom mvr geometries
		geometries := obj.Geometries.GenerateMesh(matrix)
		newMesh.Add(geometries)

		childs := obj.ChildList.GenerateMesh(matrix)
		newMesh.Add(childs)
	}
	for _, obj := range a.VideoScreens {
		if obj.GDTFSpec.Ptr != nil {
			temp := GenerateMesh(parentTransformation, obj.Matrix, obj.GDTFSpec.Ptr.Meshes[obj.GDTFMode])
			newMesh.Add(temp)
		}
		matrix := parentTransformation.Mul(obj.Matrix)

		// custom mvr geometries
		geometries := obj.Geometries.GenerateMesh(matrix)
		newMesh.Add(geometries)

		childs := obj.ChildList.GenerateMesh(matrix)
		newMesh.Add(childs)
	}
	for _, obj := range a.Projectors {
		if obj.GDTFSpec.Ptr != nil {
			temp := GenerateMesh(parentTransformation, obj.Matrix, obj.GDTFSpec.Ptr.Meshes[obj.GDTFMode])
			newMesh.Add(temp)
		}
		matrix := parentTransformation.Mul(obj.Matrix)

		// custom mvr geometries
		geometries := obj.Geometries.GenerateMesh(matrix)
		newMesh.Add(geometries)

		childs := obj.ChildList.GenerateMesh(matrix)
		newMesh.Add(childs)
	}
	return newMesh
}
