package controllers

import (
	"github.com/astaxie/beego"
	context2 "github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("usernam","",-1,"/")
		c.Ctx.SetCookie("passworld","",-1,"/")
		c.Redirect("/",301)
		return
	}
	c.TplName = "login.html"
}

func (c *LoginController)Post()  {
	//c.Ctx.WriteString(fmt.Sprint(c.Input()))
	//return
	username := c.Input().Get("uname")
	passworld := c.Input().Get("pwd")
	autologin := c.Input().Get("autologin") == "on" //判断并赋值 bool值

	if beego.AppConfig.String("username") == username && beego.AppConfig.String("passworld") == passworld {
		maxAge := 0
		if autologin {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("username",username,maxAge,"/")
		c.Ctx.SetCookie("passworld",passworld,"/")
	}
	c.Redirect("/",301)
	return
}

//检查cookie
func checkAccount(c *context2.Context) bool {
	ck, err := c.Request.Cookie("username")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = c.Request.Cookie("passworld")

	if err != nil {
		return false
	}
	pwd := ck.Value

	return beego.AppConfig.String("username") == uname && beego.AppConfig.String("passworld") == pwd
}