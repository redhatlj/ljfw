package routers
//当如包一般流程如下：
//先导入内置包
//第三方
//当前项目
import (
	"cmdb/controllers"
	"github.com/beego/beego/v2/server/web"
	"log"
)

func init()  {
	web.Router("/",&controllers.HomeController{},"*:Index")

	web.AutoRouter(&controllers.PasswordController{})
	web.AutoRouter(&controllers.AuthController{})
	web.AutoRouter(&controllers.HomeController{})
	web.AutoRouter(&controllers.UserController{})
	web.ErrorController(&controllers.ErrorController{})

	log.Println("init函数执行成功")
}