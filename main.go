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

	/*解码测试
	s := "d1d730bcf82849d53184e44fed4a2c7396dd3a60bf96b12ee7ade7a19d714ed49924e07014422dd75f2d8c91f59a5e349d7f66b9d2dde38a6269d84dc907801bdf3949512badbadfe0ab3661d3aa2b916ddcf9ba1f3b68b500fef26e4aba2525"

	json := utils.AesDecrypt(s)
	utils.ConsoleLogs.Info("JSON：%s ", json)
	*/

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
