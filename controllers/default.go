package controllers

import (
	//"apiproject/models"
	"fmt"
	"github.com/astaxie/beego"
	"myproject/models" //引入本工程的models 否则导致调用报红
	//"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//region -------------------IndexController-----------访问/index.html-----------------------
//todo 验证码的文件  表单提交
type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	//可以去	beego.Controller获取对应的属性
	//Data 参数可以自定义
	c.Data["title"] = "首页"
	c.Data["action"] = "/savedata"
	c.TplName = "trueindex.tpl"
}
type Users struct {
	username string
	pwd  string
}
func (c *IndexController) SaveData(){
	//fmt.Println("------------------------<br>")//输出到命令行
	username := c.GetString("username")
	pwd := c.GetString("pwd")

	var isok string
	c.Ctx.Input.Bind(&isok, "submit")

	//Ctx 就是 Context
	//c.Ctx.WriteString("hello world\n")//输出到页面
	c.Ctx.WriteString("username="+username+"   pwd="+pwd+"  submit="+isok+"\n")//生成的html不解析 响应头是：Content-Type: text/plain
	//c.Ctx.WriteString("username="+username+",pwd="+pwd+"<br>")
	//c.Ctx.WriteString(c.Ctx.Input.Method()+"\n")
	//data := c.Ctx.Input.RequestBody
	//arr := [3][4]string{{"1","3","4"}, {"5","6","8"}, {"9","10","11","12"}} //初始值为0
	//tree,_,error := c.MakeTree(arr)
	//c.Ctx.WriteString(c.Ctx.Input.Method()+"\n")

	value :=c.Input() //=c.Ctx.Request.Form //map[string][]string
	for k, v := range value {
		fmt.Printf("%s=%s;", k, v)
		str :=fmt.Sprintf("%s=%s;",k, v)
		c.Ctx.WriteString(str+"\n")
	}


	c.Ctx.ResponseWriter.Write([]byte("username="+username+",pwd="+pwd+"------<br>"))//依然不支持html
	if c.Ctx.Input.Method() == "POST" {
		//c.Ctx.WriteString(c.Ctx.Input.Method()+"----1------\n")
	}else if c.Ctx.Input.Method() == "GET" {
		//c.Ctx.WriteString(c.Ctx.Input.Method()+"----2------\n")
	}else{
		//c.Ctx.WriteString(c.Ctx.Input.Method()+"----3------\n")
	}

	//var user Users
	user := Users{}

	if error := c.ParseForm(&user);error != nil {
		c.Ctx.WriteString("出错了！")
	}else {
		c.Ctx.WriteString("我是结构体\n")
		c.Ctx.WriteString("Name=" + user.username + "\nPwd=" + user.pwd)

	}
	//出入数据库
	//UserModel := model.User{}

	//跳转到显示列表页
}

func (this *IndexController)UserList(){
	//this.Layout = "public/header.tpl" //只加载一个header.tpl，如果下面有this.TplName,，则只显示this.TplName //Layout是母版的设计思想
	//模板加载  {{template "public/footer.tpl" .}} 记得后面有个点，否则会导致 页面可以加载，变量死活加载不了
	this.Data["title"] = "用户列表"
	this.Data["collegelist"] = nil //输出if可以可以判断nil(不用写，用if .collegelist 就可以)
	//todo 制作model，并加载列表
	var collegeModel models.College////引入本工程的models 否则导致调用报红 有College struct就可以，与文件名无关
	Data := collegeModel.DataList()
	if Data != nil{
		this.Data["collegelist"] = Data
		//this.Ctx.WriteString("----have data------\n")
		//for _,college := range Data{
		//	this.Ctx.WriteString("---"+college.GetTableName()+"------\n")
		//	//fmt.Println(college.GetTableName())
		//}
	}else{
		fmt.Println("----no data------\n")
	}

	//Datas := models.OrderNum(6)
	//if Datas != ""{
	//	this.Ctx.WriteString("----have data------\n")
	//	//for _,college := range Data{
	//	//	this.Ctx.WriteString("---"+college.GetTableName()+"------\n")
	//	//	//fmt.Println(college.GetTableName())
	//	//}
	//}else{
	//	fmt.Println("----no data------\n")
	//}
	//var maps []orm.Params
	//num, err := o.Raw("SELECT * FROM app_ielts_college_institutions").Values(&maps)
	//if err != nil{
	//	if num >=1 {
	//		fmt.Println("查询到"+num+"条记录\n")
	//		for _,term := range maps{
	//			fmt.Println(term["id"],":",term["name"])
	//		}
	//	}else{
	//		fmt.Println("查询到"+num+"条记录\n")
	//	}
	//
	//}else{
	//	fmt.Println("表查询失败\n")
	//}

}

//func (this *IndexController) MakeTree(arr []string)([]string,bool,error){
//	return arr,_,_
//}
//endregion ----------------IndexController-----------------------------------


//region -------------------LoginController----------------------------------
type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["title"] = "登录"
	c.TplName = "login.tpl"
}
//endregion ----------------LoginController-----------------------------------

func haveRight(username string,pwd string) bool{
	return true
}


var (
	data = []string{"one","two","three","four"}
)
var vname1 int

func index(){
	var err error
	if err != nil {

	}
	var data []string
	data = []string{"one","two","three","four"}
	x:=3
	if data != nil {

	}
	if x==5 {

	}
}

