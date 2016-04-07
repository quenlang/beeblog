package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Data["isCategory"] = true
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	op := this.Input().Get("op")
	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return
	case "del":
		id := this.Input().Get("id")
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return
	}

	this.TplName = "category.html"

	var err error
	this.Data["categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
