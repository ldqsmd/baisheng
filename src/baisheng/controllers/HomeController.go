package controllers

type HomeController struct {
	BaseController
}

func (home *HomeController) Index() {
	//home.Data["Website"] = "beego.me"
	//home.Data["Email"] = "astaxie@gmail.com"
	home.setTpl()
}

func (home *HomeController)Content() {
	home.setTpl()
}
