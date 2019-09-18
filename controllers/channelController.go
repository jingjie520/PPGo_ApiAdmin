package controllers

import (
	"fmt"
	"streamConsole/libs"
	"streamConsole/models"
	"strings"
)

type ChannelController struct {
	BaseController
}

func (self *ChannelController) List() {
	self.Data["pageTitle"] = "频道列表"
	self.display()
}

func (self *ChannelController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	//sourceName := strings.TrimSpace(self.GetString("sourceName"))
	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)

	result, count := models.ChannelGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.ID
		row["channelID"] = v.ChannelID
		row["name"] = v.Name
		row["src"] = v.Src
		row["group"] = v.Group
		row["single"] = v.Single
		row["vod"] = v.Vod
		row["tsoc"] = v.TSoc
		row["toaac"] = v.ToAac
		row["demux"] = v.DeMux
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *ChannelController) Add() {
	self.Data["pageTitle"] = "新增频道"

	//装载网卡下拉框
	self.Data["netcards"] = libs.GetNetCards()
	fmt.Println(self.Data["netcards"])

	self.display()
}

func (self *ChannelController) Edit() {
	self.Data["pageTitle"] = "编辑频道"

	id := self.GetString("id", "")
	v, err := models.ChannelGetById(id)
	if err != nil {
		self.Ctx.WriteString("频道不存在")
		return
	}
	row := make(map[string]interface{})
	row["id"] = v.ID
	row["name"] = v.Name
	row["src"] = v.Src
	row["group"] = v.Group
	row["single"] = v.Single
	row["vod"] = v.Vod
	row["tsoc"] = v.TSoc
	row["toaac"] = v.ToAac
	self.Data["Source"] = row

	self.display()
}

func (self *ChannelController) ActionStart() {
	self.Data["pageTitle"] = "启动频道"
	id := self.GetString("id", "")
	v, err := models.ChannelGetById(id)
	if err != nil {
		self.Ctx.WriteString("频道不存在")
		return
	}
	row := make(map[string]interface{})
	row["id"] = v.ID.Hex()
	row["name"] = v.Name

	row["src"] = v.Src
	row["group"] = v.Group
	row["single"] = v.Single
	row["vod"] = v.Vod
	row["tsoc"] = v.TSoc

	self.Data["Source"] = row
	self.display()
}

func (self *ChannelController) ActionStop() {
	self.Data["pageTitle"] = "停止频道"
	id := self.GetString("id", "")
	v, err := models.ChannelGetById(id)
	if err != nil {
		self.Ctx.WriteString("频道不存在")
		return
	}
	row := make(map[string]interface{})
	row["id"] = v.ID
	row["name"] = v.Name

	row["group"] = v.Group
	row["single"] = v.Single
	row["vod"] = v.Vod
	row["tsoc"] = v.TSoc

	self.Data["Source"] = row
	self.display()
}

func (self *ChannelController) AjaxStartSave() {
	id := self.GetString("id", "")
	if id != "" {
		channel, _ := models.ChannelGetById(id)
		channel.Group = strings.TrimSpace(self.GetString("group"))
		channel.Single = strings.TrimSpace(self.GetString("single"))
		channel.Vod = strings.TrimSpace(self.GetString("vod"))
		channel.TSoc = strings.TrimSpace(self.GetString("tsoc"))

		//请求api
		_, err := libs.SaveChannelStatus(channel)

		if err == nil {
			self.ajaxMsg("", MSG_OK)
		} else {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	} else {
		self.ajaxMsg("记录不存在", MSG_ERR)
	}
}

func (self *ChannelController) AjaxStopSave() {
	self.AjaxStartSave()
}

//存储资源
func (self *ChannelController) AjaxSave() {
	id := self.GetString("id", "")
	if id == "" {
		channel := new(models.ChannelEntity)

		channel.ChannelID = strings.TrimSpace(self.GetString("channelID"))
		channel.Name = strings.TrimSpace(self.GetString("name"))
		channel.Src = strings.TrimSpace(self.GetString("src"))
		channel.Program, _ = self.GetInt("program", 0)

		channel.Group = strings.TrimSpace(self.GetString("group"))
		channel.Single = strings.TrimSpace(self.GetString("single"))
		channel.Vod = strings.TrimSpace(self.GetString("vod"))
		channel.TSoc = strings.TrimSpace(self.GetString("tsoc"))

		channel.NetCardin = strings.TrimSpace(self.GetString("netcardin"))

		param := "type=add&id=" + channel.ChannelID
		param += "&netcard=" + strings.TrimSpace(self.GetString("netcard"))
		param += "&groupurl=" + strings.TrimSpace(self.GetString("groupurl"))

		if channel.TSoc == "on" {
			param += "&tsoctime=36000"
		}
		if channel.Vod == "on" {
			param += "&vodtime=36000"
		}

		// 检查登录名是否已经存在
		_, err := models.ChannelGetByName(channel.Name)

		if err == nil {
			self.ajaxMsg("频道名称已经存在", MSG_ERR)
		}

		//请求api
		_, err = libs.SaveChannelEntity(channel, param)

		// err := models.ChannelAdd(channel);

		if err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	} else {
		ApiUpdate, _ := models.ChannelGetById(id)
		// 修改
		ApiUpdate.Name = strings.TrimSpace(self.GetString("name"))

		if err := ApiUpdate.Update(); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}
}

func (self *ChannelController) AjaxDel() {
	id := self.GetString("id", "")
	ApiUpdate, _ := models.ChannelGetById(id)
	if ApiUpdate != nil {
		if err := ApiUpdate.Delete(); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	}
	self.ajaxMsg("", MSG_OK)
}
