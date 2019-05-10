package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type SystemDevice struct {
	Id 				int 	`json:"id" form:"id" orm:"pk"`
	SystemId 		int 	`json:"systemId" form:"systemId" `
	DeviceId 		int 	`json:"deviceId" form:"deviceId" `
	CreateTime		string  `json:"createTime"`
	UpdateTime		string  `json:"updateTime"`
}

type SystemDeviceList struct {
	SystemDevice
	DeviceId   		int	   `json:"deviceId"`
	DeviceName 		string `json:"deviceName"`
}

func (this *SystemDevice) TableName() string {
	return "system_device"
}

func (this *SystemDevice) GetSystemDeviceList(system int)([]SystemDeviceList,error){

	var list []SystemDeviceList
	_, err := orm.NewOrm().Raw("select sd*,d.device_id,d.device_name from system_device as sd left join device as d on  sd.device_id = d.device_id   order  by sd.create_time desc").QueryRows(&list)
	if err == nil {
		return list,err
	}
	return list, nil
}


func (this *SystemDevice)Insert()error  {

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	this.CreateTime = nowTime
	this.UpdateTime = nowTime
	if id,err :=  orm.NewOrm().Insert(this); err != nil{
		return err
	}else{
		this.SystemId  = int(id)
	}
	return nil
}












