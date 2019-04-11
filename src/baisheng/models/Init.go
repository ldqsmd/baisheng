package models

import (
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(new(Admin),new(PublicStore),new(IEStore),new(NewStore),new(CloseStore))
}


func AdminTabelName() string {
	return "admin"
}
