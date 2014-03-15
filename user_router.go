package main

import (
	"github.com/astaxie/beego"
	"github.com/shxsun/gails-default/controllers" // NNNN "{{.PkgPath}}/controllers"
)

func init() {
	beego.Router("/api/user/new", &controllers.UserController{}, "post:Save")
	beego.Router("/api/user/all", &controllers.UserController{}, "get:All")
	beego.Router("/api/user/:id(\\d+)", &controllers.UserController{}) // GET + PUT  + DELETE
}
