package routers

import (
	"github.com/astaxie/beego"
	"github.com/go-mailer/register-sample/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/validate/:token", &controllers.ValidateController{})
}
