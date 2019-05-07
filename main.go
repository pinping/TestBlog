package main

import (
	_ "TestBlog/routers"
	"github.com/astaxie/beego"
	"TestBlog/models"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default",false,false)
	beego.Run()
}

