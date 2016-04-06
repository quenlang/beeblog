package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	//this.Ctx.WriteString(fmt.Sprint(autoLogin))

	if uname == beego.AppConfig.String("uname") && pwd == beego.AppConfig.String("pwd") {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
			this.Ctx.SetCookie("uname", uname, maxAge, "/")
			this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		}
	}
	this.Redirect("/", 302)
	return
}