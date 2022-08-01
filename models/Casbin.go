package models

import (
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	beegoormadapter "github.com/casbin/beego-orm-adapter"
	"github.com/casbin/casbin"
	_ "github.com/go-sql-driver/mysql"
)

var Enforcer *casbin.Enforcer

type CasbinRule struct {
	Id    int    // 自增主键
	PType string // Policy Type - 用于区分 policy和 group(role)
	V0    string // subject
	V1    string // object
	V2    string // action
	V3    string // 这个和下面的字段无用，仅预留位置，如果你的不是
	V4    string // sub, obj, act的话才会用到
	V5    string // 如 sub, obj, act, suf就会用到 V3
}

func init() {

	orm.RegisterModel(new(Role))
	orm.RegisterModel(new(User))
	// 实际上同步数据库在整个Beego项目中只需要执行一次，如果
	// 您在别的地方已经同步数据库，这里就不用在执行一次 RunSyncdb
	a := beegoormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/")

	Enforcer = casbin.NewEnforcer(beego.AppPath+"/conf/casbin.conf", a)

	err := Enforcer.LoadPolicy()
	if err != nil {
		panic(err)
	}

}
