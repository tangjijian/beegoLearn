package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"myproject/models"
)

type User struct {
	beego.Controller
}

func init() {
	registerUserPolicy()
}

// 只有管理员才能注册
// @router  /register [post]
func (c *User) Register() {}

// 只有用户、管理员才能看到别人或者自己的个人资料
// 因为管理员继承用户，所以用户能做到的，管理员也可
// @router  /profile [get]
func (c *User) Profile() {}

// 匿名也能登陆
// @router /login [post]
func (c *User) Login() {}

func registerUserPolicy() {
	// Path前缀,这个根据具体项目自行调整
	api := "/v1/user"
	// 路由的Policy
	adminPolicy := map[string][]string{
		"/register": {"post"},
	}
	userPolicy := map[string][]string{
		// 注意 - casbin.conf中使用 keyMatch2 对 obj 进行
		// 验证,这里要使用 :id 来对参数进行标识
		"/:id": {"get", "put", "delete"},
	}
	anonymousPolicy := map[string][]string{
		"/login": {"post"},
	}
	// models.RoleAdmin      = "admin"
	// models.RoleUser       = "user"
	// models.RoleAnonymous  = "anonymous"
	AddPolicyFromController(models.RoleAdmin, adminPolicy, api)
	AddPolicyFromController(models.RoleUser, userPolicy, api)
	AddPolicyFromController(models.RoleAnonymous, anonymousPolicy, api)
}
func AddPolicyFromController(role string, policy map[string][]string, api string) {
	for path := range policy {
		for _, method := range policy[path] {
			// models.Enforcer在models/Casbin.go中定义并初始化
			_ = models.Enforcer.AddPolicy(models.GetRoleString(role), fmt.Sprintf("%s%s", api, path), method)
		}
	}
}
