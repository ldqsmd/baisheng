package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ModelDevice struct {
	DeviceId 		int `json:"deviceId" form:"deviceId"  orm:"pk"`
	DeviceName		string `json:"deviceName" form:"deviceName"`
	Position		string `json:"position" form:"position"`
	Purpose			string `json:"purpose" form:"purpose"`
	DeviceModel		string `json:"deviceModel" form:"deviceModel"`
	ScrapModel		string `json:"scrapModel" form:"scrapModel"`
	DevicePic		string `json:"devicePic" form:"devicePic"`
	ModelId			int    `json:"modelId"  form:"modelId"`
	Remark			string `json:"remark"   form:"remark"`
	CreateTime 		string `json:"createTime"`
	UpdateTime 		string `json:"updateTime"`
}

func (this *ModelDevice) TableName() string {
	return "model_device"
}

func (this *ModelDevice) GetDeviceList(modelId string)([]ModelDevice,error){

	var deviceList []ModelDevice
	_, err := orm.NewOrm().Raw("select * from model_device where model_id=? order by create_time desc",modelId).QueryRows(&deviceList)
	if err == nil {
		return deviceList,err
	}
	return deviceList, nil
}


func (this *ModelDevice) GetDeviceDetail(deviceId string)(ModelDevice){

	var device  ModelDevice
	orm.NewOrm().Raw("select * from model_device where device_id=?",deviceId).QueryRow(&device)
	return device
}

func (this *ModelDevice)InsertOrUpdate()error  {

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


func (this *ModelDevice)Delete()error  {

	if  _,err :=  orm.NewOrm().Delete(this); err != nil{
		return err
	}
	return nil
}












