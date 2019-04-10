package models

import (
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
	IeStoreItDebugTime 			string  `form:"ieStoreItDebugTime"`
	NewStoreItDebugTime 			string  `form:"newStoreItDebugTime"`

	PublicStoreInfo
	NewStoreInfo
	IEStoreInfo
}


func (this *Store)AddStore()(int64,error) {

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
	return publicStore.AddStore()
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





