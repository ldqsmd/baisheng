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

	if this.actionName == "EditConfirm"{
		valid.Required(confirm.ConfirmId, "调试ID").Message("不能为空")
	}

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

	confirmList , _ := confirm.GetConfirmList()


	this.Data["confirmList"]  = confirmList
	//新店 IE 关店 转加盟  完成 准备
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
func (this *ConfirmController)EditConfirm() {

	switch this.requestMethod {
		case "GET":
			var  confirm	models.Confirm
			var  store		models.PublicStore

			confirmId := this.GetString("confirmId")
			if confirmId == ""{
				this.Abort("404")
			}

		fmt.Println(this.Data["confirmInfo"] )
			this.Data["storeList"] 	 , _ = store.GetStoreList()
			this.Data["confirmInfo"] , _ =  confirm.GetConfirmInfo(confirmId)
			this.Data["titleName"] 		= "编辑调试信息"
			this.SetTpl("base/layout_page.html","confirm/edit.html")

		case "POST":
			var confirm  models.Confirm
			if err := this.ParseForm(&confirm); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			confirm.AdminId = this.adminInfo.Id

			this.filterParams(confirm)
			if err := confirm.InsertOrUpdate(); err != nil{
				this.ReturnJson(-1,err.Error(),nil)
			}
			this.ReturnJson(0,"修改成功",nil)
	}
}

