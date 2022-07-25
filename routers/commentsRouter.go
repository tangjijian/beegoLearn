package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["myproject/controllers:DetailController"] = append(beego.GlobalControllerRouter["myproject/controllers:DetailController"],
		beego.ControllerComments{
			Method:           "Construction",
			Router:           `/construction/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myproject/controllers:DetailController"] = append(beego.GlobalControllerRouter["myproject/controllers:DetailController"],
		beego.ControllerComments{
			Method:           "Information",
			Router:           `/information/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
