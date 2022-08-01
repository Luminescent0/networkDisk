package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
	"networkDisk/model"
	"networkDisk/service"
	"networkDisk/tool"
	"os"
)

func register(ctx *gin.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	user := model.User{
		Username: username,
		Password: password,
	}
	flag,err := service.IsRepeatUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			fmt.Println("record not found")
		}else {
			log.Println("judge repeat username err:",err)
			tool.RespInternalErr(ctx)
			return
		}
	}
	if flag {
		tool.RespErrorWithDate(ctx,"用户名已存在")
		return
	}
	err = service.Register(user)
	if err != nil {
		fmt.Println("register err: ",err)
		tool.RespInternalErr(ctx)
		return
	}
	err = os.MkdirAll("../file/"+user.Username, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		tool.RespErrorWithDate(ctx,"路径有误")
		return
	}
	tool.RespSuccessFulWithDate(ctx,"注册成功")
}
func verify(ctx *gin.Context) (string, string) { //验证非法输入
	validate := validator.New() //创建验证器
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	u := model.User{Id: 0, Username: username, Password: password}

	err := validate.Struct(u)
	fmt.Println(err)
	if err != nil {
		return "存在非法输入", ""
	}
	return username, password
}
func login(ctx *gin.Context) {
	username,password := verify(ctx)
	err := service.UsernameIsExist(username)
	if err != nil {
		if err==gorm.ErrRecordNotFound{
			tool.RespErrorWithDate(ctx,"用户不存在")
			return
		}
		tool.RespInternalErr(ctx)
		return
	}
	flag,err := service.IsPasswordCorrect(username,password)
	if err != nil {
		fmt.Println("judge password correct err:",err)
		tool.RespInternalErr(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithDate(ctx,"密码错误")
		return
	}
	token,err1 := CreateToken(username)
	if err1 != nil {
		tool.RespInternalErr(ctx)
		return
	}
	err = os.Mkdir("../file/"+username,0777)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalErr(ctx)
	}
	tool.RespSuccessFulWithDate(ctx,token)
	return


}