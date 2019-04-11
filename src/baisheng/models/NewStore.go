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

func (this *NewStore)AddNewStore()(int64,error) {
	this.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return orm.NewOrm().Insert(this)
}

func (this *NewStore)UpdateStore()error  {
	_,err :=  orm.NewOrm().Update(this,
		"ApplyEmailTime",
		"BookDeviceTime",
		"Notice2gTime",
		"Open2gImsTime",
		"Open2gCmsTime",
		"ItsmRelationTime",
		"ItspTime",
		"DmbDispatchTime",
		"CallNumTime",
		"DeviceDebug",
		"SmallNoticeTime",
		"Remark",
		"ItDebugTime",
		"ImacDispatchTime",
		)
	return err
}


//获取新店ID
func (this *NewStore)GetNewStoreInfo(storeId string)NewStore{
	//获取
	orm.NewOrm().Raw("SELECT * FROM new_store WHERE store_id=?", storeId).QueryRow(&this)

	return *this

}











