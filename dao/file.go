package dao

import (
	"fmt"
	"log"
	"networkDisk/model"
)

func InsertFile(file model.File) error {
	err := dB.Table("file").Select("name", "changeName", "user", "localPath").Create(&file)
	if err != nil {
		log.Println(err.Error)
		return err.Error
	}
	return nil
}
func DelFile(file model.File) error {
	err := dB.Table("file").Where("name=? and user=? and localPath=?",
		file.Name, file.User, file.LocalPath).Delete(&file)
	if err != nil {
		log.Println(err.Error)
		return err.Error
	}
	return nil
}
func ChangePath(file model.File, updateCol string, newPath string) error {
	err := dB.Table("file").Model(&file).Where("name=?and localPath=?",
		file.Name, file.LocalPath).Update(updateCol, newPath)
	if err != nil {
		log.Println(err.Error)
		return err.Error
	}
	return nil
}
func FindFile(file model.File) (model.File, error) {
	err := dB.Table("file").Where("name=? and user=? and localPath=?", file.Name, file.User, file.LocalPath).Take(&file)
	fmt.Println(err)
	if err != nil {
		fmt.Println("查询失败:", err.Error)
		return file, err.Error
	}
	fmt.Println(file)
	return file, nil
}
