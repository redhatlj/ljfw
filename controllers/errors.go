package controllers

import "github.com/beego/beego/v2/server/web"

type ErrorController struct {
	web.Controller
}

//定义错误处理
func (c *ErrorController) Error404()  {
	c.Data["Content"] = "page not found"
	c.TplName = "error/404.html"
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "error/501.html"
}


//func (c *ErrorController) ErrorDb() {
//	c.Data["content"] = "database is now down"
//	c.TplName = "error/dberror.tpl"
//}