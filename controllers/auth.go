package controllers

import (
	"cmdb/base/controlles/base"
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/models"
	"github.com/beego/beego/v2/server/web"
	"net/http"
)


type AuthController struct {
	base.BaseController
}

// login 认证登录
func (c *AuthController) Login()  {
	sessionUser := c.GetSession(web.AppConfig.DefaultString("auth::sessionKey","user")) //判定是否是已登录，若登录则直接进入用户页面
	if sessionUser != nil{
		c.Redirect(c.URLFor(web.AppConfig.DefaultString("auth::HomeAction","UserController.Query")),http.StatusFound)
		return
	}

	from := &forms.LoginForm{}  //定义表单接收结构体
	errs := errors.New()       //定义错误
	// get 请求直接加载界面
	//POst 请求进行数据验证
	if c.Ctx.Input.IsPost() {
		// 验证成功
		//验证失败
		//获取用户提交的数据  name password  c.getstring("name)
		if err := c.ParseForm(from);err == nil{
			//fmt.Println(from,err)
			user := models.GetUserByName(from.Name)
			//fmt.Println(user)
			if user == nil {
				//用户不存在
				errs.AddError("default","用户名或密码错误！")
			} else if user.ValidPassword(from.Password)  {
					//密码正确
					//记录用户状态（session 记录服务器端）
				c.SetSession(web.AppConfig.DefaultString("auth::sessionKey","user"),user.ID)
				c.Redirect(web.URLFor(web.AppConfig.DefaultString("auth::HomeAction","UserController.Query")),302) //跳转页面
			} else {
				//用户密码不正确
				errs.AddError("default","用户名或密码错误！")

			}
		} else {
			errs.AddError("default","用户名或密码错误！")
		}

	}
	c.Data["from"] = from
	c.Data["errors"] = errs
	//定义加载页面
	c.TplName = "auth/login.html"
}

//Logout退出登录
func (c *AuthController) Logout()  {
	c.DestroySession()   //清除 全部session
	c.Redirect(c.URLFor(web.AppConfig.DefaultString("auth::LoginAction","AuthController.Login")),http.StatusFound)
}