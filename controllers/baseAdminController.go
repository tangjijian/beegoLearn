package controllers

import "fmt"

type BaseAdminController struct {
	baseController
}

func (r *BaseAdminController) NestPrepare() {

	fmt.Println("admin_NestPrepare")
	//r.Ctx.WriteString("admin_NestPrepare")

}

func (r *BaseAdminController) Get() {
	r.TplName = "Get.tpl"
}

func (r *BaseAdminController) Post() {
	r.TplName = "Post.tpl"
}
func (r *BaseAdminController) Test() {
	r.StopRun()
	r.Ctx.WriteString("test")
}
