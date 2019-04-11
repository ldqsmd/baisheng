package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type IEStore  struct {
	Id          			int		`form:"id"`
	StoreId		    		string		`form:"storeId"`
	SmallNoticeTime		    string		`form:"smallNoticeTime"`
	CreateTime				string	    `form:"createTime"`
	Remark					string	    `form:"remark"`
	ItDebugTime				string	    `form:"itDebugTime"`
	IEStoreInfo

}

type IEStoreInfo struct {

	ScrapAssessmentTime	 	string		`form:"scrapAssessmentTime"`
	DeviceBookFirsTime		string		`form:"deviceBookFirsTime"`
	OpenImacTime		    string		`form:"openImacTime"`
	CloseImacTime		    string		`form:"closeImacTime"`
	ItsmCloseStoreTime		string		`form:"itsmCloseStoreTime"`
	DmbUninstallTime		string		`form:"dmbUninstallTime"`
	DmbInstallTime			string		`form:"dmbInstallTime"`
	DmbDebugTime			string	    `form:"dmbDebugTime"`
	SamsungUnintallTime		string	    `form:"samsungUnintallTime"`
	SamsungIntallTime		string	    `form:"samsungIntallTime"`
	DmbDebugTogo			string	    `form:"dmbDebugTogo"`

}

func (this *IEStore) TableName() string {
	return "ie_store"
}

func (this *IEStore) GetIeList()([]IEStore,error){

	var ieStoreList []IEStore
	_, err := orm.NewOrm().Raw("SELECT * from ie_store where store_id=?",this.StoreId).QueryRows(&ieStoreList)
	if err == nil {
		return ieStoreList,err
	}
	return ieStoreList, nil
}

func (this *IEStore)AddIEStore()(int64,error) {
	this.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return orm.NewOrm().Insert(this)
}

func (this *IEStore)UpdateIeStore()error  {
	_,err :=  orm.NewOrm().Update(this)
	return err
}


//获取新店ID
func (this *NewStore)GetIeStoreInfo(){
	//获取
	orm.NewOrm().Raw("SELECT * FROM ie_store WHERE store_id=?", this.StoreId).QueryRow(&this)

}












