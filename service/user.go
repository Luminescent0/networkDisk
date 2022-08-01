package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"networkDisk/dao"
	"networkDisk/model"
	"time"
)

func IsRepeatUsername(username string) (bool,error) {
	_,err := dao.SelectUserByUsername(username)
	if err != nil {
		log.Println(err)
		return false,err
	}
	return true,nil
}
func Register(user model.User) error {
	password,err := Cipher(user)
	if err != nil {
		log.Println(err)
		return err
	}
	user.Password = password
	err = dao.InsertUser(user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func Cipher(user model.User) (string,error) {
	password :=[]byte(user.Password)
	nowG := time.Now()
	hashedPassword,err := bcrypt.GenerateFromPassword(password,bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return "",err
	}
	fmt.Println("加密后",string(hashedPassword),"耗时",time.Now().Sub(nowG))
	return string(hashedPassword),nil
}
func UsernameIsExist(username string) error {
	_,err := dao.SelectUserByUsername(username)
	if err != nil {
		fmt.Println(err)
		if err==gorm.ErrRecordNotFound {
			fmt.Println("不存在")
		}
		return err
	}
	return nil
}
func IsPasswordCorrect(username,password string) (bool,error) {
	user,_ := dao.SelectUserByUsername(username)
	flag := ComparePassword(user.Password,[]byte(password))
	if !flag {
		return false,nil
	}
	fmt.Println("验证密码成功")
	return true,nil

}
func ComparePassword(hashedPassword string,plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash,plainPassword)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}