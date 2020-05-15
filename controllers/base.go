package controllers

import (
	"beego-demo/models"
	"beego-demo/syserror"
	"errors"
	"github.com/astaxie/beego"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
	User    models.User
	IsLogin bool
}

func (this *BaseController) Prepare() {
	this.Data["Path"] = this.Ctx.Request.RequestURI
	user, ok := this.GetSession(SESSION_USER_KEY).(models.User)
	this.IsLogin = false
	if ok {
		this.User = user
		this.IsLogin = true
		this.Data["User"] = this.User
	}
	this.Data["isLogin"] = this.IsLogin
}
func (this *BaseController) Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500")
}
func (this *BaseController) GetMustString(key, msg string) string {
	val := this.GetString(key)
	if len(val) == 0 {
		this.Abort500(errors.New(msg))
	}
	return val
}
func (this *BaseController) MustLogin() {
	if !this.IsLogin {
		this.Abort500(syserror.NoUserError{})
	}
}

type M map[string]interface {
}

func (this *BaseController) JsonOK(msg, action string) {
	this.Data["json"] = M{
		"code":   0,
		"msg":    msg,
		"action": action,
	}
	this.ServeJSON()
}
func (this *BaseController) JsonOKM(msg string, data M) {
	data["code"] = 0
	data["msg"] = msg
	this.Data["json"] = data
	this.ServeJSON()
}
