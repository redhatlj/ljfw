package controllers

import (
	"cmdb/base/controlles/auth"
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/models"
	"fmt"
	"regexp"
)

// PasswordController 用户修改密码控制器
type PasswordController struct {
	auth.AuthorizationController
}

func (c *PasswordController) Modify() {
	form := &forms.PasswordModifyForm{}
	errs := errors.New()
	text := ""  // 修改密码成功的标识

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			// 验证
			if ok := c.LoginUser.ValidPassword(form.Old_password); !ok {
				errs.AddError("default", "旧密码错误")
				fmt.Println("旧密码错误返回",c.Ctx.Input.IsPost())
			} else {
				//验证密码范围数字，大小写英文。特殊字符（_ . $ ! % ^ & * ( ) +）
				passwordRegex := "^[0-9a-zA-Z_.\\$\\!#%^&\\*\\(\\)\\+]{6,10}$"
				fmt.Println(regexp.MatchString(passwordRegex, form.New_password1))
				if isMatch, _ := regexp.MatchString(passwordRegex, form.New_password1); !isMatch {
					errs.AddError("default", "密码只能由数字，大小写英文。特殊字符（_ . $ ! % ^ & * ( ) +）组成")
				} else if form.New_password1 != form.New_password2 {
					errs.AddError("default", "输入的两次密码不一致")
				} else if form.Old_password == form.New_password1 {
					errs.AddError("default", "新旧密码不能一致")
				} else {
					models.ModifyUserPassword(c.LoginUser.ID, form.New_password2)
					text = "密码修改成功"
				
				}
			}
		}

	}
	c.TplName = "password/modify.html"
	c.Data["errors"] = errs
	c.Data["text"] = text

}
