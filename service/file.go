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
func ChangePath(file model.File,updateCol string,newPath string) error {
	return dao.ChangePath(file,updateCol,newPath)
}