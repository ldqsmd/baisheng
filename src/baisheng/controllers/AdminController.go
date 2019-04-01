package controllers

import (
	"baisheng/models"
)

type AdminController struct {
	BaseController
}

func (this *AdminController)AdminList() {

	var admin models.Admin
	list,total := admin.GetList()

	this.Data["list"]  = list
	this.Data["total"] = total

	this.TplName  = "admin/list.html"
}


func (this *AdminController)AddAdmin() {

	this.TplName  = "admin/add.html"
}

