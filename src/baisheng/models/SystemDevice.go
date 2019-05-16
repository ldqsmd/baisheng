package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type SystemDevice struct {
	SystemDeviceId 	 int 	`json:"systemDeviceId" form:"systemDeviceId" orm:"pk"`
	SystemId 		int 	`json:"systemId" form:"systemId" `
	DeviceId 		int 	`json:"deviceId" form:"deviceId" `
	Active 			int 	`json:"active" form:"active" `
	CreateTime		string  `json:"createTime"`
	UpdateTime		string  `json:"updateTime"`
}

type SystemDeviceList struct {
	SystemId		 int 	`json:"systemId" `
	SystemDeviceId 	 int 	`json:"systemDeviceId" `
	Active 	 		 int 	`json:"active"`
	Device
}

func (this *SystemDevice) TableName() string {
	return "system_device"
}

func (this *SystemDevice) GetSystemDeviceList(systemId int)([]SystemDeviceList,error){

	var list []SystemDeviceList

	_,err := orm.NewOrm().Raw("SELECT  d.*,sd.system_device_id, sd.system_id, sd.active FROM device AS d  left join (select * from system_device where system_id=?) AS sd ON sd.device_id = d.device_id ORDER BY sd.active DESC , sd.system_device_id DESC",systemId).QueryRows(&list)

	return list, err
}


func (this *SystemDevice)InsertOrUpdate()error  {

	var systemDeviceId int

	orm.NewOrm().Raw("select system_device_id  from system_device where system_id=? and device_id=?",this.SystemId,this.DeviceId).QueryRow(&systemDeviceId)
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	this.SystemDeviceId = systemDeviceId

	if this.SystemDeviceId  == 0 {
		this.CreateTime = nowTime
		this.UpdateTime = nowTime
		if id,err :=  orm.NewOrm().Insert(this); err != nil{
			return err
		}else{
			this.SystemDeviceId  = int(id)
		}
	}else {
		this.SystemDeviceId = systemDeviceId
		this.UpdateTime = nowTime
		if  _,err :=  orm.NewOrm().Update(this); err != nil{
			return err
		}
	}
	return nil
}










