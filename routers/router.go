package routers

import (
	"github.com/astaxie/beego"
	"goweb/controllers"
)

func init() {
	beego.Router("/", &controllers.LuckController{}, "get:Get")
	beego.Router("/getList", &controllers.LuckController{}, "get:GetList")
	beego.Router("/randUser", &controllers.LuckController{}, "get:RandUser")
	beego.Router("/luck", &controllers.LuckController{}, "post:RandUser")
	beego.Router("/setting", &controllers.SettingController{}, "get:SettingConf;post:UpdateConfig")
	beego.Router("/getConfig", &controllers.SettingController{}, "get:GetConfig")
	beego.Router("/updateConfig", &controllers.SettingController{}, "post:UpdateConfig")
}
