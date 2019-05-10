package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Device struct {
	DeviceId 		int 	`json:"deviceId"    form:"deviceId"  orm:"pk"`
	DeviceName 		string 	`json:"deviceName"  form:"deviceName"  `
	DeviceJde		string 	`json:"deviceJde"   form:"deviceJde"  `
	DeviceType		string 	`json:"deviceType"  form:"deviceType"  `
	Utilize			string 	`json:"utilize" form:"utilize"  `
	Supplier		string 	`json:"supplier" form:"supplier"  `
	SupplierJde		string 	`json:"supplierJde" form:"supplierJde"  `
	Origin			string 	`json:"origin" form:"origin"`
	Remark			string  `json:"remark" form:"remark"`
	Description		string  `json:"description" form:"description"`
	CreateTime		string  `json:"createTime"`
	UpdateTime		string  `json:"updateTime"`

}

func (this *Device) TableName() string {
	return "device"
}

func (this *Device) GetDeviceList()([]Device,error){

	var deviceList []Device
	_, err := orm.NewOrm().Raw("select * from device order by create_time desc").QueryRows(&deviceList)
	if err == nil {
		return deviceList,err
	}
	return deviceList, nil
}


func (this *Device) GetDeviceInfo(deviceId string)(Device){
	var device  Device
	orm.NewOrm().Raw("select * from device where device_id=?",deviceId).QueryRow(&device)
	return device
}

func (this *Device)InsertOrUpdate()error  {

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	if this.DeviceId == 0 {
		this.CreateTime = nowTime
		this.UpdateTime = nowTime
		if id,err :=  orm.NewOrm().Insert(this); err != nil{
			return err
		}else{
			this.DeviceId  = int(id)
		}
	}else {
		this.UpdateTime = nowTime
		if  _,err :=  orm.NewOrm().Update(this); err != nil{
			return err
		}
	}
	return nil
}


func (this *Device)DeleteDevice()error  {

	if  _,err :=  orm.NewOrm().Delete(this); err != nil{
		return err
	}
	return nil
}












