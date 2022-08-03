package models

type Tag struct {
	Id    int `orm:"fk"`
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}
