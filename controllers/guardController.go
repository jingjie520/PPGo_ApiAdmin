package controllers

import (
	"streamConsole/service"
)

type GuardController struct {
	BaseController
}

func (self *GuardController) List() {
	self.Data["pageTitle"] = "Guard Aide"
	self.display()
}

func (self *GuardController) Restart() {
	self.doExec("restart")
}

func (self *GuardController) Shutdown() {
	self.doExec("shutdown")
}

func (self *GuardController) Restart_Service() {
	self.doExec("restartService")
}

func (self *GuardController) doExec(cmd string) {
	if echo, err := service.DoExec(cmd); err == nil {
		self.ajaxMsg(echo, MSG_OK)
	} else {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
}
