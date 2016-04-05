package main

import (
	"beeblog/models"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// register database
	models.RegisterDB()
}

func main() {
	// open ORM debug model
	orm.Debug = true
	// auto create object
	orm.RunSyncdb("default", false, true)

	// startup
	beego.Run()

}
