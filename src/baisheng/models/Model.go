package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Model struct {
	ModelId 		int 	`json:"modelId" form:"modelId"  orm:"pk"`
	ModelName 		string 	`json:"modelName" form:"modelName" `
	Remark 			string 	`json:"remark" form:"remark" `
	CreateTime 		string  `json:"createTime"`
	UpdateTime 		string  `json:"updateTime"`
}

func (this *Model) TableName() string {
	return "model"
}

func (this *Model) GetModelList()([]Model,error){

	var modelList []Model
	_, err := orm.NewOrm().Raw("select * from model order by create_time desc").QueryRows(&modelList)
	if err == nil {
		return modelList,err
	}
	return modelList, nil
}


func (this *Model) GetModelInfo(modelId string)(Model){

	var model  Model
	orm.NewOrm().Raw("select * from model where model_id=?",modelId).QueryRow(&model)
	return model
}

func (this *Model)InsertOrUpdate()error  {

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	if this.ModelId == 0 {
		this.CreateTime = nowTime
		this.UpdateTime = nowTime
		if id,err :=  orm.NewOrm().Insert(this); err != nil{
			return err
		}else{
			this.ModelId  = int(id)
		}
	}else {
		this.UpdateTime = nowTime
		if  _,err :=  orm.NewOrm().Update(this,"ModelName","Remark","UpdateTime"); err != nil{
			return err
		}
	}
	return nil
}


func (this *Model)DeleteModel()error  {

	if  _,err :=  orm.NewOrm().Delete(this); err != nil{
		return err
	}
	return nil
}












