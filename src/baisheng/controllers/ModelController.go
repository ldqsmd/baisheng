package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/validation"
)

type ModelController struct {
	BaseController
}

//参数过滤
func (this *ModelController)filterParams(params *models.Model) {
	//表单验证
	valid := validation.Validation{}

	if this.actionName == "EditModel"{
		valid.Required(params.ModelId, "模板ID").Message("不能为空")
	}
	valid.Required(params.ModelName, "模板名称").Message("不能为空")

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

func (this *ModelController)ModelList() {

	var model	models.Model
	this.Data["titleName"]  = "报废模板列表"
	this.Data["modelList"] , _ = model.GetModelList()
	this.SetTpl("base/layout_page.html","model/list.html")
}

//餐厅信息
func (this *ModelController)AddModel() {

	switch this.requestMethod {
		case "GET":
			this.Data["titleName"]  = "添加模板信息"
			this.SetTpl("base/layout_page.html","model/add.html")

		case "POST":
			var model models.Model
			if err := this.ParseForm(&model); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&model)

			if err := model.InsertOrUpdate();err != nil{
				this.ReturnJson(-2,err.Error(),nil)
			}
			this.ReturnJson(0,"添加成功",nil)
	}
}

//编辑餐厅信息
func (this *ModelController)EditModel() {

	switch this.requestMethod {
		case "GET":

			if modelId := this.GetString("modelId"); modelId == ""{
				this.Abort("404")
			}else{
				var model models.Model
				if modelInfo := model.GetModelInfo(modelId);modelInfo.ModelId == 0 {
					this.Abort("404")
				}else{
					this.Data["modelInfo"]  =  modelInfo
					this.Data["titleName"] 	= "编辑模板信息"
					this.SetTpl("base/layout_page.html","model/edit.html")
				}
			}

		case "POST":
			var model models.Model
			if err := this.ParseForm(&model); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&model)
			if err := model.InsertOrUpdate(); err != nil{
				this.ReturnJson(-1,err.Error(),nil)
			}
			this.ReturnJson(0,"修改成功",nil)
	}
}


func (this *ModelController)DelModel() {

	var model  models.Model
		if err := this.ParseForm(&model); err != nil {
			this.ReturnJson(-1,err.Error(),nil)
		}
	if model.ModelId == 0 {
		this.ReturnJson(-1,"模板ID不能我空",nil)
	}
	if err := model.DeleteModel(); err != nil{
		this.ReturnJson(-1,err.Error(),nil)
	}
	this.ReturnJson(0,"删除成功",nil)
}

