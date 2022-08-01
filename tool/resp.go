package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespErrorWithDate(ctx *gin.Context,data interface{})  {
	ctx.JSON(http.StatusOK,gin.H{"info":data})
}

func RespInternalErr(ctx *gin.Context)  {
	ctx.JSON(http.StatusInternalServerError,gin.H{"info":"服务器错误"})
}
func RespSuccessful(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK,gin.H{"info":"成功"})
}
func RespSuccessFulWithDate(ctx *gin.Context,data interface{})  {
	ctx.JSON(http.StatusOK,gin.H{
		"info":"成功",
		"data":data,
	})
}
func RespBadRequest(ctx *gin.Context,data interface{}) {
	ctx.JSON(http.StatusBadRequest,gin.H{
		"info":"失败",
		"data":data,
	})
}