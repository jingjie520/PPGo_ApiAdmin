package main

import (
	"github.com/astaxie/beego"
	"github.com/patrickmn/go-cache"
	"streamConsole/libs"
	"streamConsole/models"
	_ "streamConsole/routers"
	"streamConsole/utils"
	"time"
)

func main() {
	models.Init()
	utils.Che = cache.New(60*time.Minute, 120*time.Minute)

	libs.AutoCheckSerial()

	//输出
	utils.ConsoleLogs.Info("注册状态：%t ", models.SerialValid)

	/*
		s:="{ \"serial\":\"E5060100FFFBAB0F\", \"expire\":\"2099-12-31\", \"salt\":\"LICc8GAeGBpFmWNr\" }"
		sn :=  utils.AesEncrypt(s)
		utils.ConsoleLogs.Info("序列号：%s ", sn)

	*/

	beego.Run()
}
