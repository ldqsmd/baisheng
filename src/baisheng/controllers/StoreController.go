package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/orm"
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
	valid.Required(store.SystemId, "餐厅系统").Message("不能为空")

	//新店
	if store.Status == 1 {
		//新店：必选 邮箱申请、派单时间 （三星、LG 、HP）
		//IE:必选报废评估、派单时间 （HP、三星、LG）
		//关店：不用了
		//valid.Required(store.ApplyEmailTime, "申请邮箱时间").Message("不能为空")
		//valid.Required(store.NewDmbDispatchTime, "DMB派单时间").Message("不能为空")
		//valid.Required(store.CallNumTime, "叫号屏派单时间").Message("不能为空")
	}
	//IE
	if store.Status == 2 {
		//新店：必选 邮箱申请、派单时间 （三星、LG 、HP）
		//IE:必选报废评估、派单时间 （HP、三星、LG）
		//关店：不用了

		//valid.Required(store.OpenImacTime, "开店IMAC派单时间").Message("不能为空")
		//valid.Required(store.CloseImacTime, "关店IMAC派单时间").Message("不能为空")
		//valid.Required(store.DmbUninstallTime, "DMB拆除派单时间").Message("不能为空")
		//valid.Required(store.DmbInstallTime, "DMB安装派单时间").Message("不能为空")
		//valid.Required(store.SamsungIntallTime, "DMB三星安装派单时间").Message("不能为空")
		//valid.Required(store.SamsungUnintallTime, "DMB三星拆除派单时间").Message("不能为空")
	}

	if store.Status == 3 {

		//判断是否有文件上传
		uploadStr := this.GetString("uploadList")
		var closeStore models.CloseStore
		if  uploadStr != ""{
			uploadList := FilterStrSlice(strings.Split(uploadStr,","))
			//遍历结构体赋值
			storeType  		:= reflect.TypeOf(closeStore)
			storeTypeVal  	:= reflect.New(storeType)
			//遍历上传文件
			for _,formFile  := range uploadList{
				if filePath,err := this.UpFileTable(formFile,1); err != nil {
					this.ReturnJson(-1,formFile+":"+err.Error(),nil)
				}else{
					for k := 0; k < storeType.NumField(); k++ {
						if   strings.ToLower(storeType.Field(k).Name) == strings.ToLower(formFile[5:]){
							storeTypeVal.Elem().Field(k).SetString(filePath)
						}
					}
				}
			}
			closeStore = storeTypeVal.Elem().Interface().(models.CloseStore)

			store.DeviceLcTogoTable = closeStore.DeviceLcTogoTable
			store.DeviceHpTogoTable = closeStore.DeviceHpTogoTable
			store.ToLcPropertyTable = closeStore.ToLcPropertyTable
			store.PropertyDestroyTable = closeStore.PropertyDestroyTable
			store.DeviceReturnLcApply = closeStore.DeviceReturnLcApply
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
	var storeStatus = []string{1:"新店",2:"IE",3:"关店",4:"临时关店",5:"外送",6:"CPOS",7:"转加盟"}
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
			var system	models.System
			this.Data["systemList"],_ =  system.GetList()
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

			this.Data["systemList"]    =  GetStoreSystem(storeId)
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

//标记店为特别关注
func (this *StoreController)SignStore() {
	var store models.PublicStore

	if err := this.ParseForm(&store); err != nil {
		this.ReturnJson(-1,err.Error(),nil)
	}

	if  this.GetString("storeId") == ""{
		this.ReturnJson(-1,"餐厅ID不能为空",nil)
	}
	//校验必填参数
	if  this.GetString("signFlag") == ""{
		this.ReturnJson(-2,"标识不能为空",nil)
	}

	err := store.SignStore()
	if err != nil{
		this.ReturnJson(-1,err.Error(),nil)
	}
	this.ReturnJson(0,"操作成功",nil)
}

//获取餐厅系统列表
func  GetStoreSystem(storeId string) []models.StoreSystemList {
	var  list []models.StoreSystemList
	sql :=  `select s.system_id,system_name,ss.store_system_id from system as s left join  (select store_system_id, system_id from store_system  where store_id = ? ) as ss on s.system_id = ss.system_id order by ss.store_system_id desc`
	orm.NewOrm().Raw(sql, storeId).QueryRows(&list)

	return list

}
