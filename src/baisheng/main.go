package main

import (
	"baisheng/common"
	_ "baisheng/routers"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"net/http"
)

func init() {
	common.MySQLInit()
}
func page_not_found(rw http.ResponseWriter, r *http.Request){

	t,err:= template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/error/404.html")

	if err != nil {
		fmt.Println(err.Error())
	}
	data :=make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

func main() {
	//beego.ErrorHandler("404",page_not_found)
	beego.Run()
}

