package main

import (
	"github.com/astaxie/beego"
	cache "github.com/patrickmn/go-cache"
	"streamConsole/models"
	_ "streamConsole/routers"
	"streamConsole/utils"
	"time"
)

func main() {
	models.Init()
	utils.Che = cache.New(60*time.Minute, 120*time.Minute)
	beego.Run()
}
