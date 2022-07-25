package controllers

import "github.com/beego/beego/v2/server/web"

type DetailController struct {
	web.Controller
}

func (c *DetailController) URLMapping() {
	c.Mapping("Information", c.Information)
	c.Mapping("Construction", c.Construction)
}

// @router /information/:key [get]
func (c *DetailController) Information() {
	// bee.exe generate routers
	c.Ctx.WriteString("This is the information")
}

// @router /construction/:id [get]
func (c *DetailController) Construction() {
	c.Ctx.WriteString("This is the construction")
}
