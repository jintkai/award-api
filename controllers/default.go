package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"io/ioutil"
)

type LuckController struct {
	beego.Controller
}

func (c *LuckController) Get() {
	logs.SetLogger("console")
	awardTypes := beego.AppConfig.Strings("awardTypes")
	awardNames := beego.AppConfig.Strings("awardNames")
	c.Data["types"] = awardTypes
	c.Data["Names"] = awardNames
	c.TplName = "index.html"
}

type Award struct {
	AwardType   string
	AwardNumber string
}

func (c *LuckController) GetList() {

	awardTypes := beego.AppConfig.Strings("awardtypes")
	//awardNames := beego.AppConfig.Strings("awardNames")
	var t map[string]map[string][]string
	t = make(map[string]map[string][]string)
	t["index"] =  map[string][]string{"key":beego.AppConfig.Strings("awardtypes"),"name":beego.AppConfig.Strings("awardNames")}
	var conf , list []string
	for _,v := range awardTypes{
		//var temp map[string][]string
		//temp = make(map[string][]string)
		var listtemp []string
		for _,v := range beego.AppConfig.Strings(v){
			conf = append(conf,v)

		}
		for _,l := range beego.AppConfig.Strings(v+"list"){
			listtemp = append(listtemp,l)
		}
		list = append(list,strings.Join(listtemp,","))
		//conf = append(conf,strings.Join(conftemp,","))
	}
	t["index"] =  map[string][]string{"key":beego.AppConfig.Strings("awardtypes"),"name":beego.AppConfig.Strings("awardNames"),"conf":conf,"list":list}

	c.Data["json"] = t
	c.ServeJSON()
}

func (c *LuckController) RandUser() {

	f,err :=ioutil.ReadFile("./conf/user.conf")
	if err != nil {
		logs.Error("读取文件失败！")
	}
	userlist := strings.Split(string(f),";")
	logs.Info("总人数：",len(userlist))
	id := c.GetString("id")
	awardConfig := beego.AppConfig.String(id)
	configArray := strings.Split(awardConfig,",")
	total,_:= strconv.Atoi(configArray[0])
	per,_ := strconv.Atoi(configArray[1])


	l :=len(userlist)
	//获取当前等级的列表
	awardList := beego.AppConfig.Strings(id+"list")
	//获取其他等级的列表
	awardTypes := beego.AppConfig.Strings("awardTypes")
	var result []string
	for i := 0; i < per; {
		logs.Debug(total,len(awardList))
		if total <= len(awardList){
			logs.Error("无可用名额")
			beego.AppConfig.SaveConfigFile("./conf/app.conf")
			break
		}else{
			rand.Seed(time.Now().UTC().UnixNano())
			r := rand.Intn(l)
			logs.Info("随机数：",r,l,"用户信息：",userlist[r])
				isExist :=false
				LABEL:
					for _,v := range awardTypes{
						temp := beego.AppConfig.Strings(v+"list")
						for _,v1 := range temp{
							if v1 !=userlist[r]  {
								logs.Debug(userlist[r],v,temp)
								isExist = false
								continue
							}else {
								logs.Info("用户在列表中，跳过",v,":",userlist[r])
								isExist = true
								break LABEL
							}
						}
					}

				if !isExist {
					logs.Error("ok",userlist[r])
					awardList = append(awardList, userlist[r])
					result = append(result, userlist[r])
					beego.AppConfig.Set(id+"list",strings.Join(awardList,";"))
					i = i+1
				}else{
					logs.Error("重复数据",userlist[r])
				}

			beego.AppConfig.Set(id+"list",strings.Join(awardList,";"))
		}
	}
	beego.AppConfig.SaveConfigFile("./conf/app.conf")
	success := "true"
	if len(result) == 0 {
		success = "false"
	}
	c.Data["json"] = map[string]string{"award":strings.Join(result,","),"success":success}
	c.ServeJSON()

}
