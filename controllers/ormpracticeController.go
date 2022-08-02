package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"

	"github.com/beego/beego/v2/server/web"
	"myproject/models"
)

type OrmPracticeController struct {
	web.Controller
	o orm.Ormer
}

func (c *OrmPracticeController) Add() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	hash := md5.New()
	b := hash.Sum([]byte{1, 2, 3, 4, 5})
	data := models.User{Username: "张三", Password: string(b), Role: &models.Role{Name: "role_user"}}
	id, err := c.o.Insert(&data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id)
	}
	c.Ctx.WriteString("成功")
}
func (c *OrmPracticeController) Update() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	user := models.User{Id: 1, Username: "张三", Role: &models.Role{Name: "role_user"}}
	str := "123456"
	b := md5.Sum([]byte(str))
	user.Password = fmt.Sprintf("%x", b)

	update, err := c.o.Update(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(update)
	c.Ctx.WriteString("更新成功")
}
func (c *OrmPracticeController) Del() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	member := &models.Member{Id: 111}
	i, err := c.o.Delete(member)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}
	c.Ctx.WriteString("成功" + strconv.FormatInt(i, 10))

}

// 关联查询
func (c *OrmPracticeController) RelationRead() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	var post []models.Post
	qs := c.o.QueryTable("post")
	_, err := qs.Filter("Member__Username", "陈爱华").All(&post)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}
	c.Ctx.WriteString(fmt.Sprint(post))
}

// sql 查询
func (c *OrmPracticeController) SqlQuery() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	var maps []orm.Params
	_, _ = c.o.Raw("select * from member").Values(&maps)

	for _, item := range maps {
		fmt.Println(item["id"], ":", item["username"])
	}
	c.Ctx.WriteString(fmt.Sprint(maps))
}

// 关联新增
func (c *OrmPracticeController) RelationAdd() {
	profileName := c.GetString("profile_name")
	memberName := c.GetString("member_name")
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	profile := new(models.Profile)
	profile.Name = profileName
	member := new(models.Member)
	member.PostName = "测试关联"
	member.Username = memberName
	member.Profile = profile
	// 设置了关联和反向关联，在新增数据时能自动写入关联的id
	c.o.Insert(profile)
	c.o.Insert(member)
	c.Ctx.WriteString("新增成功")
}

/**
CURD
*/

// read 通过主键查询
func (c *OrmPracticeController) Read() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	//m := models.Member{Id: 1}
	m := models.Member{Username: "管理员"}
	err := c.o.Read(&m, "Username") // 默认使用主键查询，可以指定第二个参数进行条件查询
	if err == orm.ErrNoRows {
		c.Ctx.WriteString("没有数据")
	} else if err == orm.ErrMissPK {
		c.Ctx.WriteString("没有默认主键")
	} else {
		c.Ctx.WriteString(strconv.Itoa(m.Id) + m.Username + m.CreateTime)
	}
}

// readOrCreate
func (c *OrmPracticeController) ReadOrCreate() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	pass := "123456"
	h := md5.Sum([]byte(pass))
	u := models.User{Username: "李四", Password: fmt.Sprintf("%x", h)}
	created, id, err := c.o.ReadOrCreate(&u, "username")
	if created {
		c.Ctx.WriteString("Insert an object Id: " + strconv.Itoa(int(id)))
	} else if err != nil {
		c.Ctx.WriteString(err.Error())
	} else {
		c.Ctx.WriteString("Get data Id :" + strconv.FormatInt(id, 10))
	}
}

// 多行插入
func (c *OrmPracticeController) InsertMulti() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	pass := "123456"
	h := md5.Sum([]byte(pass))
	relP := fmt.Sprintf("%x", h)
	u := []models.User{
		{Username: "金财1", Password: relP},
		{Username: "金财2", Password: relP},
		{Username: "金财3", Password: relP},
		{Username: "金财4", Password: relP},
	}
	multi, err := c.o.InsertMulti(100, &u)
	if err != nil {
		c.Ctx.WriteString("Inserted lines is " + strconv.FormatInt(multi, 10))
	}
}

// Update 默认更新所有字段，可以指定更新的字段

func (c *OrmPracticeController) UpdateTwo() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	u := models.User{Id: 3}
	if c.o.Read(&u) == nil {
		u.Username = "金财发财了"
		update, err := c.o.Update(&u)
		if err != nil {
			c.Ctx.WriteString(err.Error())
		} else {
			c.Ctx.WriteString("Effect lines is " + strconv.FormatInt(update, 10))
		}
	}
}

// Delete 的反向操作
func (c *OrmPracticeController) DeleteTwo() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	m := models.Member{Id: 114}
	i, err := c.o.Delete(&m)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	} else {
		c.Ctx.WriteString("Effect lines is " + strconv.FormatInt(i, 10))
	}
}

// Expr 概览
func (c *OrmPracticeController) Expr() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	m := new(models.Member)
	var member []models.Member
	qs := c.o.QueryTable(m)
	qs = qs.Filter("profile__id__gt", 3) // where profile.id > 3
	_, err := qs.All(&member)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}
	c.Ctx.WriteString(fmt.Sprint(member))
}

// exact/ iexact
func (c *OrmPracticeController) ExactAndI() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	p := new(models.Profile)
	qs := c.o.QueryTable(p)
	qs1 := qs.Filter("name__exact", "测试关联2")
	qs2 := qs.Filter("name__iexact", "测试关联")     // 不带%的 like查询
	qs3 := qs.Filter("name__contains", "测试关联a")  // LIKE BINARY 带% %的 like查询 大小写敏感
	qs4 := qs.Filter("name__icontains", "测试关联a") // LIKE 带% %的 like查询 大小写不敏感
	qs5 := qs.Filter("name__startswith", "测试关")  // LIKE BINARY 后置% 的 like查询 大小写敏感
	qs6 := qs.Filter("name__endswith", "a5")     // LIKE BINARY 前置% 的 like查询 大小写敏感
	var maps []models.Profile
	qs1.All(&maps)
	c.Ctx.WriteString(fmt.Sprint(maps))
	var maps1 []models.Profile
	qs2.All(&maps1)
	c.Ctx.WriteString(fmt.Sprint(maps1))
	var maps2 []models.Profile
	qs3.All(&maps2)
	c.Ctx.WriteString(fmt.Sprint(maps2))
	var maps3 []models.Profile
	qs4.All(&maps3)
	c.Ctx.WriteString(fmt.Sprint(maps3))
	var maps4 []models.Profile
	qs5.All(&maps4)
	c.Ctx.WriteString(fmt.Sprint(maps4))
	var maps5 []models.Profile
	qs6.All(&maps5)
	c.Ctx.WriteString(fmt.Sprint(maps5))
}

// 高级查询
func (c *OrmPracticeController) FilterAndExclude() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	p := new(models.Profile)
	qs := c.o.QueryTable(p)
	qs = qs.Filter("name__contains", "关联").Filter("id__gt", 3).Exclude("id", 4)
	var maps []models.Profile
	qs.All(&maps)
	c.Ctx.WriteString(fmt.Sprint(maps))
}

// setCond 自定义查询
func (c *OrmPracticeController) SetCondition() {
	c.o = orm.NewOrm()
	c.o.Using("casbin")
	p := new(models.Profile)
	qs := c.o.QueryTable(p)
	cond := orm.NewCondition()
	cond1 := cond.And("name__isnull", false).AndNot("id", 3)
	qs = qs.SetCond(cond1)
	var maps []models.Profile
	qs.All(&maps)
	c.Ctx.WriteString(fmt.Sprint(maps))
}
