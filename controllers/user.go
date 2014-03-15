package controllers

import "github.com/shxsun/gails-default/models" // NNNN import "{{.PkgPath}}/models"

type UserController struct {
	APIController
}

// POST /api/user/save
func (this *UserController) Save() {
	v := new(models.User)
	v.ID, _ = this.GetInt("id")
	v.Name = this.MustString("name") // XXXX
	/* XXXX
		{{range .Cols}} v.{{.Name}} = this.Get{{.Type|title}}("{{.ORMName}}")
		{{end}}
	XXXX */
	this.err = v.Save()
}

// POST /api/user/:id
func (this *UserController) Post() {
	v := new(models.User)
	v.ID, _ = this.GetInt(":id")
	v.Name = this.MustString("name") // XXXX
	//NNNN {{range .Cols}} v.{{.Name}} = this.Must{{.Type|title}}("{{.ORMName}}")
	//NNNN {{end}}
	this.err = v.Save()
}

// GET /api/user/:id
func (this *UserController) Get() {
	id := this.MustInt64(":id")
	v, err := models.GetUser(id)
	if err != nil {
		this.err = err
		return
	}
	this.data = v
}

// GET /api/user/all
func (this *UserController) All() {
	this.data, this.err = models.AllUser()
}

// DELELE /api/user/delete
func (this *UserController) Delete() {
	id := this.MustInt64(":id")
	this.err = models.DelUser(id)
}
