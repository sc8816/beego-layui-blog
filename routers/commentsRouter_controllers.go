package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego-demo/controllers:IndexController"] = append(beego.GlobalControllerRouter["beego-demo/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-demo/controllers:IndexController"] = append(beego.GlobalControllerRouter["beego-demo/controllers:IndexController"],
        beego.ControllerComments{
            Method: "About",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-demo/controllers:IndexController"] = append(beego.GlobalControllerRouter["beego-demo/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Comment",
            Router: `/comment`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-demo/controllers:IndexController"] = append(beego.GlobalControllerRouter["beego-demo/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Message",
            Router: `/message`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-demo/controllers:IndexController"] = append(beego.GlobalControllerRouter["beego-demo/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetRegister",
            Router: `/reg`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-demo/controllers:IndexController"] = append(beego.GlobalControllerRouter["beego-demo/controllers:IndexController"],
        beego.ControllerComments{
            Method: "User",
            Router: `/user`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-demo/controllers:UserController"] = append(beego.GlobalControllerRouter["beego-demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-demo/controllers:UserController"] = append(beego.GlobalControllerRouter["beego-demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-demo/controllers:UserController"] = append(beego.GlobalControllerRouter["beego-demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
