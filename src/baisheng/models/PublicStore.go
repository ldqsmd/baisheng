package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


type PublicStore struct {

	StoreId         int		`form:"storeId" orm:"pk"`
	StoreRemark 	string	`form:"storeRemark"`
	StoreCreateTime string  `form:"storeCreateTime"`
	StoreUpdateTime	string  `form:"_"`
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
	ForbiddenStatus int	`form:"forbiddenStatus"`
}

func (this *PublicStore) TableName() string {
	return "store"
}

func (this *PublicStore) GetStoreList()([]PublicStore,error){
	var storeList []PublicStore
	_, err := orm.NewOrm().Raw("SELECT * from store where forbidden_status= 0").QueryRows(&storeList)
	if err == nil {
		return storeList,err
	}
	return storeList, nil
}

func (this *PublicStore)InsertOrUpdate()error {

	if this.StoreId == 0 {
		this.StoreCreateTime  = time.Now().Format("2006-01-02 15:04:05")
		storeId,err :=  orm.NewOrm().Insert(this)
		if err != nil{
			return err
		}
		this.StoreId = int(storeId)
	}else {
		this.StoreUpdateTime  = time.Now().Format("2006-01-02 15:04:05")
		_,err :=  orm.NewOrm().Update(this)
		if err != nil{
			return err
		}
	}
	return  nil
}
//软删除
func (this *PublicStore)DeleteStore()error  {

	this.ForbiddenStatus = 1
	_,err :=  orm.NewOrm().Update(this,"forbidden_status")
	return err
}
//获取管理员信息
func (this *PublicStore)GetStoreInfo(storeId string)PublicStore{
	//获取
	orm.NewOrm().Raw("SELECT * FROM store WHERE store_id=?",storeId).QueryRow(&this)
	return *this
}












