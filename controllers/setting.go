package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

type SettingController struct {
	beego.Controller
}

func (c *SettingController) GetConfig() {
	logs.SetLogger("console")

	id :=c.GetString("id")
	temp := beego.AppConfig.Strings(id)
	logs.Info(temp)
	c.Data["json"] = temp
	c.ServeJSON()
}

func (c *SettingController) SettingConf() {
	logs.SetLogger("console")
	awardTypes := beego.AppConfig.Strings("awardTypes")
	awardNames := beego.AppConfig.Strings("awardNames")
	for _,value := range awardTypes  {
		c.Data[value] = beego.AppConfig.Strings(value)
	}
	c.Data["types"] = awardTypes
	c.Data["Names"] = awardNames
	//beego.AppConfig.SaveConfigFile("./conf/app.conf")
	c.TplName = "setting.html"

}
func (c *SettingController) UpdateConfig() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.Info("更新配置文件-----------")
	id := c.GetString("id")
	tempTotal := c.GetString(id+"_total")
	tempPer := c.GetString(id+"_per")
	beego.AppConfig.Set(id,tempTotal+","+tempPer)
	beego.AppConfig.SaveConfigFile("./conf/app.conf")
	c.SettingConf()
}


