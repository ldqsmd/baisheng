package routers

import (
	"baisheng/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//login
	beego.Router("/login",&controllers.LoginController{},"*:Login")
	beego.Router("/logout",&controllers.LoginController{},"*:Logout")


	beego.Router("/",&controllers.HomeController{},"GET:Index")
	beego.Router("/index_v1",&controllers.HomeController{},"GET:Content")

	//admin
	beego.Router("/admin/list",&controllers.AdminController{},"GET:AdminList")
	beego.Router("/admin/add",&controllers.AdminController{},"*:AddAdmin")
	beego.Router("/admin/edit",&controllers.AdminController{},"*:EditAdmin")

	//store
	beego.Router("/store/list",&controllers.StoreController{},"GET:StoreList")
	beego.Router("/store/add",&controllers.StoreController{},"*:AddStore")
	beego.Router("/store/edit",&controllers.StoreController{},"*:EditStore")
	beego.Router("/store/delete",&controllers.StoreController{},"POST:DeleteStore")
	beego.Router("/store/sign",&controllers.StoreController{},"POST:SignStore")

	//newStore
	beego.Router("/confirm/list",&controllers.ConfirmController{},"GET:ConfirmList")
	beego.Router("/confirm/add",&controllers.ConfirmController{},"*:AddConfirm")

}
