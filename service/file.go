package service

import (
	"networkDisk/dao"
	"networkDisk/model"
)

func Upload(file model.File) error {
	return dao.InsertFile(file)
}

func DelFile(file model.File) error {
	return dao.DelFile(file)
}
func ChangePath(file model.File, updateCol string, newPath string) error {
	return dao.ChangePath(file, updateCol, newPath)
}
func FindFileInfo(file model.File) (model.File, error) {
	return dao.FindFile(file)
}
func IsPublic(file model.File) bool {
	if file.Permit == 0 {
		return false
	}
	return true
}
