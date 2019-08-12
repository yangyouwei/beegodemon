package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "home.html"

	//检查cookie
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}
