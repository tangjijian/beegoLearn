package main

import (
	"github.com/beego/beego/v2/server/web"
	"myproject/controllers"
	_ "myproject/routers"
)

func main() {

	web.BConfig.CopyRequestBody = true
	web.BConfig.Listen.EnableAdmin = true
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
	web.AutoRouter(&controllers.UserController{})
	web.AutoRouter(&controllers.BaseAdminController{})
	//tree := web.PrintTree()
	//methods := tree["Data"].(web.M)
	//for k, v := range methods {
	//	fmt.Printf("%s => %v\n", k, v)
	//}

	//web.Get("api")

	web.Run()
}
func main1() {
	//beego.AutoRouter(&controllers.UserController{})
	//beego.Run()
}
