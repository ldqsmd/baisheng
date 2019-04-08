package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type NewStore struct {
	Id          	int		`form:"id"`
	StoreId		    string	`form:"storeId"`
	CreateTime		string	`form:"createTime"`

}

func (this *NewStore) GetStoreList()([]Store,error){

	var storeList []Store
	_, err := orm.NewOrm().Raw("SELECT * from store where forbidden_status= 0").QueryRows(&storeList)
	if err == nil {
		return storeList,err
	}
	return storeList, nil
}

func (this *NewStore)AddStore()(int64,error) {
	this.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return orm.NewOrm().Insert(this)
}

func (this *NewStore)UpdateStore()error  {
	_,err :=  orm.NewOrm().Update(this)
	return err
}


//获取管理员信息
func (this *NewStore)GetStoreInfo(){
	//获取
	orm.NewOrm().Raw("SELECT * FROM store WHERE id=?", this.Id).QueryRow(&this)

}












