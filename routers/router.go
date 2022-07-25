package routers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"myproject/controllers"
)

func init() {
	fmt.Println("进来了")
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.DetailController{})
}
