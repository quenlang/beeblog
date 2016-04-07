package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	//beego.Router("/topic/add", &controllers.TopicController{}, "GET:Add")
	beego.AutoRouter(&controllers.TopicController{})
}
