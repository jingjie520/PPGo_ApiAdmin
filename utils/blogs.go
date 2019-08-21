package utils

import (
	"github.com/astaxie/beego/logs"
)

var ConsoleLogs *logs.BeeLogger

func init() {
	ConsoleLogs = logs.NewLogger()
	ConsoleLogs.SetLogger(logs.AdapterConsole)
	ConsoleLogs.SetLogger(logs.AdapterFile, `{"filename":"logs/log3-6.log","level":6,"maxlines":1000,"maxsize":1000,"daily":true,"maxdays":10,"color":true}`)

	ConsoleLogs.EnableFuncCallDepth(true)
	ConsoleLogs.SetLogFuncCallDepth(2)
	ConsoleLogs.Info("AdapterConsole Log Init Done.")

}
