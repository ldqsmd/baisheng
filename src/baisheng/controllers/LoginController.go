package controllers

import (
	"baisheng/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type LoginController struct {
	BaseController
}

func (this *LoginController)Get() {


	var errCon ErrorController
	errCon.Error404()
	this.Abort("404")
	//this.TplName = "login/login.html"
}


func (this *LoginController)Post() {

	var admin models.Admin

	fmt.Println(this.Input())
	if err := this.ParseForm(&admin); err != nil {
		fmt.Println(err.Error())
	}
	//表单验证
	valid := validation.Validation{}
	valid.Required(admin.Account, "account")
	valid.Required(admin.Password, "password")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			fmt.Println(err.Key, err.Message)
		}
		return
	}

	err := orm.NewOrm().Raw("SELECT account,user_name FROM user WHERE id = ?", 1).QueryRow(&user)


	//查看数据表
	err := orm.NewOrm().QueryTable("admin").Filter("account", admin.Account).Filter("password", admin.Password).One(&admin)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Printf("Not row found")
	}
}

