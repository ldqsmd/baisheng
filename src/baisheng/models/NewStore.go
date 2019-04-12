package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type NewStore struct {

	NewId          			int			`form:"newId" orm:"pk"`
	NewStoreId		    	int			`form:"newStoreId"`
	NewSmallNoticeTime		string		`form:"newSmallNoticeTime"`
	NewRemark				string 		`form:"newRemark"`
	NewItDebugTime			string  	`form:"newItDebugTime"`
	NewDmbDispatchTime		string  	`form:"newDmbDispatchTime"`
	NewImacDispatchTime		string		`form:"newImacDispatchTime"`
	NewCreateTime			string		`form:"_"`
	NewUpdateTime			string		`form:"_"`
	ApplyEmailTime			string		`form:"applyEmailTime"`
	BookDeviceTime			string		`form:"bookDeviceTime"`
	Notice2gTime			string		`form:"notice2gTime"`
	Open2gImsTime			string		`form:"open2gimsTime"`
	Open2gCmsTime			string		`form:"open2gcmsTime"`
	ItsmRelationTime		string		`form:"itsmRelationTime"`
	ItspTime				string		`form:"itspTime"`
	CallNumTime				string		`form:"callNumTime"`
	DeviceDebug				string		`form:"deviceDebug"`

}

func (this *NewStore) GetStoreList()([]NewStore,error){

	var newStoreList []NewStore
	_, err := orm.NewOrm().Raw("SELECT * from new_store where store_id=?",this.NewStoreId).QueryRows(&newStoreList)
	if err == nil {
		return newStoreList,err
	}
	return newStoreList, nil
}

//新增或者更新
func (this *NewStore)InsertOrUpdate()error  {

	if this.NewId == 0 {
		this.NewCreateTime  = time.Now().Format("2006-01-02 15:04:05")
		newId,err :=  orm.NewOrm().Insert(this)
		if err != nil{
			return err
		}
		this.NewId = int(newId)
	}else {
		this.NewUpdateTime  = time.Now().Format("2006-01-02 15:04:05")
		_,err :=  orm.NewOrm().Update(this)
		if err != nil{
			return err
		}
	}
	return nil
}


//获取新店ID
func (this *NewStore)GetNewStoreInfo(storeId string)NewStore{
	//获取
	orm.NewOrm().Raw("SELECT * FROM new_store WHERE new_store_id=?", storeId).QueryRow(&this)
	return *this

}












