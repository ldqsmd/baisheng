package controllers

type HomeController struct {
	BaseController
}

func (home *HomeController) Index() {
	home.SetTpl()
}

func (home *HomeController)Content() {
	home.SetTpl()
}
