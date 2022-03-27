package main

import (
	_ "cmdb/controllers"
	_ "cmdb/routers"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"log"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/cmdb?charset=utf8mb4&loc=PRC&parseTime=True",
		web.AppConfig.DefaultString("mysql::mysqluser","root"),
		web.AppConfig.DefaultString("mysql::mysqlpass","123456"),
		web.AppConfig.DefaultString("mysql::mysqlurls","127.0.0.1"),
		web.AppConfig.DefaultInt("mysql::mysqlport",3306))
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",dsn)
	//测试数据库连接
	if db,err := orm.GetDB("default");err != nil{
		log.Fatal(err)
	} else if err := db.Ping();err != nil {
		log.Fatal(err)
	}


	web.Run()
}
