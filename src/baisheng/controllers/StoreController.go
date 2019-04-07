package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type StoreController struct {
	BaseController
}

func (this *StoreController)StoreList() {

	var store models.Store
	var network 	= map[int]string{1:"餐厅宽带",2:"联通",3:"电信",4:"移动",5:"小运营商"}
	var storeStatus = []string{1:"新店",2:"IE",3:"关店",4:"转加盟",5:"完成",6:"准备"}
	var brandList 	= []string{1:"肯德基",2:"必胜客",3:"小肥羊",4:"塔可贝尔",5:"咖啡"}

	list,total := store.GetStoreList()
	//{{/*新店 IE 关店 转加盟  完成 准备*/}}
	this.Data["storeStatus"]=  storeStatus
	this.Data["network"]  	=  network
	this.Data["brandList"]	=  brandList
	this.Data["list"]  		= list
	this.Data["total"] 		= total
	this.SetTpl("base/layout_page.html","store/list.html")
}

//添加管理员
func (this *StoreController)AddStore() {

	switch this.requestMethod {
		case "GET":
			this.Data["titleName"] ="添加餐厅信息"
			this.SetTpl("base/layout_page.html","store/add.html")
		case "POST":
			var store models.Store
			if err := this.ParseForm(&store); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(store)

			_,err := store.AddStore()
			if err != nil{
				this.ReturnJson(-1,err.Error(),nil)
			}
			this.ReturnJson(0,"添加成功",nil)
	}
}

func (this *StoreController)filterParams(store models.Store) error {
	//表单验证
	valid := validation.Validation{}

	if this.requestMethod == "AddStore"{
		valid.Required(store.StoreName, "餐厅名称").Message("不能为空")
		valid.Required(store.Number, "餐厅编号").Message("不能为空")
		valid.Required(store.Brand, "餐厅品牌").Message("不能为空")

	}
	if this.actionName == "EditStore"{
		//编辑的时候 必填 adminId
		valid.Required(store.Id, "餐厅ID").Message("不能为空")

		if this.requestMethod == "POST"{
			valid.Required(store.StoreName, "餐厅名称").Message("不能为空")
			valid.Required(store.Number, "餐厅编号").Message("不能为空")
			valid.Required(store.Brand, "餐厅品牌").Message("不能为空")
		}
	}

	if this.actionName == "DeleteStore"{
		//编辑的时候 必填 adminId
		valid.Required(store.Id, "餐厅ID").Message("不能为空")
	}

	if valid.HasErrors() {
		for _, err := range valid.Errors {

			if this.IsAjax(){
				this.ReturnJson(-1,err.Key+err.Message,nil)
			}else{
				this.Data["error"] = err.Key+err.Message
				this.Error404()
				return  err
			}
		}
	}
	return nil
}

//编辑餐厅信息
func (this *StoreController)EditStore() {

	switch this.requestMethod {
		case "GET":
			var  store models.Store
			store.Id,_ = strconv.Atoi(this.GetString("storeId"))
			err :=  this.filterParams(store)
			if err != nil{
				this.StopRun()
			}
			store.GetStoreInfo()

			this.Data["titleName"] = "编辑餐厅信息"
			this.Data["storeInfo"] = store
			this.SetTpl("base/layout_page.html","store/edit.html")

		case "POST":
			var store models.Store
			if err := this.ParseForm(&store); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(store)
			err := store.UpdateStore()
			if err != nil{
				this.ReturnJson(-1,err.Error(),nil)
			}
			this.ReturnJson(0,"修改成功",nil)
	}
}



//软删除 餐厅信息
func (this *StoreController)DeleteStore() {

	var store models.Store
	if err := this.ParseForm(&store); err != nil {
		this.ReturnJson(-1,err.Error(),nil)
	}
	//校验必填参数
	this.filterParams(store)

	err := store.DeleteStore()
	if err != nil{
		this.ReturnJson(-1,err.Error(),nil)
	}
	this.ReturnJson(0,"删除成功",nil)
}

