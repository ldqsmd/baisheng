package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Confirm struct {
	ConfirmId   	int			`form:"confirmId" orm:"pk"`
	AdminId   		int			`form:"adminId" `
	StoreId   		int			`form:"storeId" `
	Remark   		string		`form:"remark" `
	ConfirmStatus  	int			`form:"confirmStatus" ` // 0 进行中1 已完成
	UpdateTime   	string		`form:"completeTime" `
	CreateTime		string		`form:"_" `

}

type ConfirmList struct {
	Confirm
	Number     		string		`form:"number" ` //餐厅编号
	UserName   		string		`form:"userName"`
	ConfirmRemark  	string		`form:"confirmRemark" `
	StoreStatus  	int			`form:"storeStatus" `
	WaitCheckNum	int
	CompleteNum		int
	FailNum			int
	NoNeedCheckNum	int
	TotalNum		int

}

func (this *Confirm) GetConfirmList()([]ConfirmList,error){

	var confirmList []ConfirmList
	_, err := orm.NewOrm().Raw("SELECT store.status as store_status, confirm_id,store.number,admin.user_name,confirm.update_time,confirm.confirm_status,confirm.remark as confirm_remark  FROM  confirm   inner join  store   on  confirm.store_id = store.store_id inner join admin on confirm.admin_id = admin.id order by confirm.create_time desc").QueryRows(&confirmList)
	if err != nil {
		return confirmList,err
	}

	//遍历添加 完成数 和未完成数
	for k,confirm :=  range confirmList {

		var softwareStatusList []orm.ParamsList

		//状态 0  没有核对 1  已经核对 2  安装失败 3  不需要核对
		orm.NewOrm().Raw("select  status,count(*) as count from software where   confirm_id= ?  group by status ", confirm.ConfirmId).ValuesList(&softwareStatusList)

		for _,val :=  range softwareStatusList{
			num ,_:= strconv.Atoi(val[1].(string))
			key ,_:= strconv.Atoi(val[0].(string))
			switch key {
				case 0:
					confirmList[k].WaitCheckNum = num
				case 1:
					confirmList[k].CompleteNum = num
				case 2:
					confirmList[k].FailNum 	=  num
				case 3:
					confirmList[k].NoNeedCheckNum 	=  num

			}
			confirmList[k].TotalNum += num
		}
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
		this.UpdateTime  = time.Now().Format("2006-01-02 15:04:05")
		if  _,err :=  orm.NewOrm().Update(this); err != nil{
			return err
		}
	}
	return nil
}


//更改该状态
func (this *Confirm)UpdateStatus(confirmId,status int)error  {

	this.ConfirmId 		= confirmId
	this.ConfirmStatus 	= status
	this.UpdateTime 	= time.Now().Format("2006-01-02 15:04:05")
	if _,err := orm.NewOrm().Update(this,"ConfirmStatus","UpdateTime") ;err != nil{

		return  err
	}else{

		return nil
	}
}






