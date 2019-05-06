package models

import (
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(new(Admin),new(PublicStore),new(IeStore),new(NewStore),new(CloseStore),new(Confirm),new(Software),new(DeviceCheck))
}


func AdminTabelName() string {
	return "admin"
}
