package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "myproject/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"

)

func main() {
	beego.Run()
}

var o orm.Ormer

var aliasName string = "default"

func init(){
	connStatus,err :=ConnectDb()
	if connStatus{
		fmt.Printf("数据库连接成功！\n")
		o = orm.NewOrm()
	}else{
		fmt.Printf("数据库连接失败！\n错误是：%s\n", err)
	}
}

func ConnectDb() (bool,error){
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	db := beego.AppConfig.String("db")

	//注册mysql Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//构造conn连接
	//用户名:密码@tcp(url地址)/数据库
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8"
	//注册数据库连接
	err :=orm.RegisterDataBase(aliasName, "mysql", conn)
	if err != nil {
		return false,err
	}
	//fmt.Printf("数据库连接成功！%s\n", conn)
	return true,nil
}