package main

import (
	"encoding/gob"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/session"
	"github.com/beego/beego/v2/server/web/session/redis"
	_ "github.com/beego/beego/v2/server/web/session/redis"
	"myproject/controllers"
	"myproject/filters"
	"myproject/models"
	_ "myproject/routers"
)

func init() {
	gob.Register(&models.User{})
}
func main() {
	pro, _ := session.GetProvider("redis")
	provider := pro.(*redis.Provider)
	provider.Password = "12345"
	//fmt.Printf("%T", pro)
	//provider := pro.(redis.Provider)
	//provider.Password = "123456"
	beego.InsertFilter("/v1/user/*", beego.BeforeRouter, filters.NewAuthorizer(models.Enforcer)) // 过滤器
	beego.BConfig.CopyRequestBody = true
	//beego.BConfig.Listen.EnableAdmin = true
	orm.RunSyncdb("casbin", false, true) // 同步
	//web.Router("api", &controllers.UserController{}, "get:GetUserByID")
	//创建namespace
	//ns := web.NewNamespace("/android",
	//	web.NSCtrlGet("/api/getuser", controllers.UserController.GetUserByID),
	//	web.NSCtrlPost("/api/update/*", controllers.UserController.UpdateUser),
	//	web.NSCtrlGet("/api/update", controllers.UserController.UpdateUser),
	//	web.NSCtrlPost("/api/delete/?:username", controllers.UserController.Del),
	//	//web.NSBefore(func(ctx *context.Context) {
	//	//	ctx.WriteString("nsBefore")
	//	//}),
	//	//web.NSAfter(func(ctx *context.Context) {
	//	//	fmt.Println("after filter")
	//	//}),
	//	web.NSNamespace("/admin",
	//		web.NSRouter("/user", &controllers.UserController{}, "get:HelloWorld"),
	//	),
	//)
	//ns.Filter("before", func(ctx *context.Context) {
	//	conf, _ := web.AppConfig.String("abc")
	//	ctx.WriteString(conf)
	//})
	//ns1 := web.NewNamespace("/ios",
	//	web.NSRouter("/list/*.*", &controllers.UserController{}),
	//)
	//web.AddNamespace(ns)
	//web.AddNamespace(ns1)
	ns := beego.NewNamespace("/v1",
		beego.NSCtrlGet("/user/:id", controllers.UserController.Info),
	)
	beego.AddNamespace(ns)
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.BaseAdminController{})
	//tree := web.PrintTree()
	//methods := tree["Data"].(web.M)
	//for k, v := range methods {
	//	fmt.Printf("%s => %v\n", k, v)
	//}

	//web.Get("api")

	beego.Run()
}
func main1() {
	//beego.AutoRouter(&controllers.UserController{})
	//beego.Run()
}
