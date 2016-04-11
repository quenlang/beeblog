package main

import (
	"beeblog/models"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
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

	os.Mkdir("attachment", os.ModePerm)
	// startup
	beego.Run()

}
