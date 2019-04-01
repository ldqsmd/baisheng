package common

import "github.com/astaxie/beego"

func StaicInit()  {

	beego.SetStaticPath("/img","static/img")
	beego.SetStaticPath("/css","static/css")
	beego.SetStaticPath("/js","static/js")
	beego.SetStaticPath("/fonts","static/fonts")
}
