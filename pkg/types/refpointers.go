package MVRTypes

import "sync"

type ReferencePointers struct {
	GDTFSpecs          map[string]*GDTF // handled seperately
	Classes            map[string]*Class
	Positions          map[string]*Position
	SymDefs            map[string]*SymDef
	MappingDefinitions map[string]*MappingDefinition
	FoucsPoints        map[string]*FocusPoint
	Object             map[string]*any
}

var refPointers ReferencePointers = ReferencePointers{
	GDTFSpecs:          map[string]*GDTF{}, // handled seperately
	Classes:            map[string]*Class{},
	Positions:          map[string]*Position{},
	SymDefs:            map[string]*SymDef{},
	MappingDefinitions: map[string]*MappingDefinition{},
	FoucsPoints:        map[string]*FocusPoint{},
	Object:             map[string]*any{},
}

type ReferenceCreation interface {
	CreateReferencePointer()
}

func CreateReferencePointers[T ReferenceCreation](source *[]T) {
	for _, element := range *source {
		element.CreateReferencePointer()
	}
}

func CreateReferencePointersMap[T ReferenceCreation](source *map[string]T) {
	for _, element := range *source {
		element.CreateReferencePointer()
	}
}

type GeometryReferenceCreation interface {
	CreateGeometryReferencePointer(parentPrefix string)
}

func CreateGeometryReferencePointers[T GeometryReferenceCreation](source *[]T, parentPrefix string) {
	for _, element := range *source {
		element.CreateGeometryReferencePointer(parentPrefix)
	}
}

type ReferenceResolver interface {
	ResolveReference()
}

func ResolveReferences[T ReferenceResolver](source *[]T) {
	if source == nil {
		return
	}
	for i := range *source {
		(*source)[i].ResolveReference()
	}
}

func ResolveReferencesMap[T ReferenceResolver](source *map[string]T) {
	if source == nil {
		return
	}
	for i := range *source {
		(*source)[i].ResolveReference()
	}
}

var gdtfLock = &sync.Mutex{}

func AddGDTFPointer(name string, a *GDTF) {
	gdtfLock.Lock()
	refPointers.GDTFSpecs[name] = a
	gdtfLock.Unlock()
}

func GetGDTFPointer(name string) *GDTF {
	gdtfLock.Lock()
	ptr := refPointers.GDTFSpecs[name]
	gdtfLock.Unlock()
	return ptr
}
