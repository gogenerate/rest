package main

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/shxsun/trs/controllers"
	// NNNN "{{.PkgPath}}/controllers"
	"github.com/gogenerate/rest/models" // NNNN "{{.PkgPath}}/models"
)

func main() {
	beego.AutoRender = false
	//beego.DirectoryIndex = true

	beego.Router("/", &controllers.MainController{})
	if err := models.InitDB(); err != nil {
		log.Fatal(err)
	}
	beego.Run()
}
