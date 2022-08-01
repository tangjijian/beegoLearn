package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	// 用户模型
	Id       int    `orm:"auto;pk" description:"用户序号" json:"uid"`
	Username string `orm:"unique" description:"用户名" json:"username"`
	Password string `description:"用户密码" json:"password"`
	Role     *Role  `orm:"rel(fk);null" description:"角色" json:"Role"`
}

// 各种ORM查询方法请自行实现，这里不强调
//func init() {
//	orm.RegisterModel(new(User))
//}

func AddOne(u User) string {
	return "1"
}
func CreateTable() {
	o := orm.NewOrm()
	// default 只能注册一个
	err := orm.RegisterDataBase("go_casbin", "mysql", "root:root@tcp(127.0.0.1:3306)/go_casbin")
	fmt.Println(err)
	o.Using("go_casbin")
	_ = orm.RunSyncdb("go_casbin", false, true)

}
