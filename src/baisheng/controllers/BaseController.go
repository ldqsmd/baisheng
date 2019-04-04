package controllers

import (
	"baisheng/models"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string             //当前控制名称
	actionName     string             //当前action名称
	requestMethod  string             //当前接口请求方式
	adminInfo  models.Admin
}

func (this *BaseController) Prepare() {
	//附值
	this.controllerName, this.actionName = this.GetControllerAndAction()
	this.requestMethod = this.Ctx.Request.Method  //当前接口请求方式
	//从Session里获取数据 设置用户信息
	this.adapterAdminInfo()
	this.checkLogin()
}


//获取session admin信息
//适配到BaseController
func (base *BaseController) adapterAdminInfo() {
	adminInfo := base.GetSession("adminInfo")

	fmt.Println(base.GetSession("adminInfo"))


	if adminInfo != nil{
		base.adminInfo = adminInfo.(models.Admin)
		base.Data["adminInfo"]  = adminInfo
	}
}

//检查用户是否登录
//没有登录返回登录页面
func (base *BaseController) checkLogin() {


	if base.adminInfo.Id == 0 {
		//登录页面地址
		loginUrl := base.URLFor("LoginController.Login") + "?url="
		//登录成功后返回的址为当前
		returnURL := base.Ctx.Request.URL.Path
		//如果ajax请求则返回相应的错码和跳转的地址
		if base.Ctx.Input.IsAjax() {
			//由于是ajax请求，因此地址是header里的Referer
			returnURL := base.Ctx.Input.Refer()
			base.ReturnJson(302, "请登录", loginUrl+returnURL)
		}
		base.Redirect(loginUrl+returnURL, 302)
		base.StopRun()
	}

}


// 设置模板
// 第一个参数模板，第二个参数为layout
func (base *BaseController) SetTpl(template ...string) {

	var tplName string
	layout := "base/layout_base.html"
	switch {
		case len(template) == 1:
			tplName = template[0]
		case len(template) == 2:
			 layout  = template[0]
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

//JSON返回
func (base *BaseController) ReturnJson(code int,message string,data interface{}) {

	var  returnJson models.ReturnJson
	returnJson.Code	 	= code
	returnJson.Message 	= message
	returnJson.Data 	= data
	base.Data["json"] 	= &returnJson
	base.ServeJSON()
	base.StopRun()
}





//
func (c *BaseController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "error/404.html"
}

func (c *BaseController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "error/501.tpl"
}

func (c *BaseController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "error/dberror.tpl"
}