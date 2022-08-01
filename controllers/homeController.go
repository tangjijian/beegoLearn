package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"myproject/models"
	"strconv"
)

type HomeController struct {
	web.Controller
}

func (c *HomeController) Content() {
	id, _ := c.GetInt("id")
	username := c.GetString("username")

	c.Ctx.WriteString(strconv.Itoa(id) + username)

}
func (c *HomeController) GetForm() {
	models.CreateTable()
	type User struct {
		Id       int    `form:"_"`
		Username string `form:"username"`
		age      int    `form:"age"`
	}
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		c.Ctx.WriteString(err.Error())
	}
	c.Ctx.WriteString(fmt.Sprint(u))
}

func (c *HomeController) ReceiveBody() {

	var ob models.User

	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &ob); err == nil {
		objectid := models.AddOne(ob)
		c.Data["json"] = "{\"ObjectId\":\"" + objectid + "\"}"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
