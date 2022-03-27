package controllers

import (
	"cmdb/base/controlles/auth"
	"cmdb/forms"
	"cmdb/models"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"log"
	"net/http"
)

//用户管理控制器
type UserController struct {
	auth.AuthorizationController
}
//web 导航列表显示
func (c *UserController) Prepare() {
	c.AuthorizationController.Prepare()
	c.Data["nav"] = "user"
}

//  Query 查询用户
func  (c *UserController) Query()   {
	q := c.GetString("q")
	users := models.QueryUser(q)
	c.Data["users"] = users
	c.Data["q"] = q
	c.TplName = "user/query.html"
}

// modify 修改用户
func (c *UserController) Modify()  {
	form := &forms.UserModifyForm{}
	//get 获取数据
	// Post 修改用户
	if c.Ctx.Input.IsPost(){
		if err := c.ParseForm(form) ; err == nil{
			//验证数据  如果Id存在
			models.ModifyUser(form)
			c.Redirect(web.URLFor("UserController.Query"),http.StatusFound)
		}


	} else if id,err := c.GetInt("id");err == nil {
		if user := models.GetUserById(id);user != nil{
			form.ID = user.ID
			form.Name = user.Name
		}
	}

	c.Data["form"] = form
	c.TplName = "user/modify.html"
}
// Delete 删除用户
func (c *UserController) DeleteDBUser()  {
	if id,err := c.GetInt("id"); err == nil && c.LoginUser.ID != id{
		log.Println("执行删除用户操作")
		models.DeleteUser(id)
	} else {
		fmt.Println(err)
	}
	c.Redirect(web.URLFor("UserController.Query"),http.StatusFound)
}