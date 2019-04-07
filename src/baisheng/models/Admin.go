package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id          	int		`form:"-"`
	Account        	string	`form:"account"`
	UserName    	string	`form:"userName"`
	Password     	string	`form:"password"`
	Email     		string 	`form:"email"`
	Status     		int
	CreateTime     	string
}



func (this *Admin) GetAdminList()([]Admin,int64)  {

	var adminList []Admin
	query := orm.NewOrm().QueryTable(AdminTabelName())
	total ,err :=query.OrderBy("id").All(&adminList)

	if err != nil{

		fmt.Println(err.Error())
		fmt.Println(err.Error())
		fmt.Println(err.Error())
	}
	return adminList, total
}

func (this *Admin)AddAdmin(){
	o := orm.NewOrm()
	_,err:=  o.Insert(this)

	if err != nil {
		fmt.Println(err.Error())

	}
}

func (this *Admin)UpdateAdmin()error  {
	_,err := orm.NewOrm().Update(this)

	return err
}


//获取管理员信息
func (this *Admin)GetAdminInfo()error{
	//登录获取用户信息
	if this.Id == 0 {
		return orm.NewOrm().Raw("SELECT id,account,user_name,email,status FROM admin WHERE account=? and password=?", this.Account,this.Password).QueryRow(&this)
	}else{
		//根据adminId获取用户信息
		return orm.NewOrm().Raw("SELECT id,account,user_name,email,status FROM admin WHERE id=?", this.Id).QueryRow(&this)

	}
}












