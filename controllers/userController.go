package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (u *UserController) HelloWorld() {
	u.Ctx.WriteString("hello world")
}
func (u UserController) GetUserByID() {
	// 拦截慌乱
	defer func() {
		if p := recover(); p != nil {
			//u.Data["abc"] = "ABC"
			m := make(map[string]string, 10)
			m["abc"] = "123"

			_ = u.Ctx.JSONResp(m)
		}
	}()
	panic("异常了")
	//var m map[string]string

	//u.Ctx.WriteString("GetUserByID")
}
func (u UserController) UpdateUser() {
	str := u.Ctx.Input.Param(":splat")
	u.Ctx.WriteString(str)
}
func (u UserController) Del() {
	username := u.Ctx.Input.Param(":username")
	u.Ctx.WriteString(username)
}
func (u *UserController) Get() {
	u.Ctx.WriteString(u.Ctx.Input.Param(":path") + "." + u.Ctx.Input.Param(":ext"))
}
func (u UserController) Params() {
	str := fmt.Sprint(u.Ctx.Input.Params())
	u.Ctx.WriteString(str)
}
