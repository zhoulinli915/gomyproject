package models

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

// 用户
type College struct{
    Id                          uint      `orm:"column(id);pk;size(11)"` // 设置主键
    MainId                      uint      `orm:"column(main_id);size(11)"` // 设置对应关系
    WxSystemUserId              uint      `orm:"column(wx_system_user_id);size(11)"` // 设置对应关系
    CollegeType                 string    `orm:"column(college_type);size(11)"` // enum 设置对应关系
    CollegeName                 string    `orm:"size(255);column(college_name)"`
    CollegeCode                 string    `orm:"size(100);column(college_code)"`
    CreaterUserId               uint      `orm:"column(creater_user_id);size(11)"` // 设置对应关系
    CreaterUserType             string    `orm:"column(creater_user_type);size(1)"` // 设置对应关系
    CreateTime                  uint      `orm:"column(create_time);size(11)"` // 设置对应关系
    UpdateTime                  uint      `orm:"column(update_time);size(11)"` // 设置对应关系
    IsDel                       string    `orm:"column(is_del);size(1)"`
    /*
id
main_id
wx_system_user_id
college_type
college_name
college_code
creater_user_id
creater_user_type
create_time
update_time
is_del
    */
}
//region 衍生的表创建语句 类型的对应关系或查看文档:https://beego.me/docs/mvc/model/models.md
//-- --------------------------------------------------
//--  Table Structure for `myproject/models.College`
//-- --------------------------------------------------
//CREATE TABLE IF NOT EXISTS `college` (
//`id` bigint NOT NULL PRIMARY KEY,
//`main_id` smallint unsigned NOT NULL DEFAULT 0 ,
//`wx_system_user_id` smallint unsigned NOT NULL DEFAULT 0 ,
//`college_type` smallint unsigned NOT NULL DEFAULT 0 ,
//`college_name` varchar(100) NOT NULL DEFAULT '' ,
//`college_code` varchar(100) NOT NULL DEFAULT '' ,
//`creater_user_id` smallint unsigned NOT NULL DEFAULT 0 ,
//`creater_user_type` smallint unsigned NOT NULL DEFAULT 0 ,
//`create_time` smallint unsigned NOT NULL DEFAULT 0 ,
//`update_time` smallint unsigned NOT NULL DEFAULT 0 ,
//`is_del` varchar(1) NOT NULL DEFAULT ''
//) ENGINE=InnoDB;
//endregion


//根据用户数据列表
func (a *College) DataList() (colleges []College) {


    o := orm.NewOrm()
    qs := o.QueryTable(TableName)
    orm.Debug = true  //todo 是每次自己定义变量，单独调用还是可以main中统一定义，然后这里使用？

    var us []College
    cnt, err :=  qs.OrderBy("-id").Limit(10, 0).All(&us)//Filter("Id", 0).
    if err == nil {
       fmt.Printf("count", cnt)
        //os.Stdout.WriteString(err.Error()) //输出错误
        // 使用不当，可能引发  Handler crashed with error runtime error: invalid memory address or nil pointer dereference  这个问题可以查看日志知道具体出在哪个controller的第几行
    }


    //sql := "SELECT * FROM app_ielts_college_institutions"
    // o.Raw(sql).QueryRows(&us)
    return us
}

var TableName string = "app_ielts_college_institutions"
func (x *College) TableName()string{
    return TableName
}

//初始化模型
func init() {
    // 需要在init中注册定义的model //可以逗号分隔多个
    // 可以增加表前缀 orm.RegisterModelWithPrefix("prefix_", new(College)) 创建后的表名为 prefix_user
    //todo  是否可以定义表名
    orm.RegisterModel(new(College))


    //region create table
    // verbose=true可以显示表的创建sql 如果已经存在，以后提示创建过了 ,false 不显示任何信息
    // 根据模型的名字，小写创建表 应该是在main中运行，会创建所有model的表 可以用于验证数据库连接的一种笨办法
    //todo 需要知道数据类型和mysql的对应关系，以及mysql对应字段的添加方式，例如：default值
    //orm.RunSyncdb("default", false, false)
    //rendregion
}