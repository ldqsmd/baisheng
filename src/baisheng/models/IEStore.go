package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type IeStore struct {
	IeId					int			`form:"IeId"`
	ScrapAssessmentTime	 	string		`form:"scrapAssessmentTime"`
	DeviceBookFirsTime		string		`form:"deviceBookFirsTime"`
	IeSmallNoticeTime		string		`form:"ieSmallNoticeTime"`
	OpenImacTime		    string		`form:"openImacTime"`
	CloseImacTime		    string		`form:"closeImacTime"`
	ItsmCloseStoreTime		string		`form:"itsmCloseStoreTime"`
	DmbUninstallTime		string		`form:"dmbUninstallTime"`
	DmbInstallTime			string		`form:"dmbInstallTime"`
	DmbDebugTime			string	    `form:"dmbDebugTime"`
	SamsungUnintallTime		string	    `form:"samsungUnintallTime"`
	SamsungIntallTime		string	    `form:"samsungIntallTime"`
	DmbDebugTogo			string	    `form:"dmbDebugTogo"`
	IeItDebugTime			string	    `form:"ieItDebugTime"`
	IeRemark				string	    `form:"ieRemark"`
	IeStoreId		    	int			`form:"IeStoreId"`
	IeCreateTime			string	    `form:"ieCreateTime"`
}

func (this *IeStore) TableName() string {
	return "ie_store"
}

func (this *IeStore) GetIeList()([]IeStore,error){

	var IeStoreList []IeStore
	_, err := orm.NewOrm().Raw("SELECT * from ie_store where store_id=?",this.IeStoreId).QueryRows(&IeStoreList)
	if err == nil {
		return IeStoreList,err
	}
	return IeStoreList, nil
}

func (this *IeStore)AddIeStore()(int64,error) {
	this.IeCreateTime = time.Now().Format("2006-01-02 15:04:05")
	return orm.NewOrm().Insert(this)
}

func (this *IeStore)CreateOrUpdateIeStore(store Store)error  {


	//this.ScrapAssessmentTime= store.ScrapAssessmentTime
	//this.DeviceBookFirsTime = store.DeviceBookFirsTime
	//this.IeSmallNoticeTime 	= store.IeStoreSmallNoticeTime
	//this.CloseImacTime 		= store.CloseImacTime
	//this.OpenImacTime 		= store.OpenImacTime
	//this.ItsmCloseStoreTime = store.ItsmCloseStoreTime
	//this.DmbUninstallTime 	= store.DmbUninstallTime
	//this.DmbInstallTime 	= store.DmbInstallTime
	//this.DmbDebugTime 		= store.IeStoreItDebugTime
	//this.SamsungUnintallTime= store.SamsungUnintallTime
	//this.SamsungIntallTime 	= store.SamsungIntallTime
	//this.DmbDebugTogo 		= store.DmbDebugTogo
	//this.IeItDebugTime  	= store.IeStoreItDebugTime
	//this.Remark 			= store.IeStoreRemark
	//this.StoreId 			= store.StoreId
	//this.StoreId 			= store.StoreId


	o := orm.NewOrm()
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	if store.IeStoreId ==  0 {
		if _,err := orm.NewOrm().Insert(&store.IeStore); err != nil {
			return err
		}
	}else{
		_,err := o.Raw("UPDATE ie_store SET " +
			"scrap_assessment_time=?," +
			"device_book_firs_time=?," +
			"small_notice_time=?," +
			"close_imac_time=?," +
			"open_imac_time=?," +
			"itsm_close_store_time=?,dmb_uninstall_time=?,dmb_install_time=?,dmb_debug_time=?,samsung_intall_time=?,samsung_unintall_time=?,dmb_debug_togo=?,it_debug_time=?,remark=?,update_time=? WHERE id=?",
			store.ScrapAssessmentTime,
			store.DeviceBookFirsTime,
			store.IeSmallNoticeTime,
			store.CloseImacTime,
			store.OpenImacTime,
			store.ItsmCloseStoreTime,
			store.DmbUninstallTime,
			store.DmbInstallTime,
			store.DmbDebugTime,
			store.SamsungIntallTime,
			store.SamsungUnintallTime,
			store.DmbDebugTogo,
			store.IeItDebugTime,
			store.IeStoreRemark,
			nowTime,
			store.IeStoreId).Exec()

	if err != nil {
		return err
	}
	}

	return nil
}

//获取新店ID
func (this *IeStore)GetIeStoreInfo(storeId string)IeStore{
	//获取
	orm.NewOrm().Raw("SELECT * FROM ie_store WHERE store_id=?", storeId).QueryRow(&this)

	return *this
}












