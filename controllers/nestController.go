package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type NestPreparer interface {
	NestPrepare()
}

// baseRouter implemented global settings for all other routers.
type baseController struct {
	beego.Controller

	isLogin bool
}

// Prepare implemented Prepare method for baseRouter.
func (c *baseController) Prepare() {

	// page start time
	c.Data["PageStartTime"] = time.Now()

	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
