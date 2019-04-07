package controllers

import (
	"baisheng/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type AdminController struct {
	BaseController
}

func (this *AdminController)AdminList() {

	var admin models.Admin
	list,total := admin.GetAdminList()
	this.Data["list"]  = list
	this.Data["total"] = total
	this.SetTpl("base/layout_page.html","admin/list.html")
}


func (this *AdminController)filterParams(admin models.Admin)  {
	//表单验证
	valid := validation.Validation{}

	if this.requestMethod == "AddAdmin"{
		valid.Required(admin.Account, "账号").Message("不能为空")
		valid.Required(admin.Password, "密码").Message("不能为空")
		valid.Required(admin.Email, "e-mail").Message("不能为空")

	}
	if this.actionName == "EditAdmin"{
		//编辑的时候 必填 adminId
		valid.Required(admin.Id, "管理员ID").Message("不能为空")

		if this.requestMethod == "POST"{
			valid.Required(admin.Account, "账号").Message("不能为空")
			valid.Required(admin.Password, "密码").Message("不能为空")
			valid.Required(admin.Email, "e-mail").Message("不能为空")
		}
	}

	if valid.HasErrors() {
		for _, err := range valid.Errors {

			if this.IsAjax(){
				this.ReturnJson(-1,err.Key+err.Message,nil)
			}else{
				this.Data["error"] = err.Key+err.Message
				this.Error404()
			}
		}
	}

}

//添加管理员
func (this *AdminController)AddAdmin() {

	switch this.requestMethod {
		case "GET":
			this.SetTpl("base/layout_page.html","admin/add.html")
		case "POST":
			var admin models.Admin
			if err := this.ParseForm(&admin); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(admin)

			 admin.AddAdmin()
			//if err != nil{
			//	this.ReturnJson(-1,err.Error(),nil)
			//}
			//this.ReturnJson(0,"添加成功",nil)

	}
}

//编辑管理员
func (this *AdminController)EditAdmin() {

	switch this.requestMethod {
		case "GET":
			var  admin models.Admin
			admin.Id,_ = strconv.Atoi(this.GetString("adminId"))
			this.filterParams(admin)
			//err := admin.GetAdminInfo()
			//if  err != nil{
			//	this.Data["error"] = err.Error()
			//	this.Error404()
			//}

		case "POST":

			var admin models.Admin
			if err := this.ParseForm(&admin); err != nil {
				fmt.Println(err.Error())
			}
			//校验必填参数
			this.filterParams(admin)

			err := admin.UpdateAdmin()
			if err != nil{
				this.ReturnJson(-1,err.Error(),nil)
			}
			this.ReturnJson(0,"添加成功",nil)

	}
}

