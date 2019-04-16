package controllers

import (
	"baisheng/models"
	"fmt"
	"github.com/astaxie/beego/validation"
)

type ConfirmController struct {
	BaseController
}

//参数过滤
func (this *ConfirmController)filterParams(confirm models.Confirm) {
	//表单验证
	valid := validation.Validation{}

	valid.Required(confirm.AdminId, "管理员Id").Message("不能为空")
	valid.Required(confirm.StoreId, "餐厅编号").Message("不能为空")

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

func (this *ConfirmController)ConfirmList() {

	var confirm	models.Confirm
	//新店 IE 关店 转加盟  完成 准备

	this.Data["confirmList"] , _ = confirm.GetConfirmList()

fmt.Println(this.Data["confirmList"])
	this.Data["storeStatus"] =  GetStatusList()
	this.SetTpl("base/layout_page.html","confirm/list.html")
}

//餐厅信息
func (this *ConfirmController)AddConfirm() {

	switch this.requestMethod {
		case "GET":
			var store	models.PublicStore
			this.Data["storeList"] 	 , _ = store.GetStoreList()
			this.Data["titleName"]  = "添加调试信息"
			this.Data["statusList"] = GetStatusList()
			this.SetTpl("base/layout_page.html","confirm/add.html")

		case "POST":

			var confirm models.Confirm

			if err := this.ParseForm(&confirm); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}

			confirm.AdminId = this.adminInfo.Id
			//校验必填参数
			this.filterParams(confirm)

			if err := confirm.InsertOrUpdate();err != nil{
				this.ReturnJson(-2,err.Error(),nil)
			}
			this.ReturnJson(0,"添加成功",nil)
	}
}

//编辑餐厅信息
func (this *ConfirmController)EditStore() {

	switch this.requestMethod {
		case "GET":
			var  store models.Store

			storeId := this.GetString("storeId")
			if storeId == ""{
				this.Abort("404")
			}
			this.Data["statusList"] 	=  GetStatusList()
			this.Data["titleName"] 		= "编辑餐厅信息"
			this.Data["storeInfo"] 		= store.PublicStore.GetStoreInfo(storeId)
			this.Data["newStoreInfo"] 	= store.NewStore.GetNewStoreInfo(storeId)
			this.Data["ieStoreInfo"] 	= store.IeStore.GetIeStoreInfo(storeId)
			this.Data["closeStoreInfo"] = store.CloseStore.GetCloseStoreInfo(storeId)
			this.SetTpl("base/layout_page.html","store/edit.html")

		case "POST":
			var store 	 models.Confirm
			if err := this.ParseForm(&store); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(store)
			if err := store.InsertOrUpdate(); err != nil{
				this.ReturnJson(-1,err.Error(),nil)
			}
			this.ReturnJson(0,"修改成功",nil)
	}
}



//软删除 餐厅信息
func (this *ConfirmController)DeleteStore() {

	var store models.PublicStore
	if err := this.ParseForm(&store); err != nil {
		this.ReturnJson(-1,err.Error(),nil)
	}
	//校验必填参数
	if  this.GetString("storeId") == ""{
		this.ReturnJson(-1,"餐厅ID不能为空",nil)
	}
	err := store.DeleteStore()
	if err != nil{
		this.ReturnJson(-1,err.Error(),nil)
	}
	this.ReturnJson(0,"删除成功",nil)
}

//标记店为特别关注
func (this *ConfirmController)SignStore() {
	var store models.PublicStore

	if err := this.ParseForm(&store); err != nil {
		this.ReturnJson(-1,err.Error(),nil)
	}

	if  this.GetString("storeId") == ""{
		this.ReturnJson(-1,"餐厅ID不能为空",nil)
	}
	//校验必填参数
	if  this.GetString("signFlag") == ""{
		this.ReturnJson(-2,"标识不能为空",nil)
	}

	err := store.SignStore()
	if err != nil{
		this.ReturnJson(-1,err.Error(),nil)
	}
	this.ReturnJson(0,"操作成功",nil)
}
