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
func (this *BaseController) adapterAdminInfo() {
	adminInfo := this.GetSession("adminInfo")
	if adminInfo != nil{
		this.adminInfo = adminInfo.(models.Admin)
		this.Data["adminInfo"]  = adminInfo
	}
}

//检查用户是否登录
//没有登录返回登录页面
func (this *BaseController) checkLogin() {

	fmt.Println(this.GetSession("adminInfo"))

	if this.adminInfo.Id == 0 {

		//登录页面地址
		loginUrl := this.URLFor("LoginController.Login") + "?url="
		//登录成功后返回的址为当前
		returnURL := this.Ctx.Request.URL.Path
		//如果ajax请求则返回相应的错码和跳转的地址
		if this.Ctx.Input.IsAjax() {
			//由于是ajax请求，因此地址是header里的Referer
			returnURL := this.Ctx.Input.Refer()
			this.ReturnJson(302, "请登录", loginUrl+returnURL)
		}
		this.Redirect(loginUrl+returnURL, 302)
		this.StopRun()
	}

}


// 设置模板
// 第一个参数模板，第二个参数为layout
func (this *BaseController) SetTpl(template ...string) {

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
			ctrlName := strings.ToLower(this.controllerName[0 : len(this.controllerName)-10])
			actionName := strings.ToLower(this.actionName)
			tplName = ctrlName + "/" + actionName + ".html"
	}

	this.Layout = layout
	this.TplName = tplName
}

//JSON返回
func (this *BaseController) ReturnJson(code int,message string,data interface{}) {

	var  returnJson models.ReturnJson
	returnJson.Code	 	= code
	returnJson.Message 	= message
	returnJson.Data 	= data
	this.Data["json"] 	= &returnJson
	this.ServeJSON()
	this.StopRun()
}

//
func (this *BaseController) Error404() {
	//this.Data["content"] = "page not found"
	this.TplName = "error/404.html"
}

func (c *BaseController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "error/501.tpl"
}

func (c *BaseController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "error/dberror.tpl"
}