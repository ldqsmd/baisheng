package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/validation"
)

type LoginController struct {
	BaseController
}

func (this *LoginController)Login() {

	switch this.requestMethod {
		case "GET":
			this.TplName = "login/login.html"
		case "POST":
			var admin models.Admin

			if err := this.ParseForm(&admin); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//表单验证
			valid := validation.Validation{}
			valid.Required(admin.Account, "账号").Message("不能为空")
			valid.Required(admin.Password, "密码").Message("不能为空")

			if valid.HasErrors() {
				for _, err := range valid.Errors {
					this.ReturnJson(-1,err.Key+err.Message,nil)
				}
			}
			err :=  admin.Login()
			if err != nil{
				this.ReturnJson(-1,"登录失败,账号或密码错误！",nil)
			}

			this.ReturnJson(0,"登录成功",nil)
	}

}

