package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/validation"
)

type ModelDeviceController struct {
	BaseController
}

//参数过滤
func (this *ModelDeviceController)filterParams(params *models.ModelDevice) {
	//表单验证
	valid := validation.Validation{}

	valid.Required(params.ModelId, "模板ID").Message("不能为空")
	valid.Required(params.DeviceName, "设备名称").Message("不能为空")
	valid.Required(params.Position, "位置").Message("不能为空")
	valid.Required(params.Purpose, "用途").Message("不能为空")
	valid.Required(params.DeviceModel, "设备型号").Message("不能为空")
	valid.Required(params.ScrapModel, "报废型号").Message("不能为空")

	if _, _,err := this.GetFile("file-devicePic");err == nil{
		filePath,err := this.UpFileTable("file-devicePic",2)
		if err != nil {
			this.ReturnJson(-1,"file-devicePic"+":"+err.Error(),nil)
		}else{
			params.DevicePic = filePath
		}
	}

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

func (this *ModelDeviceController)List() {

	var device	models.ModelDevice

	if modelId := this.GetString("modelId"); modelId == ""{
		this.Abort("404")
	}else{
		this.Data["modelId"]    =  modelId
		this.Data["titleName"]  = "设备列表"
		this.Data["deviceList"] , _ = device.GetDeviceList(modelId)
		this.SetTpl("base/layout_page.html","modelDevice/list.html")
	}

}

//餐厅信息
func (this *ModelDeviceController)Add() {

	switch this.requestMethod {
		case "GET":

			if modelId := this.GetString("modelId"); modelId == ""{
				this.Abort("404")
			}else{
				this.Data["modelId"]    =  modelId
				this.SetTpl("base/layout_page.html","modelDevice/add.html")
			}
		case "POST":
			var device models.ModelDevice
			if err := this.ParseForm(&device); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&device)

			if err := device.InsertOrUpdate();err != nil{
				this.ReturnJson(-2,err.Error(),nil)
			}
			this.ReturnJson(0,"添加成功",nil)
	}
}

//编辑餐厅信息
func (this *ModelDeviceController)Edit() {

	switch this.requestMethod {
		case "GET":
			if deviceId := this.GetString("deviceId"); deviceId == ""{
				this.Abort("404")
			}else{
				var device models.ModelDevice
				if deviceInfo := device.GetDeviceDetail(deviceId);deviceInfo.DeviceId == 0 {
					this.Abort("404")
				}else{
					this.Data["deviceInfo"] =  deviceInfo
					this.Data["titleName"] 	= "编辑设备信息"
					this.SetTpl("base/layout_page.html","modelDevice/edit.html")
				}
			}

		case "POST":
			var device models.ModelDevice
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


func (this *ModelDeviceController)Del() {

	var device  models.ModelDevice
		if err := this.ParseForm(&device); err != nil {
			this.ReturnJson(-1,err.Error(),nil)
		}
	if device.DeviceId == 0 {
		this.ReturnJson(-1,"设备id不能我空",nil)
	}
	if err := device.Delete(); err != nil{
		this.ReturnJson(-1,err.Error(),nil)
	}
	this.ReturnJson(0,"删除成功",nil)
}

