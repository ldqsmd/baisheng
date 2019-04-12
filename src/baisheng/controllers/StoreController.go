package controllers

import (
	"baisheng/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"reflect"
	"strings"
)

type StoreController struct {
	BaseController
}

//参数过滤
func (this *StoreController)filterParams(store *models.Store) {
	//表单验证
	valid := validation.Validation{}

	valid.Required(store.StoreName, "餐厅名称").Message("不能为空")
	valid.Required(store.Number, "餐厅编号").Message("不能为空")
	valid.Required(store.Brand, "餐厅品牌").Message("不能为空")
	valid.Required(store.Status, "餐厅状态").Message("不能为空")


	//新店
	if store.Status == 1 {
		//新店：必选 邮箱申请、派单时间 （三星、LG 、HP）
		//IE:必选报废评估、派单时间 （HP、三星、LG）
		//关店：不用了
		valid.Required(store.ApplyEmailTime, "申请邮箱时间").Message("不能为空")
		valid.Required(store.NewDmbDispatchTime, "DMB派单时间").Message("不能为空")
		valid.Required(store.CallNumTime, "叫号屏派单时间").Message("不能为空")
	}
	//IE
	if store.Status == 2 {
		//新店：必选 邮箱申请、派单时间 （三星、LG 、HP）
		//IE:必选报废评估、派单时间 （HP、三星、LG）
		//关店：不用了

		valid.Required(store.OpenImacTime, "开店IMAC派单时间").Message("不能为空")
		valid.Required(store.CloseImacTime, "关店IMAC派单时间").Message("不能为空")
		valid.Required(store.DmbUninstallTime, "DMB拆除派单时间").Message("不能为空")
		valid.Required(store.DmbInstallTime, "DMB安装派单时间").Message("不能为空")
		valid.Required(store.SamsungIntallTime, "DMB三星安装派单时间").Message("不能为空")
		valid.Required(store.SamsungUnintallTime, "DMB三星拆除派单时间").Message("不能为空")
	}

	if store.Status == 3 {
		if this.actionName == "EditStore" {
			valid.Required(store.StoreId, "餐厅ID").Message("不能为空")
			valid.Required(store.IeStoreId, "IEID").Message("不能为空")
		}
		var closeStoreInfo models.CloseStore
		//判断是否有文件上传
		uploadStr := this.GetString("uploadList")
		if  uploadStr != ""{
			uploadList := FilterStrSlice(strings.Split(uploadStr,","))
			//遍历结构体赋值
			storeType  		:= reflect.TypeOf(closeStoreInfo)
			storeTypeVal  	:= reflect.New(storeType)
			//遍历上传文件
			for _,formFile  := range uploadList{

				if filePath,err := this.UpFileTable(formFile); err != nil {
					this.ReturnJson(-1,formFile+":"+err.Error(),nil)
				}else{
					for k := 0; k < storeType.NumField(); k++ {
						if   strings.ToLower(storeType.Field(k).Name) == strings.ToLower(formFile){
							storeTypeVal.Elem().Field(k).SetString(filePath)
						}
					}
				}
			}
			closeStoreInfo = storeTypeVal.Elem().Interface().(models.CloseStore)
			store.DeviceLcTogoTable = closeStoreInfo.DeviceLcTogoTable
			store.DeviceHpTogoTable = closeStoreInfo.DeviceHpTogoTable
			store.ToLcPropertyTable = closeStoreInfo.ToLcPropertyTable
			store.PropertyDestroyTable = closeStoreInfo.PropertyDestroyTable
			store.DeviceReturnLcApply = closeStoreInfo.DeviceReturnLcApply
		}

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

func (this *StoreController)StoreList() {

	var store	models.PublicStore
	var network 	= map[int]string{1:"餐厅宽带",2:"联通",3:"电信",4:"移动",5:"小运营商"}
	var storeStatus = []string{1:"新店",2:"IE",3:"关店",4:"转加盟",5:"完成",6:"准备"}
	var brandList 	= []string{1:"肯德基",2:"必胜客",3:"小肥羊",4:"塔可贝尔",5:"咖啡"}

	list,total := store.GetStoreList()
	//{{/*新店 IE 关店 转加盟  完成 准备*/}}
	this.Data["storeStatus"]=  storeStatus
	this.Data["statusList"] =  GetStatusList()
	this.Data["network"]  	=  network
	this.Data["brandList"]	=  brandList
	this.Data["list"]  		= list
	this.Data["total"] 		= total
	this.SetTpl("base/layout_page.html","store/list.html")
}

//餐厅信息
func (this *StoreController)AddStore() {

	switch this.requestMethod {
		case "GET":
			this.Data["titleName"]  = "添加餐厅信息"
			this.Data["statusList"] = GetStatusList()
			this.SetTpl("base/layout_page.html","store/add.html")

		case "POST":

			var store models.Store
			if err := this.ParseForm(&store); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&store)
			if err := store.InsertOrUpdate();err != nil{
				this.ReturnJson(-2,err.Error(),nil)
			}
			this.ReturnJson(0,"添加成功",nil)
	}
}

//编辑餐厅信息
func (this *StoreController)EditStore() {

	switch this.requestMethod {
		case "GET":
			var  store models.Store

			storeId := this.GetString("storeId")
			if storeId == ""{
				this.Abort("404")
			}
			fmt.Println(store.GetStoreInfo(storeId))

			this.Data["statusList"] 	=  GetStatusList()
			this.Data["titleName"] 		= "编辑餐厅信息"
			this.Data["storeInfo"] 		= store.PublicStore.GetStoreInfo(storeId)
			this.Data["newStoreInfo"] 	= store.NewStore.GetNewStoreInfo(storeId)
			this.Data["ieStoreInfo"] 	= store.IeStore.GetIeStoreInfo(storeId)
			this.Data["closeStoreInfo"] = store.CloseStore.GetCloseStoreInfo(storeId)
			this.SetTpl("base/layout_page.html","store/edit.html")

		case "POST":
			var store 	 models.Store
			if err := this.ParseForm(&store); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(&store)

			if err := store.InsertOrUpdate(); err != nil{
				this.ReturnJson(-1,err.Error(),nil)
			}
			this.ReturnJson(0,"修改成功",nil)
	}
}



//软删除 餐厅信息
func (this *StoreController)DeleteStore() {

	var store models.PublicStore
	if err := this.ParseForm(&store); err != nil {
		this.ReturnJson(-1,err.Error(),nil)
	}
	//校验必填参数
	if  this.GetString("storeId") == ""{
		this.ReturnJson(-1,"餐厅ID不能为空",nil)
	}

	err := store.DeleteStore()
	if err != nil{
		this.ReturnJson(-1,err.Error(),nil)
	}
	this.ReturnJson(0,"删除成功",nil)
}

