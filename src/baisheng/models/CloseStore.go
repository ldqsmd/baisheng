package models

import (
	"github.com/astaxie/beego/orm"
)

type CloseStore struct {
	CloseId          		int			`form:"closeId" orm:"pk"`
	CloseStoreId		    int			`form:"closeStoreId"`
	CloseCreateTime			string		`form:"closeCreateTime"`
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

func (this *CloseStore) GetNewStoreList()([]CloseStore,error){

	var closeStoreList []CloseStore
	_, err := orm.NewOrm().Raw("SELECT * from close_store where store_id=?",this.CloseStoreId).QueryRows(&closeStoreList)
	if err == nil {
		return closeStoreList,err
	}
	return closeStoreList, nil
}

func (this *CloseStore)InsertOrUpdate()(error) {
	return nil
}

func (this *CloseStore)UpdateCloseStore()error  {
	_,err :=  orm.NewOrm().Update(this,
		"Remark",
		"DmbDispatchTime",
		"ImacDispatchTime",
		"PropertyEvaluateTime",
		"BusinessRecoveryTime",
		"BusinessCloseTime",
		"DmbUninstallNumber",
		"DispatchSamsungTime",
		"DeviceHpTogoTable",
		"DeviceLcTogoTable",
		"ToLcPropertyTable",
		"DeviceReturnLcApply",
		"PropertyDestroyTable",
		)
	return err
}

//获取新店ID
func (this *CloseStore)GetCloseStoreInfo(storeId string)CloseStore{
	//获取
	orm.NewOrm().Raw("SELECT * FROM close_store WHERE store_id=?", storeId).QueryRow(&this)
	return *this
}












