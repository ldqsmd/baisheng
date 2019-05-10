package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/validation"
)

type DeviceCheckController struct {
	BaseController
}

//参数过滤
func (this *DeviceCheckController)filterParams(params *models.DeviceCheck) {
	//表单验证
	valid := validation.Validation{}
	params.AdminId  = this.adminInfo.Id

	if this.actionName == "EditCheck"{
		valid.Required(params.CheckId, "审查ID").Message("不能为空")
	}
	valid.Required(params.StoreId, "餐厅").Message("不能为空")
	valid.Required(params.ReplyDate, "确认回复日期").Message("不能为空")

	readyFile 	 := "file-ready_file"
	completeFile := "file-complete_file"

	if _, _,err := this.GetFile(readyFile);err == nil{
		filePath,err := this.UpFileTable(readyFile,1)
		if err != nil {
			this.ReturnJson(-1,readyFile+":"+err.Error(),nil)
		}
		if  params.Status == 0{
			params.Status = 1
		}
		params.ReadyFile = filePath
	}

	if _, _,err := this.GetFile(completeFile);err == nil{
		if  params.Status == 0{
			this.ReturnJson(-1,"请上传初始详情表",nil)
		}

		filePath,err := this.UpFileTable(completeFile,1)
		if err != nil {
			this.ReturnJson(-1,completeFile+":"+err.Error(),nil)
		}
		params.Status = 2
		params.CompleteFile = filePath
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

func (this *DeviceCheckController)CheckList() {

	var checkList	models.CheckList

	this.Data["titleName"]  = "设备审查"
	this.Data["checkList"] , _ = checkList.GetCheckList()
	this.SetTpl("base/layout_page.html","deviceCheck/list.html")
}

//餐厅信息
func (this *DeviceCheckController)AddCheck() {

	switch this.requestMethod {
		case "GET":
			var storeOption 	models.StoreOption
			this.Data["optionList"] 	=  storeOption.GetStoreOption()
			this.Data["titleName"]  = "添加审查信息"
			this.SetTpl("base/layout_page.html","deviceCheck/add.html")

		case "POST":
			var check models.DeviceCheck
			if err := this.ParseForm(&check); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&check)

			if err := check.InsertOrUpdate();err != nil{
				this.ReturnJson(-2,err.Error(),nil)
			}
			this.ReturnJson(0,"添加成功",nil)
	}
}

//编辑餐厅信息
func (this *DeviceCheckController)EditCheck() {

	switch this.requestMethod {
		case "GET":

			if checkId := this.GetString("checkId"); checkId == ""{
				this.Abort("404")
			}else{
				var check models.DeviceCheck
				var storeOption 	models.StoreOption
				if checkInfo := check.GetCheckInfo(checkId);checkInfo.CheckId == 0 {
					this.Abort("404")
				}
				this.Data["optionList"] 	=  storeOption.GetStoreOption()
				this.Data["checkInfo"]  =  check.GetCheckInfo(checkId)
				this.Data["titleName"] 	= "编辑审查信息"
				this.SetTpl("base/layout_page.html","deviceCheck/edit.html")
			}

		case "POST":
			var check models.DeviceCheck
			if err := this.ParseForm(&check); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&check)
			if err := check.InsertOrUpdate(); err != nil{
				this.ReturnJson(-1,err.Error(),nil)
			}
			this.ReturnJson(0,"修改成功",nil)
	}
}



func (this *DeviceCheckController)DelCheck() {

	var check  models.DeviceCheck
		if err := this.ParseForm(&check); err != nil {
			this.ReturnJson(-1,err.Error(),nil)
		}
	if check.CheckId == 0 {
		this.ReturnJson(-1,"审查ID不能我空",nil)
	}
	if err := check.DeleteDeviceCheck(); err != nil{
		this.ReturnJson(-1,err.Error(),nil)
	}
	this.ReturnJson(0,"删除成功",nil)
}

