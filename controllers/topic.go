package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["isTopic"] = true
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	this.TplName = "topic.html"
}

func (this *TopicController) Post() {
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	err := models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}

	this.Data["isTopic"] = true
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	this.TplName = "topic.html"
	this.Data["Topics"], err = models.GetAllTopics()
	if err != nil {
		beego.Error(err)
	}
}

func (this *TopicController) Add() {
	if !CheckAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.TplName = "topic_add.html"
}
