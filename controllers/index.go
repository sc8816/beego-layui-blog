package controllers

type IndexController struct {
	BaseController
}
// @router / [get]
func (that *IndexController) Index()  {
	that.TplName = "index.html"
}

// @router /about [get]
func (that *IndexController) About()  {
	that.TplName = "about.html"
}

// @router /message [get]
func (that *IndexController) Message()  {
	that.TplName = "message.html"
}

// @router /comment [get]
func (that *IndexController) Comment()  {
	that.TplName = "comment.html"
}
// @router /user [get]
func (that *IndexController) User()  {
	that.TplName = "user.html"
}
// @router /reg [get]
func (that *IndexController) GetRegister()  {
	that.TplName = "register.html"
}