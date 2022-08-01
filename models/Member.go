package models

type Member struct {
	Id         int `orm:"pk"`
	Code       int
	Username   string
	Phone      string
	PostName   string
	Profile    *Profile `orm:"rel(one)"` // 一对一关联
	CreateTime string   `orm:"auto_now_add;type(timestamp)"`
	UpdateTime string   `orm:"auto_now;type(timestamp)"`
	Post       []*Post  `orm:"null;reverse(many);on_delete(set_null)"` // 一对多的反向关系
}

// 设置引擎为 INNODB
func (u *Member) TableEngine() string {
	return "INNODB"
}
