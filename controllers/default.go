package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	status := this.GetSession("Status")
	if status == nil {
		this.Redirect("/register", http.StatusFound)
		return
	}
	beego.Debug(status)
	var isValid bool
	if v, ok := status.(int); ok && v == 1 {
		isValid = true
	}
	this.Data["Status"] = isValid
	this.Data["Register"] = this.GetSession("Register")
	this.TplName = "index.tpl"
}
