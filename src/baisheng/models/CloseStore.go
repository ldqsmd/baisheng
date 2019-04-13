package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type CloseStore struct {
	CloseId          		int			`form:"closeId" orm:"pk"`
	CloseStoreId		    int			`form:"closeStoreId"`
	CloseCreateTime			string		`form:"_"`
	CloseUpdateTime			string		`form:"_"`
	CloseRemark				string		`form:"closeRemark"`
	CloseDmbDispatchTime	string		`form:"closeDmbDispatchTime"`
	CloseImacDispatchTime string		`form:"closeImacDispatchTime"`


	PropertyEvaluateTime	string		`form:"propertyEvaluateTime"`
	BusinessRecoveryTime	string		`form:"businessRecoveryTime"`
	BusinessCloseTime		string		`form:"businessCloseTime"`
	DmbUninstallNumber		int			`form:"dmbUninstallNumber"`
	DispatchSamsungTime		string		`form:"dispatchSamsungTime"`
	DeviceHpTogoTable		string		`form:"deviceHpTogoTable"`
	DeviceLcTogoTable		string		`form:"deviceLcTogoTable"`
	ToLcPropertyTable		string		`form:"toLcPropertyTable"`
	DeviceReturnLcApply		string		`form:"deviceReturnLcApply"`
	PropertyDestroyTable	string		`form:"propertyDestroyTable"`


}

func (this *CloseStore) GetCloseStoreList()([]CloseStore,error){

	var closeStoreList []CloseStore
	_, err := orm.NewOrm().Raw("SELECT * from close_store where store_id=?",this.CloseStoreId).QueryRows(&closeStoreList)
	if err == nil {
		return closeStoreList,err
	}
	return closeStoreList, nil
}

func (this *CloseStore)InsertOrUpdate()error  {

	if this.CloseId == 0 {
		this.CloseCreateTime  = time.Now().Format("2006-01-02 15:04:05")
		closeId,err :=  orm.NewOrm().Insert(this)
		if err != nil{
			return err
		}
		this.CloseId = int(closeId)
	}else {
		this.CloseUpdateTime  = time.Now().Format("2006-01-02 15:04:05")
		if  _,err :=  orm.NewOrm().Update(this); err != nil{
			return err
		}
	}

	return nil
}

//获取新店ID
func (this *CloseStore)GetCloseStoreInfo(storeId string)CloseStore{
	//获取
	orm.NewOrm().Raw("SELECT * FROM close_store WHERE close_store_id=?", storeId).QueryRow(&this)
	return *this
}












