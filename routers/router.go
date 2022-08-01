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
	ns1 := beego.NewNamespace("/ormpractice",
		beego.NSRouter("/add", &controllers.OrmPracticeController{}, "post:Add"),
		beego.NSRouter("/edit", &controllers.OrmPracticeController{}, "post:Update"),
		beego.NSRouter("/del", &controllers.OrmPracticeController{}, "post:Del"),
		beego.NSRouter("/relationread", &controllers.OrmPracticeController{}, "post:RelationRead"),
		beego.NSRouter("/sqlquery", &controllers.OrmPracticeController{}, "post:SqlQuery"),
		beego.NSRouter("/relationadd", &controllers.OrmPracticeController{}, "post:RelationAdd"),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(ns1)
	beego.Include(&controllers.DetailController{})
}
