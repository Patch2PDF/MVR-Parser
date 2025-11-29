package MVRXML

import (
	"archive/zip"

	GDTFReader "github.com/Patch2PDF/MVR-Parser/internal/gdtfreader"
)

type GetGDTF interface {
	GetGDTF(fileMap map[string]*zip.File, modelLevel int, readThumbnail bool) error
}

func (obj *ChildList) GetGDTFs(fileMap map[string]*zip.File, modelLevel int, readThumbnail bool) error {
	for _, element := range obj.SceneObjects {
		err := element.GetGDTF(fileMap, modelLevel, readThumbnail)
		if err != nil {
			return err
		}
		err = element.GetGDTFs(fileMap, modelLevel, readThumbnail)
		if err != nil {
			return err
		}
	}
	for _, element := range obj.GroupObjects {
		err := element.GetGDTFs(fileMap, modelLevel, readThumbnail)
		if err != nil {
			return err
		}
	}
	for _, element := range obj.Fixtures {
		err := element.GetGDTF(fileMap, modelLevel, readThumbnail)
		if err != nil {
			return err
		}
		err = element.GetGDTFs(fileMap, modelLevel, readThumbnail)
		if err != nil {
			return err
		}
	}
	for _, element := range obj.Supports {
		err := element.GetGDTF(fileMap, modelLevel, readThumbnail)
		if err != nil {
			return err
		}
		err = element.GetGDTFs(fileMap, modelLevel, readThumbnail)
		if err != nil {
			return err
		}
	}
	for _, element := range obj.Trusses {
		err := element.GetGDTF(fileMap, modelLevel, readThumbnail)
		if err != nil {
			return err
		}
		err = element.GetGDTFs(fileMap, modelLevel, readThumbnail)
		if err != nil {
			return err
		}
	}
	for _, element := range obj.VideoScreens {
		err := element.GetGDTF(fileMap, modelLevel, readThumbnail)
		if err != nil {
			return err
		}
		err = element.GetGDTFs(fileMap, modelLevel, readThumbnail)
		if err != nil {
			return err
		}
	}
	for _, element := range obj.Projectors {
		err := element.GetGDTF(fileMap, modelLevel, readThumbnail)
		if err != nil {
			return err
		}
		err = element.GetGDTFs(fileMap, modelLevel, readThumbnail)
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *SceneObject) GetGDTF(fileMap map[string]*zip.File, modelLevel int, readThumbnail bool) error {
	return GDTFReader.GetGDTF(fileMap, obj.GDTFSpec, obj.GDTFMode, modelLevel, readThumbnail)
}

func (obj *Fixture) GetGDTF(fileMap map[string]*zip.File, modelLevel int, readThumbnail bool) error {
	return GDTFReader.GetGDTF(fileMap, obj.GDTFSpec, obj.GDTFMode, modelLevel, readThumbnail)
}

func (obj *Support) GetGDTF(fileMap map[string]*zip.File, modelLevel int, readThumbnail bool) error {
	return GDTFReader.GetGDTF(fileMap, obj.GDTFSpec, obj.GDTFMode, modelLevel, readThumbnail)
}

func (obj *Truss) GetGDTF(fileMap map[string]*zip.File, modelLevel int, readThumbnail bool) error {
	return GDTFReader.GetGDTF(fileMap, obj.GDTFSpec, obj.GDTFMode, modelLevel, readThumbnail)
}

func (obj *VideoScreen) GetGDTF(fileMap map[string]*zip.File, modelLevel int, readThumbnail bool) error {
	return GDTFReader.GetGDTF(fileMap, obj.GDTFSpec, obj.GDTFMode, modelLevel, readThumbnail)
}

func (obj *Projector) GetGDTF(fileMap map[string]*zip.File, modelLevel int, readThumbnail bool) error {
	return GDTFReader.GetGDTF(fileMap, obj.GDTFSpec, obj.GDTFMode, modelLevel, readThumbnail)
}
