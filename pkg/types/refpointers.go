package MVRTypes

type ReferencePointers struct {
	GDTFSpecs          map[string]*GDTF // handled seperately
	Classes            map[string]*Class
	Positions          map[string]*Position
	SymDefs            map[string]*SymDef
	MappingDefinitions map[string]*MappingDefinition
	FoucsPoints        map[string]*FocusPoint
	Object             map[string]*any
}

func CreateRefPointersMap() *ReferencePointers {
	return &ReferencePointers{
		GDTFSpecs:          map[string]*GDTF{}, // handled seperately
		Classes:            map[string]*Class{},
		Positions:          map[string]*Position{},
		SymDefs:            map[string]*SymDef{},
		MappingDefinitions: map[string]*MappingDefinition{},
		FoucsPoints:        map[string]*FocusPoint{},
		Object:             map[string]*any{},
	}
}

type ReferenceCreation interface {
	CreateReferencePointer(refPointers *ReferencePointers)
}

func CreateReferencePointers[T ReferenceCreation](refPointers *ReferencePointers, source *[]T) {
	for _, element := range *source {
		element.CreateReferencePointer(refPointers)
	}
}

func CreateReferencePointersMap[T ReferenceCreation](refPointers *ReferencePointers, source *map[string]T) {
	for _, element := range *source {
		element.CreateReferencePointer(refPointers)
	}
}

type GeometryReferenceCreation interface {
	CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string)
}

func CreateGeometryReferencePointers[T GeometryReferenceCreation](refPointers *ReferencePointers, source *[]T, parentPrefix string) {
	for _, element := range *source {
		element.CreateGeometryReferencePointer(refPointers, parentPrefix)
	}
}

type ReferenceResolver interface {
	ResolveReference(refPointers *ReferencePointers)
}

func ResolveReferences[T ReferenceResolver](refPointers *ReferencePointers, source *[]T) {
	if source == nil {
		return
	}
	for i := range *source {
		(*source)[i].ResolveReference(refPointers)
	}
}

func ResolveReferencesMap[T ReferenceResolver](refPointers *ReferencePointers, source *map[string]T) {
	if source == nil {
		return
	}
	for i := range *source {
		(*source)[i].ResolveReference(refPointers)
	}
}

// func AddGDTFPointer(name string, a *GDTF) {
// 	refPointers.GDTFSpecs[name] = a
// }
