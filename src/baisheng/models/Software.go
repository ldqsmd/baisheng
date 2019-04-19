package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Software struct {	
	SoftwareId 		int 	`json:"softwareId" orm:"pk"`
	SoftwareName 	string 	`json:"softwareName" form:"softwareName"`
	Version 		string 	`json:"version" form:"version"`
	Status  		int 	`json:"softwareStatus"`
	CreateTime		string  `json:"createTime"`
	ConfirmId		int 	`json:"confirmId" form:"confirmId"`
	AdminId			int 	`json:"adminId"`
	CheckTime		string `json:"checkTime"`
	Remark			string `json:"remark" form:"remark"`
}


func (this *Software) GetSoftwareList(confirmId string)([]Software,error){

	var softwareList []Software
	_, err := orm.NewOrm().Raw("SELECT * FROM  software where confirm_id="+confirmId).QueryRows(&softwareList)
	if err == nil {
		return softwareList,err
	}
	return softwareList, nil
}


func (this *Software)InsertOrUpdate()error  {

	if this.SoftwareId == 0 {

		softwareId,err :=  orm.NewOrm().Insert(this)
		if err != nil{
			return err
		}
		this.SoftwareId = int(softwareId)
	}else {
		if  _,err :=  orm.NewOrm().Update(this); err != nil{
			return err
		}
	}
	return nil
}


func (this *Software)SignSoftware()error  {

	this.Status = 1
	this.CheckTime  = time.Now().Format("2006-01-02 15:04:05")
	if  _,err :=  orm.NewOrm().Update(this,"Status"); err != nil{
		return err
	}
	return nil
}












