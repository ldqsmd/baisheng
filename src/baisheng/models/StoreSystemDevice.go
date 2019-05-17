package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type StoreSystemDevice struct {
	StoreSystemDeviceId  int 	`json:"storeSystemDeviceId" form:"storeSystemDeviceId" orm:"pk"`
	StoreId 			int 	`json:"storeId" form:"storeId" `
	SystemDeviceId 		int 	`json:"systemDeviceId" form:"systemDeviceId" `
	Number 				int 	`json:"number" form:"number" `
	Price 				float64 `json:"price" form:"price" `
	SmallBook 			string 	`json:"smallBook" form:"smallBook" `
	OldDeviceNum 		int 	`json:"oldDeviceNum" form:"oldDeviceNum" `
	Remark 				string 	`json:"remark" form:"remark" `
	OrderNum 			string 	`json:"orderNum" form:"orderNum" `
	CreateTime			string  `json:"createTime"`
	UpdateTime			string  `json:"updateTime"`
}

type StoreSystemDeviceList struct {

	Supplier		string 	`json:"supplier" `
	SupplierJde		string 	`json:"supplierJde"  `
	DeviceName 		string 	`json:"deviceName"`
	DeviceJde		string 	`json:"deviceJde" `
	DeviceType		string 	`json:"deviceType"`
	TotalPrice		float64 `json:"totalPrice"`
	SystemDeviceId	int 	`json:"systemDeviceId" `
	SystemName		string 	`json:"systemName" `
	Active			int 	`json:"active" `
	StoreSystemDevice
}

func (this *StoreSystemDevice) TableName() string {
	return "store_system_device"
}




func (this *StoreSystemDevice) GetStoreSystemDeviceList(storeId int)([]StoreSystemDeviceList,error){

	var list []StoreSystemDeviceList
	sql := `SELECT 
			ssd.store_system_device_id,
			sd.system_device_id,
			sd.active,
			s.system_name,
			d.supplier,
			d.supplier_jde,
			ssd.order_num,
			d.device_name,
			d.device_jde,
			d.device_type,
			ssd.number,
			ssd.old_device_num,
			ssd.price,
			ifnull(ssd.price,0) * ifnull(ssd.number,0) as total_price,
			ssd.remark,
			ssd.small_book
			FROM  		store_system 		as ss  
			left join  	system 				as s 	on 	ss.system_id 		= s.system_id
			left join  	system_device 		as sd 	on 	ss.system_id 		= sd.system_id
			left join  	device 		 		as d 	on 	sd.device_id 		= d.device_id
			left join (select * from store_system_device  where  store_id = ?  ) as ssd  on sd.system_device_id 	= ssd.system_device_id

			where ss.store_id=?  order by ss.system_id,ssd.update_time desc `
	_,err := orm.NewOrm().Raw(sql,storeId,storeId).QueryRows(&list)
	return list, err
}


func (this *StoreSystemDevice)GetInfo(id int)StoreSystemDevice  {
	var storeSystemDevice StoreSystemDevice
	orm.NewOrm().Raw("select * from store_system_device where store_system_device_id=?",id).QueryRow(&storeSystemDevice)
	return storeSystemDevice
}


func (this *StoreSystemDevice)InsertOrUpdate()error  {

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	if this.StoreSystemDeviceId == 0 {
		this.CreateTime = nowTime
		this.UpdateTime = nowTime
		if id,err :=  orm.NewOrm().Insert(this); err != nil{
			return err
		}else{
			this.StoreSystemDeviceId  = int(id)
		}
	}else {
		this.UpdateTime = nowTime
		if  _,err :=  orm.NewOrm().Update(this); err != nil{
			return err
		}
	}
	return nil
}


func (this *StoreSystemDevice) List(store int)([]StoreSystemDeviceList,error){

	var list []StoreSystemDeviceList
	sql := `SELECT 
store.store_name,
sd.system_device_id,
s.system_name,
d.supplier,
d.supplier_jde,
ssd.order_num,
d.device_name,
d.device_jde,
d.device_type,
ssd.number,
ssd.price,
ifnull(ssd.price,0) * ifnull(ssd.number,0) as total_price,
ssd.remark,
ssd.small_book
FROM  		store_system_device as ssd
left join   store 				    on  ssd.store_id = store.store_id
left join  	system_device 		as sd 	on 	ssd.system_device_id = sd.system_device_id
left join  	device 		 		as d 	on 	sd.device_id 		= d.device_id
left join  	system 				as s 	on 	s.system_id 		= sd.system_id
order by store.store_id`
	_,err := orm.NewOrm().Raw(sql).QueryRows(&list)
	return list, err
}





