package forms

// Web前端提交密码修改表单接收结构体
// PasswordModifyForm 修改用户密码

type PasswordModifyForm struct {
	Old_password  string `form:"old_password"`
	New_password1 string `form:"new_password1"`
	New_password2 string `form:"new_password2"`
}
