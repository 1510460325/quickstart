package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"quickstart/models"
)

type UserController struct {
	beego.Controller
}

func (this UserController) Get() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("query is wrong.")
	} else {
		user := models.Read(id)
		data, err := json.Marshal(user)
		if err != nil {
			this.Ctx.WriteString("json is wrong.")
		} else {
			this.Ctx.WriteString(string(data))
		}
	}
}
