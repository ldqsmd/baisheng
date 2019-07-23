package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/validation"
	"time"
)

type NoticeController struct {
	BaseController
}

func (this *NoticeController)filterParams(p models.Notice)  {
	//表单验证
	valid := validation.Validation{}


	valid.Required(p.Email, "收件人").Message("不能为空")
	valid.Required(p.Subject, "主题").Message("不能为空")
	valid.Required(p.Body, "详情").Message("不能为空")
	valid.Required(p.NoticeTime, "提醒时间").Message("不能为空")

	if p.NoticeTime < time.Now().Format("2006-01-02 15:04:05"){
		this.ReturnJson(-1,"时间不能小于当前时间",nil)
	}

	if this.actionName == "edit"{
		valid.Required(p.Id, "通知ID").Message("不能为空")
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


func (this *NoticeController)List() {
	var notice models.Notice
	list,total := notice.List()
	this.Data["list"]  = list
	this.Data["total"] = total
	this.SetTpl("base/layout_page.html","notice/list.html")
}

func (this *NoticeController)Add() {

	switch this.requestMethod {
		case "GET":
			this.SetTpl("base/layout_page.html","notice/add.html")
		case "POST":
			var notice models.Notice
			if err := this.ParseForm(&notice); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}

			//校验必填参数
			this.filterParams(notice)

			if  err := notice.InsertOrUpdate() ; err != nil{
				this.ReturnJson(-1,err.Error(),nil)
			}else{
				this.ReturnJson(0,"添加成功",nil)
			}
	}
}

func (this *NoticeController)Edit() {

	switch this.requestMethod {
		case "GET":
			var  noitce models.Notice

			if data :=  noitce.Detail(this.GetString("id")) ;data.Id ==0{
				this.ReturnJson(-2,"信息不存在",nil)
			}else{
				this.Data["data"] = data
				this.SetTpl("base/layout_page.html","notice/edit.html")
			}
		case "POST":
			var  noitce models.Notice
			if err := this.ParseForm(&noitce); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(noitce)
			if err := noitce.InsertOrUpdate() ;err != nil{
				this.ReturnJson(-2,err.Error(),nil)
			}else{
				this.ReturnJson(0,"修改成功",nil)
			}
	}
}

func (this *NoticeController)Del() {
	var notice  models.Notice
	if err := this.ParseForm(&notice); err != nil {
		this.ReturnJson(-1,err.Error(),nil)
	}

	if notice.Id == 0 {
		this.ReturnJson(-1,"提醒ID不能我空",nil)
	}
	if err := notice.Delete(); err != nil{
		this.ReturnJson(-1,err.Error(),nil)
	}
	this.ReturnJson(0,"删除成功",nil)
}


