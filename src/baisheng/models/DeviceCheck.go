package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


type DeviceCheck struct {
	CheckId 		int 	`json:"checkId" form:"checkId"  orm:"pk"`
	StoreId			int 	`json:"storeId" form:"storeId"`
	Status 			int 	`json:"status" form:"status"`
	ReplyDate		string 	`json:"replyDate" form:"replyDate"`
	UpdateTime		string 	`json:"updateTime"`
	CreateTime		string 	`json:"createTime"`
	ReadyFile		string 	`json:"readyFile" form:"readyFile"`
	CompleteFile	string 	`json:"completeFile" form:"completeFile"`
	AdminId			int 	`json:"adminId"`
	Remark			string  `json:"remark" form:"remark"`
}

type CheckList struct {
	AdminName		string  `json:"adminName" `
	StoreName		string  `json:"storeName" `
	Number			string  `json:"number" `
	DeviceCheck
}

func (this *DeviceCheck) TableName() string {
	return "device_check"
}

func (this *CheckList) GetCheckList()([]CheckList,error){

	var checkList []CheckList
	_, err := orm.NewOrm().Raw("SELECT d.*,store.store_name,store.number,a.user_name as admin_name FROM  device_check as d left join admin as a on d.admin_id=a.id left join store on store.store_id = d.store_id ").QueryRows(&checkList)
	if err == nil {
		return checkList,err
	}
	return checkList, nil
}


func (this *DeviceCheck) GetCheckInfo(checkId string)(DeviceCheck){

	var deviceCheck  DeviceCheck
	orm.NewOrm().Raw("select * from device_check where check_id=?",checkId).QueryRow(&deviceCheck)
	return deviceCheck
}

func (this *DeviceCheck)InsertOrUpdate()error  {

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	if this.CheckId == 0 {
		this.CreateTime = nowTime
		this.UpdateTime = nowTime
		checkId,err :=  orm.NewOrm().Insert(this)
		if err != nil{
			return err
		}
		this.CheckId = int(checkId)
	}else {
		this.UpdateTime = nowTime
		if  _,err :=  orm.NewOrm().Update(this,"StoreId","Status","ReplyDate","UpdateTime","ReadyFile","CompleteFile","AdminId","Remark"); err != nil{
			return err
		}
	}
	return nil
}


func (this *DeviceCheck)DeleteDeviceCheck()error  {

	if  _,err :=  orm.NewOrm().Delete(this); err != nil{
		return err
	}
	return nil
}












