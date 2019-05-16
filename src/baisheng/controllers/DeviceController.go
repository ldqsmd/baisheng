package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/validation"
)

type DeviceController struct {
	BaseController
}

//参数过滤
func (this *DeviceController)filterParams(params *models.Device) {
	//表单验证
	valid := validation.Validation{}

	if this.actionName == "Edit"{
		valid.Required(params.DeviceId, "设备ID").Message("不能为空")
	}
	valid.Required(params.DeviceName, "设备名称").Message("不能为空")

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
func (this *DeviceController)List() {

	var device	models.Device
	this.Data["deviceList"] , _ = device.GetDeviceList()
	this.SetTpl("base/layout_page.html","device/list.html")
}

//添加
func (this *DeviceController)Add() {

	switch this.requestMethod {
		case "GET":
			this.SetTpl("base/layout_page.html","device/add.html")

		case "POST":

			var device	models.Device
			if err := this.ParseForm(&device); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&device)

			if err := device.InsertOrUpdate();err != nil{
				this.ReturnJson(-2,err.Error(),nil)
			}else{
				this.ReturnJson(0,"添加成功",nil)
			}
	}
}

//编辑
func (this *DeviceController)Edit() {

	switch this.requestMethod {
		case "GET":
			if deviceId := this.GetString("deviceId"); deviceId == ""{
				this.Abort("404")
			}else{
				var device models.Device
				if deviceInfo := device.GetDeviceInfo(deviceId);deviceInfo.DeviceId == 0 {
					this.Abort("404")
				}else{
					this.Data["deviceInfo"]  =  deviceInfo
					this.SetTpl("base/layout_page.html","device/edit.html")
				}
			}

		case "POST":
			var device models.Device
			if err := this.ParseForm(&device); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&device)
			if err := device.InsertOrUpdate(); err != nil{
				this.ReturnJson(-1,err.Error(),nil)
			}else{
				this.ReturnJson(0,"修改成功",nil)
			}
	}
}

//删除
func (this *DeviceController)Del() {

	var device models.Device
	if err := this.ParseForm(&device); err != nil {
		this.ReturnJson(-1,err.Error(),nil)
	}
	if device.DeviceId == 0 {
		this.ReturnJson(-1,"设备ID不能为空",nil)
	}
	if err := device.DeleteDevice(); err != nil{
		this.ReturnJson(-1,err.Error(),nil)
	}else{
		this.ReturnJson(0,"删除成功",nil)
	}

}

