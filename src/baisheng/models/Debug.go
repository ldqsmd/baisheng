package models

import (
	"github.com/astaxie/beego/orm"
)

type Confirm struct {
	ConfirmId   	int			`form:"confirmId" orm:"pk"`
	AdminId   		int			`form:"adminId" `
	StoreId   		int			`form:"storeId" `

}
type ConfirmList struct {
	Confirm
	Number     		string		`form:"number" `
	UserName   		string		`form:"userName"`
	CompleteTime   	string		`form:"completeTime" `
	ConfirmStatus  	int			`form:"confirmStatus" ` // 0 进行中1 已完成
	ConfirmRemark  	string		`form:"confirmRemark" `
	StoreStatus  	int			`form:"storeStatus" `

}

func (this *Confirm) GetConfirmList()([]ConfirmList,error){

	var confirmList []ConfirmList
	_, err := orm.NewOrm().Raw("SELECT store.status as store_status, confirm.id as confirm_id,store.number,admin.user_name,confirm.complete_time,confirm.confirm_status,confirm.remark as confirm_remark  FROM  confirm   inner join  store   on  confirm.store_id = store.store_id inner join admin on confirm.admin_id = admin.id").QueryRows(&confirmList)
	if err == nil {
		return confirmList,err
	}
	return confirmList, nil
}


func (this *Confirm)InsertOrUpdate()error  {

	if this.ConfirmId == 0 {

		conformId,err :=  orm.NewOrm().Insert(this)
		if err != nil{
			return err
		}
		this.ConfirmId = int(conformId)
	}else {
		if  _,err :=  orm.NewOrm().Update(this); err != nil{
			return err
		}
	}

	return nil
}












