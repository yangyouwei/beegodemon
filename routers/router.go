package routers

import (
	"github.com/yangyouwei/beegodemon/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.HomeController{})
	beego.Router("/login",&controllers.LoginController{})
}
