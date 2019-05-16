package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/validation"
)

type StoreSystemController struct {
	BaseController
}

//参数过滤
func (this *StoreSystemController)filterParams(params *models.StoreSystem) {
	//表单验证
	valid := validation.Validation{}

	valid.Required(params.StoreId, "餐厅ID名称").Message("不能为空")
	valid.Required(params.SystemId, "系统ID").Message("不能为空")

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


//添加
func (this *StoreSystemController)Add() {

	var storeSystem	models.StoreSystem
	if err := this.ParseForm(&storeSystem); err != nil {
		this.ReturnJson(-1,err.Error(),nil)
	}
	//校验必填参数
	this.filterParams(&storeSystem)

	if err := storeSystem.InsertOrUpdate();err != nil{
		this.ReturnJson(-2,err.Error(),nil)
	}
	this.ReturnJson(0,"添加成功",nil)
}

//删除
func (this *StoreSystemController)Del() {

	var storeSystem	models.StoreSystem
	if err := this.ParseForm(&storeSystem); err != nil {
		this.ReturnJson(-1,err.Error(),nil)
	}	//校验必填参数
	this.filterParams(&storeSystem)

	if err := storeSystem.Delete(); err != nil{
		this.ReturnJson(-1,err.Error(),nil)
	}else{
		this.ReturnJson(0,"解除系统成功",nil)
	}

}




