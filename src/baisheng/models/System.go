package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type System struct {
	SystemId 		int 	`json:"systemId" form:"systemId" orm:"pk"`
	SystemName		string  `json:"systemName" form:"systemName"`
	Remark			string	`json:"Remark" form:"Remark"`
	Status			string	`json:"status" form:"status"`
	CreateTime		string  `json:"createTime"`
	UpdateTime		string  `json:"updateTime"`

}

func (this *System) TableName() string {
	return "system"
}

func (this *System) GetList()([]System,error){

	var systemList []System
	_, err := orm.NewOrm().Raw("select * from system order  by create_time desc").QueryRows(&systemList)
	if err == nil {
		return systemList,err
	}
	return systemList, nil
}


func (this *System) GetInfo(systemId string)(System){
	var system  System
	orm.NewOrm().Raw("select * from system where system_id=?",systemId).QueryRow(&system)
	return system
}


func (this *System)InsertOrUpdate()error  {

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	if this.SystemId == 0 {
		this.CreateTime = nowTime
		this.UpdateTime = nowTime
		if id,err :=  orm.NewOrm().Insert(this); err != nil{
			return err
		}else{
			this.SystemId  = int(id)
		}
	}else {
		this.UpdateTime = nowTime
		if  _,err :=  orm.NewOrm().Update(this); err != nil{
			return err
		}
	}
	return nil
}

func (this *System)ChangeStatus()error  {

	this.UpdateTime =  time.Now().Format("2006-01-02 15:04:05")
	if  _,err :=  orm.NewOrm().Update(this,"UpdateTime","Status"); err != nil{
		return err
	}
	return nil
}












