package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

func InitEngine()  {
	engine := gin.Default()
	engine.POST("/register",register) //注册
	engine.POST("/login",login) //登录

	fileGroup := engine.Group("/:username")
	{
		fileGroup.Use(JwtAuthMiddleware)
		fileGroup.POST("/upload",upload) //上传文件
		fileGroup.DELETE("/delete",delFile) //删除文件
		fileGroup.PUT("/changePath",changePath) //修改文件路径
		fileGroup.PUT("/changeName",changeName) //重命名
		fileGroup.GET("/download",downloadByLink)//下载
	}

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}