package models

import (
	"time"
)

type Post struct {
	Id             int `orm:"pk"`
	PostName       string
	Code           int
	DepartmentCode int
	DepartmentName string
	CreateTime     time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdateTime     time.Time `orm:"auto_now;type(timestamp)"`
	Member         *Member   `orm:"null;rel(fk)"` // 会给数据表加字符member_id  ； 设置一对多的关系,删除member时会反向删除post
}

// 设置引擎为 INNODB
func (u *Post) TableEngine() string {
	return "INNODB"
}
