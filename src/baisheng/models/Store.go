package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Store struct {
	PublicStore
	NewStore
	IeStore
	CloseStore
	SystemId  []int `json:"systemId" form:"systemId"`
}
//新增
func (this *Store)InsertOrUpdate()(error) {

	var storeSystem  StoreSystem
	var numberCount  int
	orm.NewOrm().Raw("SELECT count(*)  from store where number=?", this.Number).QueryRow(&numberCount)
	if numberCount>1 {
		return errors.New("餐厅编号已经存在")
	}
	//默认指定 标识为 特别关注
	this.SignFlag = 2
	if err := this.PublicStore.InsertOrUpdate(); err != nil{
		return err
	}
	this.NewStore.NewStoreId 		= this.PublicStore.StoreId
	this.IeStore.IeStoreId   		= this.PublicStore.StoreId
	this.CloseStore.CloseStoreId   	= this.PublicStore.StoreId

	//选择添加餐厅系统
	for _,systemId := range this.SystemId{
		storeSystem.SystemId  = systemId
		storeSystem.StoreId  = this.StoreId
		if err := storeSystem.InsertOrUpdate();err != nil {
			return err
		}
	}


	//NEW store
	if  this.PublicStore.Status   == 1 {
		//默认 不标记
		this.SignFlag = 1
		 if err := this.NewStore.InsertOrUpdate();err != nil {
			return  err
		}
	}
	//IE store
	if  this.PublicStore.Status  == 2 {
		if err := this.IeStore.InsertOrUpdate();err != nil {
			return  err
		}
	}
	//clsose store
	if  this.PublicStore.Status  == 3 {
		if err := this.CloseStore.InsertOrUpdate();err != nil {
			return  err
		}
	}

	return nil
}


func (this *Store)GetStoreInfo(storeId string)(Store) {

	var store  Store
	orm.NewOrm().Raw("SELECT * FROM store as s left join ie_store as i on s.store_id = i.ie_store_id  left join new_store as n on s.store_id=n.new_store_id WHERE s.store_id=?", storeId).QueryRow(&store)

	fmt.Println(store)
	return store
}




