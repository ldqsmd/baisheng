package main

import (
	"baisheng/common"
	_ "baisheng/routers"
	"github.com/astaxie/beego"
)

func init() {
	common.MySQLInit()
}


func main() {
	common.ErrorInit()
	beego.Run()
}

