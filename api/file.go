package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
	"networkDisk/model"
	"networkDisk/service"
	"networkDisk/tool"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func upload(ctx *gin.Context) {
	var UpFile model.File
	var permit string
	permit = ctx.PostForm("permit") //默认不公开
	if permit != "" {
		UpFile.Permit = 1
	}
	//获取文件头
	file, err := ctx.FormFile("upload")
	if err != nil {
		tool.RespBadRequest(ctx, "请求失败")
		return
	}
	//获取文件名
	UpFile.Name = file.Filename
	fmt.Println("文件名:", UpFile.Name)
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	UpFile.ChangeName = GenFileName(UpFile.Name)
	file.Filename = UpFile.ChangeName //修改存在本地的文件名 防止注入（
	UpFile.LocalPath = "../file/" + username
	UpFile.DownloadAddr = "http://127.0.0.1:8080/:username/download?path=" + UpFile.LocalPath + UpFile.Name
	err = ctx.SaveUploadedFile(file, UpFile.LocalPath)
	if err != nil {
		tool.RespBadRequest(ctx, "保存失败")
		return
	}
	err = service.Upload(UpFile)
	if err != nil {
		tool.RespErrorWithDate(ctx, "插入数据失败")
		return
	}
	tool.RespSuccessful(ctx)
}
func GenFileName(file string) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 5)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	//filenameall := path.Base(file)
	//filesuffix := path.Ext(file)
	//prefix := file[0:len(filenameall)-len(filesuffix)]
	filename := []rune(file)
	b = append(b, filename...)
	return string(b)
}
func delFile(ctx *gin.Context) {
	var delFile model.File
	delFile.Name = ctx.PostForm("fileName")
	delFile.LocalPath = ctx.PostForm("filepath")
	iUsername, _ := ctx.Get("username")
	delFile.User = iUsername.(string)
	err := os.Remove(delFile.LocalPath)
	if err != nil {
		fmt.Println(err)
		tool.RespBadRequest(ctx, "请求失败")
		return
	}
	err = service.DelFile(delFile)
	if err != nil {
		fmt.Println(err)
		tool.RespErrorWithDate(ctx, "数据删除失败")
		return
	}
	tool.RespSuccessful(ctx)
	return
}
func changePath(ctx *gin.Context) {
	var file model.File
	file.Name = ctx.PostForm("fileName")
	file.LocalPath = ctx.PostForm("oldPath")
	newPath := ctx.PostForm("newPath")
	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		tool.RespErrorWithDate(ctx, "新路径格式有误")
		return
	}
	err = os.Rename(file.LocalPath+file.Name, newPath+file.Name)
	if err != nil {
		fmt.Println(err)
		tool.RespBadRequest(ctx, "请求失败")
		return
	}
	updateCol := "localPath"
	err = service.ChangePath(file, updateCol, newPath)
	if err != nil {
		fmt.Println(err)
		tool.RespErrorWithDate(ctx, "数据更新失败")
		return
	}
	tool.RespSuccessful(ctx)
	return
}
func changeName(ctx *gin.Context) {
	var file model.File
	file.Name = ctx.PostForm("oldName")
	file.LocalPath = ctx.PostForm("localPath")
	newName := ctx.PostForm("newName")
	err := os.Rename(file.LocalPath+file.Name, file.LocalPath+newName)
	if err != nil {
		fmt.Println(err)
		tool.RespBadRequest(ctx, "请求失败")
		return
	}
	updateCol := "name"
	err = service.ChangePath(file, updateCol, newName)
	if err != nil {
		fmt.Println(err)
		tool.RespErrorWithDate(ctx, "数据更新失败")
		return
	}
	tool.RespSuccessful(ctx)
	return
}
func downloadByLink(ctx *gin.Context) {
	path := ctx.Query("path")
	filename := filepath.Base(path)
	//读取文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		tool.RespBadRequest(ctx, "请求失败")
		return
	}
	defer file.Close()
	//读取文件头部信息
	fileHeader := make([]byte, 512)
	file.Read(fileHeader)

	fileStat, _ := file.Stat()
	//定义响应头
	ctx.Writer.Header().Set("Content-Disposition", "attachment;filename="+filename)   //返回文件名
	ctx.Writer.Header().Set("Content-Type", http.DetectContentType(fileHeader))       //返回检测到的文件类型
	ctx.Writer.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10)) //返回文件大小

	//把文件读取指针归零
	file.Seek(0, 0)
	//io.Copy(ctx.Writer,file)
	for {
		tmp := make([]byte, 10) //通过设置每次读取切片长度控制流速
		n, err := file.Read(tmp)
		if err == io.EOF {
			return
		}
		ctx.Writer.Write(tmp[:n])
		time.Sleep(time.Microsecond) //通过sleep时间控制流速
	}
}
func share(ctx *gin.Context) {
	var shareFile model.File
	var err error
	shareFile.Name = ctx.PostForm("fileName")
	shareFile.LocalPath = ctx.PostForm("path")
	iUsername, _ := ctx.Get("username")
	shareFile.User = iUsername.(string)
	shareFile, err = service.FindFileInfo(shareFile)
	if err != nil {
		fmt.Println(err)
		tool.RespBadRequest(ctx, "查询文件失败")
		return
	}
	if !service.IsPublic(shareFile) {
		tool.RespErrorWithDate(ctx, "文件权限不允许公开下载")
		return
	}
	tool.RespSuccessFulWithDate(ctx, shareFile.DownloadAddr)
}
