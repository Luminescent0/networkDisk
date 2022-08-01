package main

import (
	"networkDisk/api"
	"networkDisk/dao"
)

func main()  {
	dao.InitDB()
	api.InitEngine()
}