package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type StoreSystemDeviceController struct {
	BaseController
}

//参数过滤
func (this *StoreSystemDeviceController)filterParams(params *models.StoreSystemDevice) {
	//表单验证
	valid := validation.Validation{}

	valid.Required(params.StoreId, "餐厅ID").Message("不能为空")
	valid.Required(params.SystemDeviceId, "系统设备ID").Message("不能为空")

	if this.requestMethod == "POST" {
		valid.Required(params.Price, "单价").Message("不能为空")
		valid.Required(params.Number, "数量").Message("不能为空")
	}

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			this.ReturnJson(-1,err.Key+err.Message,nil)

			//if this.IsAjax(){
			//	this.ReturnJson(-1,err.Key+err.Message,nil)
			//}else{
			//	this.Data["error"] = err.Key+err.Message
			//	this.Abort("404")
			//}
		}
	}
}

//餐厅的设备列表
func (this *StoreSystemDeviceController)List() {

	var ssd			models.StoreSystemDevice
	var store		models.Store

	if err := this.ParseForm(&store); err != nil {
		this.ReturnJson(-1,err.Error(),nil)
	}

	this.Data["storeInfo"]  = store.PublicStore.GetStoreInfo(strconv.Itoa(store.StoreId))
	this.Data["storeSystemDeviceList"],_ = ssd.GetStoreSystemDeviceList(store.StoreId)
	this.Data["storeId"]  		= store.StoreId
	this.SetTpl("base/layout_page.html","storeSystemDevice/list.html")

}

//添加报价
func (this *StoreSystemDeviceController)Add() {

	switch this.requestMethod {

	case "GET":
		var ssd	models.StoreSystemDevice
		if err := this.ParseForm(&ssd); err != nil {
			this.ReturnJson(-1,err.Error(),nil)
		}
		this.filterParams(&ssd)
		this.Data["ssdInfo"]  = ssd
		this.SetTpl("base/layout_page.html","storeSystemDevice/add.html")

	case "POST":
			var ssd	models.StoreSystemDevice
			if err := this.ParseForm(&ssd); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&ssd)


		if err := ssd.InsertOrUpdate();err != nil{
			this.ReturnJson(-2,err.Error(),nil)
		}
		this.ReturnJson(0,"添加成功",nil)
	}
}



//编辑
func (this *StoreSystemDeviceController)Edit() {

	switch this.requestMethod {

	case "GET":
		var ssd	models.StoreSystemDevice
		if err := this.ParseForm(&ssd); err != nil {
			this.ReturnJson(-1,err.Error(),nil)
		}
		this.Data["ssdInfo"]  = ssd.GetInfo(ssd.StoreSystemDeviceId)
		this.SetTpl("base/layout_page.html","storeSystemDevice/edit.html")

	case "POST":
			var ssd	models.StoreSystemDevice
			if err := this.ParseForm(&ssd); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&ssd)
			if err := ssd.InsertOrUpdate();err != nil{
				this.ReturnJson(-2,err.Error(),nil)
			}
			this.ReturnJson(0,"修改成功",nil)
	}
}


