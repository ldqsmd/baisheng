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
	beego.Router("/page404",&controllers.HomeController{},"GET:Page404")

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

	//confirm
	beego.Router("/confirm/list",&controllers.ConfirmController{},"GET:ConfirmList")
	beego.Router("/confirm/add",&controllers.ConfirmController{},"*:AddConfirm")
	beego.Router("/confirm/edit",&controllers.ConfirmController{},"*:EditConfirm")

	//software
	beego.Router("/software/list",&controllers.SoftwareController{},"GET:SoftwareList")
	beego.Router("/software/edit",&controllers.SoftwareController{},"*:EditSoftware")
	beego.Router("/software/del",&controllers.SoftwareController{},"POST:DelSoftware")
	beego.Router("/software/add",&controllers.SoftwareController{},"*:AddSoftware")
	beego.Router("/software/sign",&controllers.SoftwareController{},"POST:SignSoftware")

	//deviceCheck
	beego.Router("/deviceCheck/list",&controllers.DeviceCheckController{},"GET:CheckList")
	beego.Router("/deviceCheck/add",&controllers.DeviceCheckController{},"*:AddCheck")
	beego.Router("/deviceCheck/edit",&controllers.DeviceCheckController{},"*:EditCheck")
	beego.Router("/deviceCheck/del",&controllers.DeviceCheckController{},"*:DelCheck")
}
