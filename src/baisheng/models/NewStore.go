package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type NewStore struct {
	Id          	int			`form:"id"`
	StoreId		    int		`form:"storeId"`
	SmallNoticeTime	string		`form:"smallNoticeTime"`
	CreateTime		string		`form:"createTime"`
	Remark				string 	`form:"remark"`
	ItDebugTime			string  `form:"newStoreItDebugTime"`
	ImacDispatchTime	string	`form:"newStoreImacDispatchTime"`
	NewStoreInfo
}

type NewStoreInfo struct {

	ApplyEmailTime	string		`form:"applyEmailTime"`
	BookDeviceTime	string		`form:"bookDeviceTime"`
	Notice2gTime	string		`form:"notice2gTime"`
	Open2gImsTime	string		`form:"open2gimsTime"`
	Open2gCmsTime	string		`form:"open2gcmsTime"`
	ItsmRelationTime	string		`form:"itsmRelationTime"`
	ItspTime		string		`form:"itspTime"`
	DmbDispatchTime	string		`form:"dmbDispatchTime"`
	CallNumTime		string		`form:"callNumTime"`
	DeviceDebug		string		`form:"deviceDebug"`

}

func (this *NewStore) GetStoreList()([]NewStore,error){

	var newStoreList []NewStore
	_, err := orm.NewOrm().Raw("SELECT * from new_store where store_id=?",this.StoreId).QueryRows(&newStoreList)
	if err == nil {
		return newStoreList,err
	}
	return newStoreList, nil
}

//新增或者更新
func (this *NewStore)CreateOrUpdateNewStore(store Store)error  {

	o := orm.NewOrm()
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	//新增
	if  store.NewStoreId == 0 {
		_, err := o.Raw("INSERT INTO  new_store (store_id,apply_email_time,book_device_time,small_notice_time,notice2g_time,open2g_ims_time," +
			"open2g_cms_time,itsm_relation_time,itsp_time,imac_dispatch_time,dmb_dispatch_time,call_num_time,it_debug_time,device_debug,remark,create_time) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
			store.StoreId, store.ApplyEmailTime,store.BookDeviceTime,store.NewStoreSmallNoticeTime,store.Notice2gTime,store.Open2gImsTime,store.Open2gCmsTime,store.ItsmRelationTime, store.ItspTime,store.NewStoreImacDispatchTime,store.NewStoreDmbDispatchTime,store.CallNumTime,store.NewStoreItDebugTime,store.DeviceDebug,store.NewStoreRemark,nowTime).Exec()
		if err != nil {
			return  err
		}
	}else{
		//更新
		_, err := o.Raw("UPDATE new_store SET  " +
			"apply_email_time=?," +
			"book_device_time=?," +
			"small_notice_time=?," +
			"notice2g_time=?," +
			"open2g_ims_time=?," +
			"open2g_cms_time=?," +
			"itsm_relation_time=?," +
			"itsp_time=?," +
			"imac_dispatch_time=?," +
			"dmb_dispatch_time=?," +
			"call_num_time=?," +
			"it_debug_time=?," +
			"device_debug=?," +
			"remark=? where id=?" +
			"", store.ApplyEmailTime,store.BookDeviceTime,store.NewStoreSmallNoticeTime,store.Notice2gTime,store.Open2gImsTime,store.Open2gCmsTime,store.ItsmRelationTime, store.ItspTime,store.NewStoreImacDispatchTime,store.NewStoreDmbDispatchTime,store.CallNumTime,store.NewStoreItDebugTime,store.DeviceDebug,store.NewStoreRemark,store.NewStoreId).Exec()
		if err != nil {
			return  err
		}
	}
	return nil
}


//获取新店ID
func (this *NewStore)GetNewStoreInfo(storeId string)NewStore{
	//获取
	orm.NewOrm().Raw("SELECT * FROM new_store WHERE store_id=?", storeId).QueryRow(&this)

	return *this

}












