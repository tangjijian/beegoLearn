package models

type Profile struct {
	Id     int
	Name   string
	Member *Member `orm:"null;reverse(one)"` // 反向一对一关联
}
