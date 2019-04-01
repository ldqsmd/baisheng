package routers

import (
	"baisheng/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.HomeController{},"GET:Index")
	beego.Router("/index_v1",&controllers.HomeController{},"GET:Content")

	//admin
	beego.Router("/adminList",&controllers.AdminController{},"GET:AdminList")
	beego.Router("/addAdmin",&controllers.AdminController{},"GET:AddAdmin")

}
