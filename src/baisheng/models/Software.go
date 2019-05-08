package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


type Software struct {
	SoftwareId 		int 	`form:"softwareId" orm:"pk"`
	SoftwareName 	string 	`json:"softwareName" form:"softwareName"`
	Version 		string 	`json:"version" form:"version"`
	Status  		int 	`json:"softwareStatus"`
	CreateTime		string  `json:"createTime"`
	ConfirmId		int 	`json:"confirmId" form:"confirmId"`
	AdminId			int 	`json:"adminId"`
	CheckTime		string  `json:"checkTime"`
	Remark			string  `json:"remark" form:"remark"`
}

type SoftwareList struct {
	AdminName		string  `json:"adminName" `
	Software
}

func (this *Software) GetSoftwareList(confirmId string)([]SoftwareList,error){

	var softwareList []SoftwareList
	_, err := orm.NewOrm().Raw("SELECT s.*,a.user_name as admin_name FROM  software as s left join admin as a on s.admin_id=a.id where confirm_id="+confirmId).QueryRows(&softwareList)
	if err == nil {
		return softwareList,err
	}
	return softwareList, nil
}



func (this *Software) GetSoftwareInfo(softwareId string)(Software,error){

	var software  Software
	orm.NewOrm().Raw("select * from software where software_id=?",softwareId).QueryRow(&software)
	return software, nil
}


func (this *Software)InsertOrUpdate()error  {

	if this.SoftwareId == 0 {
		softwareId,err :=  orm.NewOrm().Insert(this/*,"SoftwareName","AdminId","Version","Remark"*/)
		if err != nil{
			return err
		}
		this.SoftwareId = int(softwareId)
	}else {
		this.Status  = 0
		if  _,err :=  orm.NewOrm().Update(this,"Status","SoftwareName","Version","Remark"); err != nil{
			return err
		}
	}
	var confirm Confirm
	confirm.ConfirmId = this.ConfirmId
	confirm.ConfirmStatus = 0
	orm.NewOrm().Update(&confirm,"ConfirmStatus")

	return nil
}

//标记完成
func (this *Software)SignSoftware()error  {

	this.Status  = 1
	this.CheckTime  = time.Now().Format("2006-01-02 15:04:05")
	if  _,err :=  orm.NewOrm().Update(this,"Status","CheckTime","AdminId"); err != nil{
		return err
	}

	return nil
}


func (this *Software)DeleteSofeware()error  {

	if  _,err :=  orm.NewOrm().Delete(this); err != nil{
		return err
	}
	return nil
}












