package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type IeStore struct {
	IeId					int			`form:"ieId" orm:"pk"`
	ScrapAssessmentTime	 	string		`form:"scrapAssessmentTime"`
	DeviceBookFirsTime		string		`form:"deviceBookFirsTime"`
	IeSmallNoticeTime		string		`form:"ieSmallNoticeTime"`
	OpenImacTime		    string		`form:"openImacTime"`
	CloseImacTime		    string		`form:"closeImacTime"`
	ItsmCloseStoreTime		string		`form:"itsmCloseStoreTime"`
	DmbUninstallTime		string		`form:"dmbUninstallTime"`
	DmbInstallTime			string		`form:"dmbInstallTime"`
	SamsungUnintallTime		string	    `form:"samsungUnintallTime"`
	SamsungIntallTime		string	    `form:"samsungIntallTime"`
	DmbDebugTogo			string	    `form:"dmbDebugTogo"`
	IeItDebugTime			string	    `form:"ieItDebugTime"`
	IeRemark				string	    `form:"ieRemark"`
	IeStoreId		    	int			`form:"IeStoreId"`
	IeCreateTime			string	    `form:"_"`
	IeUpdateTime			string	    `form:"_"`
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

func (this *IeStore)InsertOrUpdate()error  {

	if this.IeId == 0 {
		this.IeCreateTime  = time.Now().Format("2006-01-02 15:04:05")
		ieId,err :=  orm.NewOrm().Insert(this)
		if err != nil{
			return err
		}
		this.IeId = int(ieId)
	}else {
		this.IeUpdateTime  = time.Now().Format("2006-01-02 15:04:05")
		if  _,err :=  orm.NewOrm().Update(this); err != nil{
			return err
		}
	}

	return nil
}

//获取新店ID
func (this *IeStore)GetIeStoreInfo(storeId string)IeStore{
	//获取
	orm.NewOrm().Raw("SELECT * FROM ie_store WHERE ie_store_id=?", storeId).QueryRow(&this)

	return *this
}












