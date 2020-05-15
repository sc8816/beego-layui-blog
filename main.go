package main

import (
	"beego-demo/models"
	_ "beego-demo/models"
	_ "beego-demo/routers"
	"encoding/gob"
	"github.com/astaxie/beego"
	"strings"
)

func main() {
	initSession()
	initTemplate()
	beego.Run()
}
func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		x1 := strings.Trim(x, "/")
		y1 := strings.Trim(y, "/")
		return strings.Compare(x1, y1) == 0
	})
}
func initSession() {
	gob.Register(models.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "lite_blog"
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}
