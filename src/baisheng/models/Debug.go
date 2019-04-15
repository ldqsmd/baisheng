package models

import (
	"github.com/astaxie/beego/orm"
)

type Debug struct {

	ConfirmId   	int			`form:"confirmId" `
	Number     		string		`form:"number" `
	UserName   		string		`form:"userName" `
	CompleteTime   	string		`form:"completeTime" `
	ConfirmStatus  	int		`form:"confirmStatus" ` // 0 进行中1 已完成
	ConfirmRemark  	string		`form:"confirmRemark" `
	StoreStatus  	int		`form:"storeStatus" `

}

func (this *Debug) GetDebugList()([]Debug,error){

	var debugList []Debug
	_, err := orm.NewOrm().Raw("SELECT store.status as store_status, confirm.id as confirm_id,store.number,admin.user_name,confirm.complete_time,confirm.confirm_status,confirm.remark as confirm_remark  FROM  confirm   inner join  store   on  confirm.store_id = store.store_id inner join admin on confirm.admin_id = admin.id").QueryRows(&debugList)
	if err == nil {
		return debugList,err
	}
	return debugList, nil
}
















