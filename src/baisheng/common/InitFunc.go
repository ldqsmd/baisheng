package common

import (
	"baisheng/controllers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func MySQLInit()  {

	host 		:= beego.AppConfig.String("mysql::host")
	port 		:= beego.AppConfig.String("mysql::port")
	userName 	:= beego.AppConfig.String("mysql::userName")
	password 	:= beego.AppConfig.String("mysql::password")
	database 	:= beego.AppConfig.String("mysql::databse")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", userName+":"+password+"@tcp("+host+":"+ port+")/"+database+"?charset=utf8", 30)

	if  err != nil {
		fmt.Println(err.Error())
	}
}

func CheckNotie()  {
	turn 		:= beego.AppConfig.String("email::turn")
	if turn != "on"{
		fmt.Println("邮件轮询没有开启")
		return
	}
	fmt.Println("========================>邮件轮询已经开启<========================")

	tick := time.Tick(60 * time.Second)
	for {
		select {
			case <-tick:
				if err := controllers.NotcieAlertByEmail();err != nil {
					fmt.Println("通过邮件提醒错误"+err.Error())
				}else{
					fmt.Println("========================>没有需要发的邮件<========================")
				}
		}
	}
}

