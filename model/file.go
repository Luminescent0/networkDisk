package model
type File struct {
	Id int
	Name string//用户定义的
	ChangeName string//服务端定义的
	User string
	LocalPath string
	Permit int//0为私有 1为公开
	DownloadAddr string
}
