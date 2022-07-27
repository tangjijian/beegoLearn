package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
)

type Role struct {
	Id    int     `orm:"auto;pk" description:"角色序号" json:"role_id"`
	Name  string  `orm:"unique"  description:"角色名"   json:"role_name"`
	Users []*User `orm:"reverse(many)" description:"用户列表" json:"users"`
}

func init() {
	orm.RegisterModel(new(Role))
	// 这里因为在别的地方已经同步过数据库了，就不同步了
	RegisterRoles()
	AddRolesGroupPolicy()
}

var (
	RoleAdmin     = "admin"
	RoleUser      = "user"
	RoleAnonymous = "anonymous"
	RolesId       = map[string]int{
		RoleAdmin:     -1,
		RoleUser:      -1,
		RoleAnonymous: -1,
	}
)

// 注册角色模型 - 初始化
func RegisterRoles() {
	o := orm.NewOrm()
	// 这里我通过遍历上面构建的一个字典来写入数据库
	// 如果不愿意使用骚操作的话，直接写三个ReadOrCreate就好了
	// GetRoleString方法是必须的
	for key, _ := range RolesId {
		_, id, err := o.ReadOrCreate(&Role{Name: GetRoleString(key)}, "Name")
		if err != nil {
			panic(err)
		}
		RolesId[key] = int(id)
	}
}

// 这个方法主要用于在Name字段加个前缀role_
func GetRoleString(s string) string {
	if strings.HasPrefix(s, "role_") {
		return s
	}
	return fmt.Sprintf("role_%s", s)
}

// 向Casbin添加角色继承策略规则
func AddRolesGroupPolicy() {
	// 普通管理员继承用户
	_ = Enforcer.AddGroupingPolicy(GetRoleString(RoleAdmin), GetRoleString(RoleUser))
	// 用户继承匿名者
	_ = Enforcer.AddGroupingPolicy(GetRoleString(RoleUser), GetRoleString(RoleAnonymous))
}
