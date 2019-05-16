package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type StoreSystem struct {
	StoreSystemId 	int 	`json:"storeSystemId" form:"storeSystemId" orm:"pk"`
	SystemId 		int 	`json:"systemId" form:"systemId"`
	StoreId			int  	`json:"storeId"  form:"storeId"`
	CreateTime		string  `json:"createTime"`
}

//餐厅系统列表
type StoreSystemList struct {
	StoreSystem
	SystemName		string  `json:"systemName" 	form:"systemName"`
}


func (this *StoreSystem) TableName() string {
	return "store_system"
}

func (this *StoreSystem)InsertOrUpdate()error  {

	this.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	if _,err :=  orm.NewOrm().Insert(this); err != nil{
		return err
	}
	return nil
}

func (this *StoreSystem)Delete()error  {

	orm.NewOrm().Raw("SELECT store_system_id FROM store_system WHERE store_id=? and system_id=?",this.StoreId,this.SystemId).QueryRow(&this.StoreSystemId)

	if  _,err :=  orm.NewOrm().Delete(this); err != nil{
		return err
	}
	return nil
}















