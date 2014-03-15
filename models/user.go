package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	tables = append(tables, new(User))
}

type User struct {
	ID   int64  `json:"id" xorm:"id pk autoincr"`
	Name string `json:"name"` // XXXX
	/*XXXX
	{{range .Cols}}
	{{.Name}} {{.Type}} `json:"{{.ORMName}}"`
	{{end}}
	XXXX*/
}

// create
func (v *User) Create() error {
	_, err := x.Insert(v)
	return err
}

func (v *User) Update(id int64) error {
	affec, err := x.Id(id).Update(v)
	if err == nil && affec == 0 {
		err = fmt.Errorf("update {{.Table}}(id:%d) failed", id)
	}
	return err
}

func GetUser(id int64) (v *User, err error) {
	v = new(User)
	ok, err := x.Id(id).Get(v)
	if err == nil && !ok {
		err = fmt.Errorf("get {{.Table}}(id:%d) failed", id)
	}
	return
}

func (v *User) Save() error {
	if v.ID == 0 {
		return v.Create()
	}
	return v.Update(v.ID)
}

func AllUser() (vs []User, err error) {
	err = x.Find(&vs)
	return
}

func DelUser(id int64) error {
	affec, err := x.Id(id).Delete(new(User))
	if err == nil && affec == 0 {
		err = fmt.Errorf("{{.Table}}(id:%d) already deleted", id)
	}
	return err
}
