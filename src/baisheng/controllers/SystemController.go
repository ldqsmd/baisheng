package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/validation"
)

type SystemController struct {
	BaseController
}

//参数过滤
func (this *SystemController)filterParams(params *models.System) {
	//表单验证
	valid := validation.Validation{}

	if this.actionName == "Edit"{
		valid.Required(params.SystemId, "系统ID").Message("不能为空")
	}
	valid.Required(params.SystemName, "系统名称").Message("不能为空")

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
func (this *SystemController)List() {

	var system	models.System
	this.Data["systemList"] , _ = system.GetList()
	this.SetTpl("base/layout_page.html","system/list.html")
}

//添加
func (this *SystemController)Add() {

	switch this.requestMethod {
		case "GET":
			this.SetTpl("base/layout_page.html","system/add.html")

		case "POST":

			var system	models.System
			if err := this.ParseForm(&system); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&system)

			if err := system.InsertOrUpdate();err != nil{
				this.ReturnJson(-2,err.Error(),nil)
			}
			this.ReturnJson(0,"添加成功",nil)
	}
}

//编辑
func (this *SystemController)Edit() {

	switch this.requestMethod {
		case "GET":
			if systemId := this.GetString("systemId"); systemId == ""{
				this.Abort("404")
			}else{
				var system  models.System
				if systemInfo := system.GetInfo(systemId);systemInfo.SystemId == 0 {
					this.Abort("404")
				}else{
					this.Data["systemInfo"]  =  systemInfo
					this.SetTpl("base/layout_page.html","system/edit.html")
				}
			}

		case "POST":
			var system  models.System
			if err := this.ParseForm(&system); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&system)
			if err := system.InsertOrUpdate(); err != nil{
				this.ReturnJson(-1,err.Error(),nil)
			}else{
				this.ReturnJson(0,"修改成功",nil)
			}
	}
}

//更改状态
func (this *SystemController)ChangeStatus() {

	var system  models.System
	if err := this.ParseForm(&system); err != nil {
		this.ReturnJson(-1,err.Error(),nil)
	}
	if system.SystemId == 0 {
		this.ReturnJson(-1,"系统ID不能为空",nil)
	}

	if system.Status == "" {
		this.ReturnJson(-1,"状态ID不能为空",nil)
	}

	if err := system.ChangeStatus(); err != nil{
		this.ReturnJson(-1,err.Error(),nil)
	}else{
		this.ReturnJson(0,"操作成功",nil)
	}

}




