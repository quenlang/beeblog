package controllers

import (
	"beeblog/models"
	//"fmt"
	"github.com/astaxie/beego"
	"path"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["isTopic"] = true
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	var err error
	this.Data["Topics"], err = models.GetAllTopics()
	if err != nil {
		beego.Error(err)
	}
	this.TplName = "topic.html"
}

func (this *TopicController) Post() {
	if !CheckAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")

	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	if fh != nil {
		attachment = fh.Filename
		beego.Info(attachment)
		err := this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}

	if len(tid) == 0 {
		err := models.AddTopic(title, content, attachment)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.UpdateTopic(tid, title, content, attachment)
		if err != nil {
			beego.Error(err)
		}
	}

	this.Data["isTopic"] = true
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	if !CheckAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	this.TplName = "topic_add.html"
}

func (this *TopicController) Modify() {
	if !CheckAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	tid := this.Input().Get("tid")
	var err error
	this.Data["Topic"], err = models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
	}
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	this.Data["Tid"] = tid
	//fmt.Println("controller:" + tid)
	this.TplName = "topic_modify.html"
}

func (this *TopicController) Delete() {
	if !CheckAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	tid := this.Input().Get("tid")
	err := models.DelTopic(tid)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}

func (this *TopicController) View() {
	this.Data["isTopic"] = true
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	this.TplName = "topic_view.html"

	id := this.Ctx.Input.Params()["0"]
	var err error
	this.Data["Topic"], err = models.GetTopic(id)
	if err != nil {
		beego.Error(err)
	}
}
