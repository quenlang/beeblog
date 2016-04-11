package controllers

import (
	//"fmt"
	"beeblog/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["isHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	var err error
	this.Data["Topics"], err = models.GetAllTopics()
	if err != nil {
		beego.Error(err)
	}
	//this.Ctx.WriteString(fmt.Sprintln(CheckAccount(this.Ctx)))
}
