package utils

import "github.com/astaxie/beego/logs"

var ConsoleLogs *logs.BeeLogger
var FileLogs *logs.BeeLogger

func init() {
	ConsoleLogs = logs.NewLogger(1000)
	ConsoleLogs.SetLogger("console")
	FileLogs = logs.NewLogger(1000)
	FileLogs.SetLogger("file", `{"filename":”logs/test.log"}`)
}
