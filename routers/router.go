package routers

import (
	"myproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index.html", &controllers.IndexController{})
	beego.Router("/login.html", &controllers.LoginController{})
	beego.Router("/savedata", &controllers.IndexController{},"post:SaveData")
	//beego.Router("/savedata", &controllers.IndexController{},"get:SaveData")

	beego.Router("/userlist", &controllers.IndexController{},"get:UserList")
    //三参可以分号分隔
	//beego.Router("/test_model", &controllers.IndexController{}, "get:Get;post:Post")
}
