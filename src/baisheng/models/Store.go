package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Store struct {
	PublicStore
	NewStore
	IeStore
	CloseStore
}
//新增
func (this *Store)InsertOrUpdate()(error) {

	//public store
	if err := this.PublicStore.InsertOrUpdate(); err != nil{
		return err
	}
	this.NewStore.NewStoreId 		= this.PublicStore.StoreId
	this.IeStore.IeStoreId   		= this.PublicStore.StoreId
	this.CloseStore.CloseStoreId   	= this.PublicStore.StoreId

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

	orm.NewOrm().Raw("SELECT * FROM store as s left join ie_store as i on s.store_id = i.ie_store_id  left join new_store as n on s.store_id=n.new_store_id WHERE s.store_id=?", storeId).QueryRow(this)
fmt.Println(*this)
	return *this
}





