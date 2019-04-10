package models

import (
	"github.com/astaxie/beego/orm"
)

type StoreStatus struct {
	Id          	int		`form:"id"`
	Status		    int		`form:"status"`
	Content		    string		`form:"content"`
}

func (this *StoreStatus) GetStatusList()([]StoreStatus,error){

	var statusList []StoreStatus
	_, err := orm.NewOrm().Raw("SELECT * from store_status").QueryRows(&statusList)
	if err == nil {
		return statusList,err
	}
	return statusList, nil
}












