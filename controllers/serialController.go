package controllers

import (
	"github.com/astaxie/beego"
	"streamConsole/libs"
	"streamConsole/models"
	"time"
)

type SerialController struct {
	BaseController
}

//查看详情
func (self *SerialController) Detail() {

	detail, _ := models.SerialGetById(1)
	row := make(map[string]interface{})
	row["id"] = detail.Id
	row["hardware_code"] = detail.HardwareCode
	row["serial_code"] = detail.SerialCode
	row["remark"] = detail.Remark
	row["valid_time"] = beego.Date(time.Unix(detail.ValidTime, 0), "Y-m-d H:i:s")

	self.Data["pageTitle"] = "序列号"
	self.Data["Detail"] = row
	self.display()
}

func (self *SerialController) AjaxSave() {
	serialCode := self.GetString("serialCode")
	serial := libs.ManualCheckSerial(serialCode)
	if models.SerialValid {
		self.ajaxMsg("", MSG_OK)
	} else {
		self.ajaxMsg(serial.Remark, MSG_ERR)
	}

}
