package controllers

import (
	"time"

	"github.com/go-mailer/validate"
)

var (
	tokenValidate *validate.TokenValidate
)

func init() {
	// 创建验证信息存储，每10分钟执行一次GC
	store := validate.NewMemoryStore(time.Second * 60 * 10)
	// 创建验证信息管理器，验证信息的过期时间为1小时
	tokenValidate = validate.NewTokenValidate(store, validate.Config{Expire: time.Second * 60 * 60})
}
