package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"myproject/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/handleRQ",
		beego.NSRouter("/home", &controllers.HomeController{}, "post:Content"),
		beego.NSRouter("/home/parseform", &controllers.HomeController{}, "post:GetForm"),
	)
	beego.AddNamespace(ns)
	beego.Include(&controllers.DetailController{})
}
