package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Confirm struct {
	ConfirmId   	int			`form:"confirmId" orm:"pk"`
	AdminId   		int			`form:"adminId" `
	StoreId   		int			`form:"storeId" `
	Remark   		string		`form:"remark" `
	ConfirmStatus  	int			`form:"confirmStatus" ` // 0 进行中1 已完成
	CompleteTime   	string		`form:"completeTime" `
	CreateTime		string		`form:"_" `

}
type ConfirmList struct {
	Confirm
	Number     		string		`form:"number" ` //餐厅编号
	UserName   		string		`form:"userName"`
	ConfirmRemark  	string		`form:"confirmRemark" `
	StoreStatus  	int			`form:"storeStatus" `

}

func (this *Confirm) GetConfirmList()([]ConfirmList,error){

	var confirmList []ConfirmList
	_, err := orm.NewOrm().Raw("SELECT store.status as store_status, confirm_id,store.number,admin.user_name,confirm.complete_time,confirm.confirm_status,confirm.remark as confirm_remark  FROM  confirm   inner join  store   on  confirm.store_id = store.store_id inner join admin on confirm.admin_id = admin.id order by confirm.create_time desc").QueryRows(&confirmList)
	if err == nil {
		return confirmList,err
	}
	return confirmList, nil
}

func (this *Confirm) GetConfirmInfo(confirmId string)(Confirm,error){

	var confirm  Confirm
	orm.NewOrm().Raw("select * from confirm where confirm_id=?",confirmId).QueryRow(&confirm)
	return confirm, nil
}


func (this *Confirm)InsertOrUpdate()error  {

	if this.ConfirmId == 0 {
		this.CreateTime  = time.Now().Format("2006-01-02 15:04:05")
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

func (this *Confirm)SignCofirm()error  {

	_,err :=  orm.NewOrm().Update(this,"ConfirmFlag")
	return err
}












