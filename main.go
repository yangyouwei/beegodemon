package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/yangyouwei/beegodemon/controllers"
	"github.com/yangyouwei/beegodemon/models"
	_ "github.com/yangyouwei/beegodemon/routers"
)

func init()  {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default",false,true)  //默认是不会自动建表的。只建一次表
	beego.Router("/",&controllers.HomeController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Run()
}