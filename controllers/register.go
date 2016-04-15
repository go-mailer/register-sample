package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/mail"
	"text/template"

	"github.com/go-mailer/send"

	"github.com/go-mailer/register-sample/models"

	"github.com/astaxie/beego"
)

// RegisterController 注册管理
type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	status := this.GetSession("Status")
	if status != nil {
		this.Redirect("/", http.StatusFound)
		return
	}
	this.TplName = "register.tpl"
}

func (this *RegisterController) Post() {
	var reg models.Register
	if err := this.ParseForm(&reg); err != nil {
		panic(err)
	}
	beego.Debug(reg)
	err := this.sendEmail(reg)
	if err != nil {
		panic(err)
	}
	this.SetSession("Status", 0)
	this.SetSession("Register", reg)
	this.Redirect("/", http.StatusFound)
}

func (this *RegisterController) sendEmail(reg models.Register) error {
	config, err := beego.AppConfig.GetSection("email")
	if err != nil {
		return err
	}
	from := mail.Address{Name: config["name"], Address: config["address"]}
	sender, err := send.NewSmtpSender(config["server"], from, config["authpwd"])
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	token, err := tokenValidate.Generate(reg.Email)
	if err != nil {
		return err
	}
	link := fmt.Sprintf("%s:%d/validate/%s", config["link"], beego.AppConfig.DefaultInt("httpport", 8080), token)
	tmpl := template.Must(template.New("").Parse(this.getEmailTpl()))
	err = tmpl.Execute(buf, map[string]interface{}{"Link": link})
	if err != nil {
		return err
	}
	msg := &send.Message{
		Subject: "Beego邮箱验证范例",
		Content: buf,
		To:      []string{reg.Email},
	}
	err = sender.AsyncSend(msg, false, func(err error) {
		if err != nil {
			beego.Error(err)
		}
	})
	return err
}

func (this *RegisterController) getEmailTpl() string {
	return `
<!DOCTYPE html>
<html lang="zh-CN">
	<head>
		<meta charset="UTF-8">
		<title>邮件验证</title>
		<style type="text/css">
		.alert {
		padding: 15px;
		margin-bottom: 20px;
		border: 1px solid transparent;
		border-radius: 4px;
		}
		.alert-info {
		color: #31708f;
		background-color: #d9edf7;
		border-color: #bce8f1;
		}
		</style>
	</head>
	<body>
		<h1>欢迎您注册Beego的邮件测试范例</h1>
		<div class="alert alert-info">
			请点击链接激活邮箱：
			<a href="{{.Link}}" class="alert-link">{{.Link}}</a>
		</div>
	</body>
</html>
	`
}
