package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Author"] = "skyblue"
	this.TplNames = "index.tpl"
	this.Render()
}
