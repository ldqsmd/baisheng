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
	beego.Router("/admin/forbid",&controllers.AdminController{},"POST:ForbidAdmin")

	//store
	beego.Router("/store/list",&controllers.StoreController{},"GET:StoreList")
	beego.Router("/store/add",&controllers.StoreController{},"*:AddStore")
	beego.Router("/store/edit",&controllers.StoreController{},"*:EditStore")
	beego.Router("/store/delete",&controllers.StoreController{},"POST:DeleteStore")
	beego.Router("/store/sign",&controllers.StoreController{},"POST:SignStore")

	//confirmSignSoftware
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

	//model
	beego.Router("/model/list",&controllers.ModelController{},"GET:ModelList")
	beego.Router("/model/add",&controllers.ModelController{},"POST:AddModel")
	beego.Router("/model/edit",&controllers.ModelController{},"*:EditModel")
	beego.Router("/model/del",&controllers.ModelController{},"*:DelModel")

	//modelDevice
	beego.Router("/modelDevice/list",&controllers.ModelDeviceController{},"GET:List")
	beego.Router("/modelDevice/add",&controllers.ModelDeviceController{},"*:Add")
	beego.Router("/modelDevice/edit",&controllers.ModelDeviceController{},"*:Edit")
	beego.Router("/modelDevice/del",&controllers.ModelDeviceController{},"POST:Del")

	//device
	beego.Router("/device/list",&controllers.DeviceController{},"GET:List")
	beego.Router("/device/add",&controllers.DeviceController{},"*:Add")
	beego.Router("/device/edit",&controllers.DeviceController{},"*:Edit")
	beego.Router("/device/del",&controllers.DeviceController{},"POST:Del")

	//system
	beego.Router("/system/list",&controllers.SystemController{},"GET:List")
	beego.Router("/system/add",&controllers.SystemController{},"*:Add")
	beego.Router("/system/edit",&controllers.SystemController{},"*:Edit")
	beego.Router("/system/changeStatus",&controllers.SystemController{},"POST:ChangeStatus")

	//SystemDeviceController
	beego.Router("/systemDevice/list",&controllers.SystemDeviceController{},"GET:List")
	beego.Router("/systemDevice/edit",&controllers.SystemDeviceController{},"*:Edit")

	//StoreSystemController
	beego.Router("/storeSystem/add",&controllers.StoreSystemController{},"POST:Add")
	beego.Router("/storeSystem/del",&controllers.StoreSystemController{},"POST:Del")

	//StoreSystemDeviceController
	beego.Router("/storeSystemDevice/list",&controllers.StoreSystemDeviceController{},"GET:List")
	beego.Router("/storeSystemDevice/edit",&controllers.StoreSystemDeviceController{},"*:Edit")
	beego.Router("/storeSystemDevice/add",&controllers.StoreSystemDeviceController{},"*:Add")

	//NoticeController
	beego.Router("/notice/list",&controllers.NoticeController{},"GET:List")
	beego.Router("/notice/edit",&controllers.NoticeController{},"*:Edit")
	beego.Router("/notice/add",&controllers.NoticeController{},"*:Add")
	beego.Router("/notice/del",&controllers.NoticeController{},"*:Del")
}
