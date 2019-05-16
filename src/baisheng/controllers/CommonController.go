package controllers

import (
	"baisheng/models"
	"github.com/astaxie/beego/orm"
)

// 通过map主键唯一的特性过滤重复元素
func FilterStrSlice(slc []string) []string {

	result := []string{}
	tempMap := map[string]byte{}  // 存放不重复主键
	for _, e := range slc{
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l{  // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

func GetStatusList()[]models.StoreStatus{

	var statusList []models.StoreStatus
	_, err := orm.NewOrm().Raw("SELECT * from store_status").QueryRows(&statusList)
	if err == nil {
		return statusList
	}
	return statusList
}


