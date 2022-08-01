package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"myproject/models"
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
	str := u.Ctx.Input.Param(":splat") // 获取正则路由 * 的匹配值
	u.Ctx.WriteString(str)
}
func (u UserController) Del() {
	username := u.Ctx.Input.Param(":username")
	u.Ctx.WriteString(username)
}
func (u *UserController) Get() {
	u.Ctx.WriteString(u.Ctx.Input.Param(":path") + "." + u.Ctx.Input.Param(":ext")) // 回去URL和后缀
}
func (u UserController) Params() {
	str := fmt.Sprint(u.Ctx.Input.Params())
	u.Ctx.WriteString(str)
}

func (u *UserController) Login() {
	//username := u.GetString("username")
	//password := u.GetString("password")
	user := &models.User{Id: 1, Username: "张三", Password: "123456", Role: &models.Role{Id: 2, Name: "role_user"}}
	err := u.SetSession("user", user)
	if err != nil {
		u.Ctx.WriteString(err.Error())
	}
	u.Ctx.WriteString("成功")
}
func (u UserController) Info() {
	u.Ctx.WriteString("用户详情")
}
