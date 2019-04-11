package models

import (
	"errors"
	"strconv"
	"time"
)

type Store struct {

	StoreId          		int		`form:"storeId"`
	NewStoreId        		int		`form:"newStoreId"`
	IeStoreId        		int		`form:"ieStoreId"`
	IeStoreSmallNoticeTime 	string  `form:"ieStoreSmallNoticeTime"`
	NewStoreSmallNoticeTime string  `form:"newStoreSmallNoticeTime"`
	StoreRemark 			string  `form:"storeRemark"`
	IeStoreRemark 			string  `form:"ieStoreRemark"`
	NewStoreRemark 			string  `form:"newStoreRemark"`
	CloseStoreRemark 		string  `form:"closeStoreRemark"`
	IeStoreItDebugTime 		string  `form:"ieStoreItDebugTime"`
	NewStoreItDebugTime 	string  `form:"newStoreItDebugTime"`
	NewStoreImacDispatchTime 	string  `form:"newStoreImacDispatchTime"`
	NewStoreDmbDispatchTime 	string  `form:"newStoreDmbDispatchTime"`
	CloseStoreDmbDispatchTime 	string  `form:"closeStoreDmbDispatchTime"`
	CloseStoreImacDispatchTime 	string  `form:"closeStoreImacDispatchTime"`


	PublicStoreInfo
	NewStoreInfo
	IEStoreInfo
	CloseStoreInfo
}



func (this *Store)AddStore()(error) {

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
	publicStore.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	storeId,err := publicStore.AddStore()

	if err != nil{
		return err
	}
	//NEW store
	if  publicStore.Status  == 1 {
		var  newStore NewStore
		newStore.StoreId 			= strconv.Itoa(int(storeId))
		newStore.SmallNoticeTime	= this.NewStoreSmallNoticeTime
		newStore.Remark 			= this.NewStoreRemark
		newStore.ItDebugTime 		= this.NewStoreItDebugTime
		newStore.ApplyEmailTime 	= this.ApplyEmailTime
		newStore.BookDeviceTime 	= this.BookDeviceTime
		newStore.Notice2gTime 		= this.Notice2gTime
		newStore.Open2gImsTime 		= this.Open2gImsTime
		newStore.Open2gCmsTime 		= this.Open2gCmsTime
		newStore.ItsmRelationTime 	= this.ItsmRelationTime
		newStore.ItspTime 			= this.ItspTime
		newStore.ImacDispatchTime 	= this.NewStoreImacDispatchTime
		newStore.DmbDispatchTime 	= this.NewStoreDmbDispatchTime
		newStore.CallNumTime 		= this.CallNumTime
		newStore.DeviceDebug 		= this.DeviceDebug
		newStore.CreateTime 		= time.Now().Format("2006-01-02 15:04:05")

		newStoreId,err := newStore.AddNewStore()
		if err != nil {
			return  err
		}
		return errors.New(string(newStoreId))
	}
	//IE store
	if  publicStore.Status  == 2 {
		var ieStore IEStore

		ieStore.StoreId 			= strconv.Itoa(int(storeId))
		ieStore.SmallNoticeTime 	= this.IeStoreSmallNoticeTime
		ieStore.Remark 				= this.IeStoreRemark
		ieStore.ItDebugTime  		= this.IeStoreItDebugTime
		ieStore.ScrapAssessmentTime = this.ScrapAssessmentTime
		ieStore.DeviceBookFirsTime 	= this.DeviceBookFirsTime
		ieStore.OpenImacTime 		= this.OpenImacTime
		ieStore.CloseImacTime 		= this.CloseImacTime
		ieStore.ItsmCloseStoreTime 	= this.ItsmCloseStoreTime
		ieStore.DmbUninstallTime 	= this.DmbUninstallTime
		ieStore.DmbInstallTime 		= this.DmbInstallTime
		ieStore.DmbDebugTime 		= this.IeStoreItDebugTime
		ieStore.SamsungUnintallTime = this.SamsungUnintallTime
		ieStore.SamsungIntallTime 	= this.SamsungIntallTime
		ieStore.DmbDebugTogo 		= this.DmbDebugTogo
		ieStore.CreateTime 			= time.Now().Format("2006-01-02 15:04:05")

		_,err := ieStore.AddIEStore()
		if err != nil {
			return  err
		}
	}
	//clsose store
	if  publicStore.Status  == 3 {
		var closeStore CloseStore

		closeStore.StoreId 			= strconv.Itoa(int(storeId))
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

		_,err := closeStore.AddCloseStore()
		if err != nil {
			return  err
		}
	}


	return nil

}


func (this *Store)UpdateStore()error  {

	var publicStore PublicStore

	publicStore.Id 				= this.StoreId
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

	return publicStore.UpdateStore()
}





