package models

import (
	"cmdb/forms"
	"cmdb/utils"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// User 用户对象
type User struct {
	ID         int        `orm:"column(id)"`
	StaffID    string     `orm:"column(staffID);size(32)"`
	Name       string     `orm:"size(64)"`
	Nickname   string     `orm:"size(64)"`
	Password   string     `orm:"size(1024)"`
	Gender     int        `orm:""`
	Tel        string     `orm:"size(32)"`
	Addr       string     `orm:"size(64)"`
	Email      string     `orm:"column(Email);size(64)"`
	Department string     `orm:"column(department);size(128)"`
	Status     int        `orm:""`
	Created_at *time.Time `orm:"column(created_at);auto_now_add"`
	Updated_at *time.Time `orm:"column(updated_at);auto_now"`
	DeletedAt  *time.Time `orm:"column(deleted_at);null"`
}

func init() { //注册模型
	orm.Debug = true
	orm.RegisterModel(new(User))
}

//GenderText web模板男女状态判定
func (u *User) GenderText() string {
	if u.Gender == 0 {
		return "女"
	}
	if u.Gender == 1 {
		return "男"
	}
	return "性别设置错误，应设置0/1"
}

//状态判定
func (u *User) StatusText() string {
	switch u.Status {
	case 0:
		return "正常"
	case 1:
		return "锁定"
	case 2:
		return "离职"
	default:
		return "状态设置错误"
	}
}

//时间格式化
func (u *User) TimeCRText() string {
	return u.Created_at.Format("2006-01-02 15:04")
}
func (u *User) TimeUPText() string {
	return u.Updated_at.Format("2006-01-02")
}

//验证用户密码是否正确
func (u *User) ValidPassword(password string) bool {
	return utils.CheckPassword(password, u.Password)
	//fmt.Println( password , u.Password)
	//return u.Password == utils.Md5Text(password)
}

// 通过用户名获取用户
func GetUserByName(name string) *User {
	user := &User{Name: name}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

//查询用户
func QueryUser(q string) []*User {
	var users []*User
	queryset := orm.NewOrm().QueryTable(User{}) //获取表
	if q != "" {
		cond := orm.NewCondition()
		cond = cond.Or("name__icontains", q)
		cond = cond.Or("nickname__icontains", q)
		cond = cond.Or("tel__icontains", q)
		cond = cond.Or("addr__icontains", q)
		cond = cond.Or("Email__icontains", q)
		cond = cond.Or("department__icontains", q)
		queryset = queryset.SetCond(cond)
	}
	queryset.All(&users)
	return users
}

//GetUserById 通过用户ID获取用户信息
func GetUserById(Id int) *User {
	user := &User{ID: Id}
	ormer := orm.NewOrm()
	if err := ormer.Read(user); err == nil {
		return user
	}
	return nil
}

//ModifyUser  修改用户信息
func ModifyUser(form *forms.UserModifyForm) {
	if user := GetUserById(form.ID); user != nil {
		user.Name = form.Name
		ormer := orm.NewOrm()
		ormer.Update(user, "Name")
	}
}

//DeleteUser 删除用户
func DeleteUser(id int) {
	ormer := orm.NewOrm()
	fmt.Println(ormer.Read(&User{ID: id}))
	ormer.Delete(&User{ID: id},"id")
}

// 用户修密码
func ModifyUserPassword(id int,password string)  {
	if user := GetUserById(id); user != nil{
		user.Password = utils.GeneratePassword(password) // 通过算法加密密码然后存储数据库
		ormer := orm.NewOrm()
		ormer.Update(user,"Password")
	}
}