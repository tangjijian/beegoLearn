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

	// 初始化 Casbin
	//RegisterCasbin()
}

// 注意，这个Enforcer很重要，Casbin使用都是调用这个变量

//
//type Adapter struct {
//	o orm.Ormer
//}

//func RegisterCasbin() {
//	a := &Adapter{}
//	a.o = orm.NewOrm()
//	// 这个我不知道干嘛的
//	runtime.SetFinalizer(a, finalizer)
//	// Enforcer初始化 - 即传入Adapter对象
//	Enforcer = casbin.NewEnforcer("conf/casbin.conf", a) // 初始化为orm-adapter
//	// Enforcer读取Policy
//	err := Enforcer.LoadPolicy()
//	if err != nil {
//		panic(err)
//	}
//}

// finalizer is the destructor for Adapter.
// 这个函数里面啥都没有，就是这样
//func finalizer(_ *Adapter) {}

// 注意，方法对应的具体代码要从beego-ORM-Adapter/adapter.go中复制过来
// 这里方法里面使用的orm操作还是要根据自己的
// 实际情况作出调整，不要盲目复制
//func loadPolicyLine(line CasbinRule, model model.Model)     {}
//func savePolicyLine(ptype string, rule []string) CasbinRule {}

//func (a *Adapter) LoadPolicy(model model.Model) error                         {}
//func (a *Adapter) SavePolicy(model model.Model) error                         {}
//func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error    {}
//func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {}
//func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
//}
//func (a *Adapter) SavePolicy(model model.Model) error {
//	var lines []CasbinRule
//	for ptype, ast := range model["p"] {
//		for _, rule := range ast.Policy {
//			line := savePolicyLine(ptype, rule)
//			lines = append(lines, line)
//		}
//	}
//	for ptype, ast := range model["g"] {
//		for _, rule := range ast.Policy {
//			line := savePolicyLine(ptype, rule)
//			lines = append(lines, line)
//		}
//	}
//	_, err := a.o.InsertMulti(len(lines), lines)
//	return err
//}
