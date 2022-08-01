package dao

import (
	"fmt"
	"log"
	"networkDisk/model"
)

func SelectUserByUsername(username string) (model.User,error) {
	user := model.User{}
	err := dB.Table("user").Where("username=?",username).Take(&user)
	fmt.Println(err)
	if err != nil {
		fmt.Println("查询失败:",err.Error)
		return user,err.Error
	}
	fmt.Println(user)
	return user,nil
}
func InsertUser(user model.User) error {
	err := dB.Table("user").Select("username","password").Create(&user)
	if err != nil {
		log.Println(err.Error)
		return err.Error
	}
	return nil
}