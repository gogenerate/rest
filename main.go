package main

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/shxsun/gails-default/models" // NNNN "{{.PkgPath}}/models"
)

func main() {
	beego.AutoRender = false
	beego.DirectoryIndex = true

	if err := models.InitDB(); err != nil {
		log.Fatal(err)
	}
	beego.Run()
}
