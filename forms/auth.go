package forms

//定义用户登录表单    数据验收结构体
type LoginForm struct {
	Name	 string 	`form:"name"`
	Password string 	`form:"password"`
}