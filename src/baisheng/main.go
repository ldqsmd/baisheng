package main

import (
	"baisheng/common"
	"baisheng/models"
	_ "baisheng/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func init() {
	common.StaicInit()
	common.MySQLInit()
}


func main() {

	var admin models.Admin
	data,total := admin.GetList()
	fmt.Println(data)
	fmt.Println(total)

	for k,v := range data{
		fmt.Println(k)
		fmt.Println(v)
	}
	beego.Run()
}

