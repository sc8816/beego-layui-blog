package controllers

import (
	"beego-demo/models"
	"beego-demo/syserror"
	"errors"
	"github.com/astaxie/beego/logs"
	"strings"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (this *UserController) Login() {
	email := this.GetMustString("email", "邮箱不能为空")
	password := this.GetMustString("password", "密码不能为空")
	user, err := models.QueryUserByEmailAndPwd(email, password)
	if err != nil {
		this.Abort500(syserror.New("登陆失败", err))
	}
	logs.Info(user)
	this.SetSession(SESSION_USER_KEY, user)
	this.JsonOK("登陆成功", "/")
}

// @router /register [post]
func (this *UserController) Register() {
	//邮箱 密码 确认密码 用户名都不为空
	name := this.GetMustString("name", "用户名不为空")
	email := this.GetMustString("email", "邮箱不能为空")
	password := this.GetMustString("password", "密码不能为空")
	again := this.GetMustString("again", "确认密码不能为空")
	if strings.Compare(password, again) != 0 {
		this.Abort500(syserror.New("用户名重复",errors.New("")))
	}
	if u, err := models.QueryUserByName(name); err == nil && u.ID > 0 {
		this.Abort500(syserror.New("用户名重复", err))
	}
	if u, err := models.QueryUserByEmail(email); err == nil && u.ID > 0 {
		this.Abort500(syserror.New("用户名重复",errors.New("")))
	}
	if err := models.SaveUser(&models.User{
		Name:   name,
		Email:  email,
		Pwd:    password,
		Avatar: "/static/images/item.png",
		Role:   1,
	}); err != nil {
		this.Abort500(syserror.New("用户保存失败", err))
	}
	this.JsonOK("注册成功", "/user")
}

// @router /logout [get]
func (this *UserController) Logout() {
	this.MustLogin()
	this.DelSession(SESSION_USER_KEY)
	this.Redirect("/", 302)

}
