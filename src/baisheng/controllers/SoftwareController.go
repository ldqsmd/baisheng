package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type SoftwareController struct {
	BaseController
}

//参数过滤
func (this *SoftwareController)filterParams(software models.Software) {
	//表单验证
	valid := validation.Validation{}


	valid.Required(software.SoftwareName, "软件名称").Message("不能为空")
	valid.Required(software.Version, "软件版本").Message("不能为空")
	valid.Required(software.ConfirmId, "confimId").Message("不能为空")

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

func (this *SoftwareController)SoftwareList() {

	var software	models.Software
	confirmId := this.GetString("confirmId")
	if confirmId == "" {
		this.Abort("404")
	}
	this.Data["titleName"]  = "调试餐厅"
	this.Data["confirmId"]  = confirmId
	this.Data["softwareList"] , _ = software.GetSoftwareList(confirmId)
	this.SetTpl("base/layout_page.html","software/list.html")
}

//餐厅信息
func (this *SoftwareController)AddSoftware() {

	switch this.requestMethod {
		case "GET":
			var store	models.PublicStore
			this.Data["storeList"] 	 , _ = store.GetStoreList()
			this.Data["titleName"]  = "添加调试信息"
			this.Data["statusList"] = GetStatusList()
			this.SetTpl("base/layout_page.html","confirm/add.html")

		case "POST":

			var software models.Software

			if err := this.ParseForm(&software); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(software)

			software.AdminId  = this.adminInfo.Id
			if err := software.InsertOrUpdate();err != nil{
				this.ReturnJson(-2,err.Error(),nil)
			}
			this.ReturnJson(0,"添加成功",nil)
	}
}

//编辑餐厅信息
func (this *SoftwareController)EditSoftware() {

	switch this.requestMethod {
		case "GET":
			var  confirm	models.Confirm
			var  store		models.PublicStore

			confirmId := this.GetString("confirmId")
			if confirmId == ""{
				this.Abort("404")
			}

			this.Data["storeList"] 	 , _ = store.GetStoreList()
			this.Data["confirmInfo"] , _ =  confirm.GetConfirmInfo(confirmId)
			this.Data["titleName"] 		= "编辑调试信息"
			this.SetTpl("base/layout_page.html","confirm/edit.html")

		case "POST":
			var confirm  models.Software
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
func (this *SoftwareController)DelSoftware() {

}


//标记店为特别关注
func (this *SoftwareController)SignSoftware() {
	var software models.Software

	softwareId := this.GetString("softwareId")
	if  softwareId == ""{
		this.ReturnJson(-1,"软件ID",nil)
	}
	software.SoftwareId,_ = strconv.Atoi(softwareId)
	err := software.SignSoftware()
	if err != nil{
		this.ReturnJson(-1,err.Error(),nil)
	}
	this.ReturnJson(0,"操作成功",nil)
}