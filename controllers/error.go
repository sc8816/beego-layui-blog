package controllers

import (
	"beego-demo/syserror"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ErrorController struct {
	beego.Controller
}

//ajax return {code:"",message:""}
func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
	c.Data["content"] = "page not found"
	if c.IsAjax() {
		c.jsonerr(syserror.Error404{})
	}
}
func (c *ErrorController) Error500() {
	c.TplName = "error/500.html"
	err, ok := c.Data["error"].(error)
	if !ok {
		err = syserror.New("未知错误", nil)
	}
	serr, ok := err.(syserror.Error)
	if !ok {
		err = syserror.New(err.Error(), nil)
	}
	if serr.ReasonErr() != nil {
		logs.Info(serr.Error(), serr.ReasonErr())
	}
	if c.IsAjax() {
		c.jsonerr(serr)
	} else {
		c.Data["content"] = serr.Error()
	}
}
func (c *ErrorController) jsonerr(serr syserror.Error) {
	c.Ctx.Output.Status = 200
	c.Data["json"] = map[string]interface{}{
		"code":    serr.Code(),
		"message": serr.Error(),
	}
	c.ServeJSON()
}
