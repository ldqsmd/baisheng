package models

import (
	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id          	int
	Account        	string
	UserName    	string
	Password     	string
	Status     		int
	CreateTime     	string
}


func (this *Admin) GetList()([]Admin,int64)  {

	query := orm.NewOrm().QueryTable(AdminTabelName())
	data := make([]Admin, 0)
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

	total, _ := query.Count()
	query.OrderBy("id").All(&data)
	return data, total
}