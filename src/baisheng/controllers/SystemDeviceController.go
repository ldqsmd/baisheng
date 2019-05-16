package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type SystemDeviceController struct {
	BaseController
}

//参数过滤
func (this *SystemDeviceController)filterParams(params *models.SystemDevice) {
	//表单验证
	valid := validation.Validation{}

	valid.Required(params.SystemId, "系统ID").Message("不能为空")
	valid.Required(params.DeviceId, "设备ID").Message("不能为空")

	if valid.HasErrors() {
		for _, err := range valid.Errors {

			if this.IsAjax(){
				this.ReturnJson(-1,err.Key+err.Message,nil)
			}else{
				this.Data["error"] = err.Key+err.Message
				this.Abort("404")
			}
		}
	}
}

//列表
func (this *SystemDeviceController)List() {

	var systemDevice	models.SystemDevice
	var system 	models.System

	systemId ,_ := strconv.Atoi(this.GetString("systemId"))
	systemDeviceList, _ := systemDevice.GetSystemDeviceList(systemId)
	this.Data["systemDeviceList"]  = systemDeviceList
	this.Data["systemId"]  = systemId
	this.Data["systemList"] , _ = system.GetList()
	this.SetTpl("base/layout_page.html","systemDevice/list.html")

}

//添加
func (this *SystemDeviceController)Edit() {

	switch this.requestMethod {
		case "POST":

			var systemDevice	models.SystemDevice
			if err := this.ParseForm(&systemDevice); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&systemDevice)
			if err := systemDevice.InsertOrUpdate();err != nil{
				this.ReturnJson(-2,err.Error(),nil)
			}
			this.ReturnJson(0,"修改成功",nil)
	}
}
//
////编辑
//func (this *SystemDeviceController)Edit() {
//
//	switch this.requestMethod {
//		case "GET":
//			if systemId := this.GetString("systemId"); systemId == ""{
//				this.Abort("404")
//			}else{
//				var system  models.System
//				if systemInfo := system.GetInfo(systemId);systemInfo.SystemId == 0 {
//					this.Abort("404")
//				}else{
//					this.Data["systemInfo"]  =  systemInfo
//					this.SetTpl("base/layout_page.html","system/edit.html")
//				}
//			}
//
//		case "POST":
//			var system  models.System
//			if err := this.ParseForm(&system); err != nil {
//				this.ReturnJson(-1,err.Error(),nil)
//			}
//			//校验必填参数
//			this.filterParams(&system)
//			if err := system.InsertOrUpdate(); err != nil{
//				this.ReturnJson(-1,err.Error(),nil)
//			}else{
//				this.ReturnJson(0,"修改成功",nil)
//			}
//	}
//}
//
