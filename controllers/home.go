package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["isHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	//this.Ctx.WriteString(fmt.Sprintln(CheckAccount(this.Ctx)))
}
