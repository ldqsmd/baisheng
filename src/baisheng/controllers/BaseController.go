package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string             //当前控制名称
	actionName     string             //当前action名称
}

func (c *BaseController) Prepare() {
	//附值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	//从Session里获取数据 设置用户信息
	//c.adapterUserInfo()
}

// 设置模板
// 第一个参数模板，第二个参数为layout
func (base *BaseController) setTpl(template ...string) {

	var tplName string
	layout := "base/layout_base.html"
	switch {
		case len(template) == 1:
			tplName = template[0]
		case len(template) == 2:
			 layout = template[0]
			 tplName = template[1]
		default:
			//不要Controller这个10个字母
			ctrlName := strings.ToLower(base.controllerName[0 : len(base.controllerName)-10])
			actionName := strings.ToLower(base.actionName)
			tplName = ctrlName + "/" + actionName + ".html"
	}
	base.Layout = layout
	base.TplName = tplName
}