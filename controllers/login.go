package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController)Post()  {
	c.Ctx.WriteString(fmt.Sprint(c.Input()))
	return
}