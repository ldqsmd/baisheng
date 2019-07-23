package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Notice struct {
	Id          	int		`form:"id"`
	Email        	string	`form:"email"`
	Subject     	string	`form:"subject"`
	Body     		string	`form:"body"`
	NoticeTime      string 	`form:"noticeTime"`
	PlanNoticeTime  int64 	`form:"planNoticeTime"`
	Status     		int
	CreateTime     	string
}

func (this *Notice) List()([]Notice,int64)  {
	var notice []Notice
	query := orm.NewOrm().QueryTable("notice")
	total ,_ :=query.OrderBy("id").OrderBy("-status").All(&notice)
	return notice, total
}


func (this *Notice) Detail(id string)(Notice){
	var notice  Notice
	orm.NewOrm().Raw("select * from notice where id=?",id).QueryRow(&notice)
	return notice
}

func (this *Notice) Delete( )error{
	_,err := orm.NewOrm().Delete(this)
	if err != nil {
		return err
	}
	return nil
}

func (this *Notice)InsertOrUpdate()error  {

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	this.CreateTime = nowTime

	if this.Id == 0 {
		this.PlanNoticeTime = StrToMicTimeInt64(this.NoticeTime)
		if id,err :=  orm.NewOrm().Insert(this); err != nil{
			return err
		}else{
			this.Id  = int(id)
		}
	}else {
		this.PlanNoticeTime = StrToMicTimeInt64(this.NoticeTime)
		if  _,err :=  orm.NewOrm().Update(this); err != nil{
			return err
		}
	}
	return nil
}


func  GetNoticeToList()([]Notice,error ) {
	var notice []Notice
	nowTime := time.Now().Unix()
	cond := orm.NewCondition().And("status__eq",0).And("plan_notice_time__lt",nowTime)
	if _,err := orm.NewOrm().QueryTable("notice").SetCond(cond).All(&notice);err != nil {
		return notice, err
	}
	return notice, nil
}

func  (this *Notice) UpdateNoticeStatus( )error {
	this.Status = 1
	if  _,err :=  orm.NewOrm().Update(this,"status"); err != nil{
		return  err
	}
	return   nil
}

func StrToMicTimeInt64(s string) int64 {
	//获取本地location   	//待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                    //转化所需模板
	loc, _ := time.LoadLocation("Local")                   //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, s, loc) //使用模板在对应时区转化为time.time类型
	return theTime.Unix()
}








