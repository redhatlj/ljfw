package auth

import (
	"cmdb/base/controlles/base"
	"cmdb/models"
	"github.com/beego/beego/v2/server/web"
	"net/http"
	"strings"
)
//认证业务逻辑控制器
type AuthorizationController struct {
	base.BaseController
	LoginUser *models.User
}

// nav 判断 从控制器获取首部分字符
func (c *AuthorizationController) getNay() string {
	controllerName,_ := c.GetControllerAndAction()
	return strings.ToLower(strings.TrimSuffix(controllerName,"Controller"))
}

// session 认证检查
func (c *AuthorizationController) Prepare()  {
	 sessionValus := c.GetSession(web.AppConfig.DefaultString("auth::sessionKey","user"))
	 c.Data["LoginUser"] = nil
	 c.Data["nav"] = c.getNay()

	if sessionValus != nil {
		 if id, ok := sessionValus.(int); ok {
			 if user := models.GetUserById(id); user != nil {
				 c.Data["LoginUser"] = user
				 c.LoginUser = user
				 return
			 }
		 }
	 }
	c.Redirect(c.URLFor(web.AppConfig.DefaultString("auth::LoginAction","AuthController.Login")),http.StatusFound )

	//查询用户信息 --》 Data
	 // 通过ID获取用户数据
}