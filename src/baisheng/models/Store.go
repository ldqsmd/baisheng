package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Store struct {
	Id          	int		`form:"storeId"`
	StoreName       string	`form:"storeName"`
	Number        	string	`form:"number"`
	Brand        	int		`form:"brand"`
	Status        	int		`form:"status"`
	Manager        	string	`form:"manager"`
	ManagerPhone    string	`form:"managerPhone"`
	Address        	string	`form:"address"`
	AreaManager     string	`form:"areaManager"`
	RegionalManager string	`form:"regionalManager"`
	Ip       		string	`form:"ip"`
	Network        	int		`form:"network"`
	OpenTime       	string	`form:"openTime"`
	DecorationTime 	string 	`form:"decorationTime"`
	CloseTime       string	`form:"closeTime"`
	BookStartTime	string  `form:"bookStartTime"`
	EmployeeStartTime	string  `form:"employeeStartTime"`
	WaitTime        string	`form:"waitTime"`
	DeviceTime 		string	`form:"deviceTime"`
	BuildName 		string	`form:"buildName"`
	TempCloseTime 	string	`form:"tempCloseTime"`
	Remark			string	`form:"remark"`
	ForbiddenStatus int	`form:"forbiddenStatus"`
	CreateTime 		string  `form:"createTime"`
}

func (this *Store) GetStoreList()([]Store,error){

	var storeList []Store
	_, err := orm.NewOrm().Raw("SELECT * from store where forbidden_status= 0").QueryRows(&storeList)
	if err == nil {
		return storeList,err
	}
	return storeList, nil
}

func (this *Store)AddStore()(int64,error) {
	this.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return orm.NewOrm().Insert(this)
}

func (this *Store)UpdateStore()error  {
	_,err :=  orm.NewOrm().Update(this)
	return err
}

//软删除
func (this *Store)DeleteStore()error  {

	this.ForbiddenStatus = 1
	_,err :=  orm.NewOrm().Update(this,"forbidden_status")
	return err
}



//获取管理员信息
func (this *Store)GetStoreInfo(){
	//获取
	orm.NewOrm().Raw("SELECT * FROM store WHERE id=?", this.Id).QueryRow(&this)

}












