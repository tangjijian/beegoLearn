package filters

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/casbin/casbin"
	"myproject/models"
	"strings"
)

type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

//type beego struct {
//	FilterFunc func()
//}

func (a *BasicAuthorizer) GetUserRole(input *context.BeegoInput) string {

	user, ok := input.Session("user").(*models.User)
	fmt.Println(user.Role.Name)
	// 判断是否成功通过Session获取用户信息
	if !ok || user.Role.Name == "" {
		// 不成功的话直接返回匿名
		return models.GetRoleString(models.RoleAnonymous)
	}
	return user.Role.Name
}

func NewAuthorizer(e *casbin.Enforcer) beego.FilterFunc {
	return func(ctx *context.Context) {
		// 通过创建结构体，存放Enforcer
		a := &BasicAuthorizer{enforcer: e}
		// 获取用户角色

		userRole := a.GetUserRole(ctx.Input)
		// 获取访问路径
		method := strings.ToLower(ctx.Request.Method)
		// 获取访问方式
		path := strings.ToLower(ctx.Request.URL.Path)
		// 进行验证 - 失败则返回401
		if status := a.enforcer.Enforce(userRole, path, method); !status {
			ctx.Output.Status = 401
			_ = ctx.Output.JSON(map[string]string{"msg": "用户权限不足"}, beego.BConfig.RunMode != "prod", false)
		}
	}
}
