package main

import (
	"baisheng/common"
	_ "baisheng/routers"
	"github.com/astaxie/beego"
)

func init() {
	common.MySQLInit()
	go common.CheckNotie()
}

func main() {
	//beego.ErrorHandler("404",page_not_found)
	beego.Run()
}

