package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

// ValidateController 邮箱验证
type ValidateController struct {
	beego.Controller
}

func (this *ValidateController) Get() {
	this.TplName = "validate.tpl"
	token := this.Ctx.Input.Param(":token")
	if token == "" {
		return
	}
	isValid, _, err := tokenValidate.Validate(token)
	if err != nil {
		return
	}
	if isValid {
		this.SetSession("Status", 1)
		this.Redirect("/", http.StatusFound)
		return
	}
}
