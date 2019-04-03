package controllers

import (
	"baisheng/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type AdminController struct {
	BaseController
}

func (this *AdminController)AdminList() {

	var admin models.Admin
	list,total := admin.GetList()
	this.Data["list"]  = list
	this.Data["total"] = total
	this.SetTpl("base/layout_page.html","admin/list.html")
}

func (this *AdminController)AddAdmin() {

	switch this.requestMethod {
		case "GET":
			this.SetTpl("base/layout_page.html","admin/add.html")
		case "POST":
			var admin models.Admin

			if err := this.ParseForm(&admin); err != nil {
				fmt.Println(err.Error())
			}
			o := orm.NewOrm()
			id, err := o.Insert(&admin)
			if err != nil {
				fmt.Println(err.Error())
			}else{
				fmt.Println(id)
			}
			this.Redirect("/adminList",200)
		}
}


