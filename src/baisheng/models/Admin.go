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


func (this *Admin) GetList()([]Admin,int64)  {


	var adminList []Admin
	//o := orm.NewOrm()
	//num, err := o.Raw("SELECT id,account,user_name,email,status FROM admin").QueryRows(&adminList)
	//if err == nil {
	//	fmt.Println("user nums: ", num)
	//}
	//fmt.Println(adminList)
	//return adminList,num
	query := orm.NewOrm().QueryTable(AdminTabelName())
	//默认排序
	//sortorder := "Id"
	//switch params.Sort {
	//case "Id":
	//	sortorder = "Id"
	//case "Seq":
	//	sortorder = "Seq"
	//}
	//if params.Order == "desc" {
	//	sortorder = "-" + sortorder
	//}
	//query = query.Filter("name__istartswith", params.NameLike)
	total ,err :=query.OrderBy("id").All(&adminList)

	if err != nil{

		fmt.Println(err.Error())
		fmt.Println(err.Error())
		fmt.Println(err.Error())
	}
	return adminList, total
}




//获取管理员信息
func (this *Admin)Login()error{

	err := orm.NewOrm().Raw("SELECT id,account,user_name,email,status FROM admin WHERE account=? and password=?", this.Account,this.Password).QueryRow(&this)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Println("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Println("Not row found")
	}
	//设置session记录
	return err
}


















