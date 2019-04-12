package models

import (
	"time"
)

type Store struct {

	StoreId          		int		`form:"storeId"`
	NewStoreId        		int		`form:"newStoreId"`
	CloseStoreId        	int		`form:"closeStoreId"`
	NewStoreSmallNoticeTime string  `form:"newStoreSmallNoticeTime"`
	StoreRemark 			string  `form:"storeRemark"`
	IeStoreRemark 			string  `form:"ieStoreRemark"`
	NewStoreRemark 			string  `form:"newStoreRemark"`
	CloseStoreRemark 		string  `form:"closeStoreRemark"`
	NewStoreItDebugTime 	string  `form:"newStoreItDebugTime"`
	NewStoreImacDispatchTime 	string  `form:"newStoreImacDispatchTime"`
	NewStoreDmbDispatchTime 	string  `form:"newStoreDmbDispatchTime"`
	CloseStoreDmbDispatchTime 	string  `form:"closeStoreDmbDispatchTime"`
	CloseStoreImacDispatchTime 	string  `form:"closeStoreImacDispatchTime"`
	PublicStoreInfo
	NewStoreInfo
	IeStore
	CloseStoreInfo
}

func (this *Store)InitPublicStore()PublicStore {

	var publicStore PublicStore
	publicStore.StoreName 		= this.StoreName
	publicStore.Number 			= this.Number
	publicStore.Brand 			= this.Brand
	publicStore.Status 			= this.Status
	publicStore.Manager 		= this.Manager
	publicStore.ManagerPhone 	= this.ManagerPhone
	publicStore.Address	 		= this.Address
	publicStore.AreaManager	  	= this.AreaManager
	publicStore.RegionalManager	= this.RegionalManager
	publicStore.Ip	  			= this.Ip
	publicStore.Network	  		= this.Network
	publicStore.OpenTime	 	= this.OpenTime
	publicStore.DecorationTime	= this.DecorationTime
	publicStore.CloseTime	  	= this.CloseTime
	publicStore.BookStartTime	= this.BookStartTime
	publicStore.EmployeeStartTime = this.EmployeeStartTime
	publicStore.WaitTime	  	= this.WaitTime
	publicStore.DeviceTime	  	= this.DeviceTime
	publicStore.BuildName	  	= this.BuildName
	publicStore.TempCloseTime	= this.TempCloseTime
	publicStore.ForbiddenStatus	= this.ForbiddenStatus
	publicStore.Remark			= this.StoreRemark

	return publicStore
}

//初始化 close store
func (this *Store)InitCloseStore()CloseStore{
	var closeStore CloseStore

	closeStore.StoreId 			= this.StoreId
	closeStore.Remark 			= this.CloseStoreRemark
	closeStore.CreateTime 		= time.Now().Format("2006-01-02 15:04:05")
	closeStore.DmbDispatchTime 	= this.CloseStoreDmbDispatchTime
	closeStore.ImacDispatchTime = this.CloseStoreImacDispatchTime
	closeStore.PropertyEvaluateTime = this.PropertyEvaluateTime
	closeStore.BusinessRecoveryTime = this.BusinessRecoveryTime
	closeStore.BusinessCloseTime = this.BusinessCloseTime
	closeStore.DmbUninstallNumber = this.DmbUninstallNumber
	closeStore.DispatchSamsungTime = this.DispatchSamsungTime
	closeStore.DeviceHpTogoTable = this.DeviceHpTogoTable
	closeStore.DeviceLcTogoTable = this.DeviceLcTogoTable
	closeStore.ToLcPropertyTable = this.ToLcPropertyTable
	closeStore.DeviceReturnLcApply = this.DeviceReturnLcApply
	closeStore.PropertyDestroyTable = this.PropertyDestroyTable

	return closeStore

}

//新增
func (this *Store)AddStore()(error) {

	publicStore := this.InitPublicStore()
	publicStore.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	storeId,err := publicStore.AddStore()
	if err != nil{
		return err
	}
	this.StoreId = int(storeId)

	//NEW store
	if  publicStore.Status  == 1 {
		var newStore NewStore
		 if err := newStore.CreateOrUpdateNewStore(*this);err != nil {
			return  err
		}
	}
	//IE store
	if  publicStore.Status  == 2 {
		var ieStore IeStore
		ieStore.CreateOrUpdateIeStore()
		ieStore.CreateTime 	= time.Now().Format("2006-01-02 15:04:05")
		_,err := ieStore.AddIEStore()
		if err != nil {
			return  err
		}
	}
	//clsose store
	if  publicStore.Status  == 3 {
		closeStore := this.InitCloseStore()
		_,err := closeStore.AddCloseStore()
		if err != nil {
			return  err
		}
	}
	return nil
}
//更新
func (this *Store)UpdateStore()error  {

	publicStore := this.InitPublicStore()
	publicStore.Id = this.StoreId
	publicStore.UpdateTime 	= time.Now().Format("2006-01-02 15:04:05")
	err :=  publicStore.UpdateStore()
	if err != nil {
		return err
	}
	//NEW store
	if  publicStore.Status  == 1 {
		var newStore NewStore
		err := newStore.CreateOrUpdateNewStore(*this)
		if err != nil {
			return  err
		}
	}
	//IE store
	if  publicStore.Status  == 2 {
		var ieStore IeStore
		err := ieStore.CreateOrUpdateIeStore(*this)
		if err != nil {
			return  err
		}
	}
	//clsose store
	if  publicStore.Status  == 3 {
		closeStore := this.InitCloseStore()
		closeStore.Id = this.CloseStoreId
		err := closeStore.UpdateCloseStore()
		if err != nil {
			return  err
		}
	}
	return nil
}





