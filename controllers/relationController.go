package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/server/web"
	"myproject/models"
)

type RelationController struct {
	web.Controller
	o orm.Ormer
}

// 多对一
/*

postName:尚能实业质检中心 userName:任松
postName:尚能实业销售部 userName:任松
postName:尚能实业人力资源部 userName:任松
postName:尚能实业企管中心 userName:任松

*/
func (c *RelationController) ManyToOne() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	p := new(models.Post)
	qs := c.o.QueryTable(p)
	var posts []models.Post
	_, err := qs.Filter("member__id", 144).RelatedSel().All(&posts)
	if err != nil {
		return
	}
	for _, post := range posts {
		fmt.Println("postName:"+post.PostName, "userName:"+post.Member.Username)
	}
	c.Ctx.WriteString(fmt.Sprint(posts))
}
