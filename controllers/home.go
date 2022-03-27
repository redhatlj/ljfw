package controllers

import (
	"cmdb/base/controlles/auth"
)

type HomeController struct {
	auth.AuthorizationController
}

func (c *HomeController) Prepare() {
	c.AuthorizationController.Prepare()
	c.Data["nav"] = "home"
}

func (c *HomeController) Index()  {
	c.TplName = "home/index.html"
}
