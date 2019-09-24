package controllers

import (
	"fmt"
	"net/url"
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
	row["id"] = v.ID.Hex()
	row["channelID"] = v.ChannelID
	row["name"] = v.Name

	if v.NetCardin == "" {
		row["mode"] = "single"
	} else {
		row["mode"] = "group"
	}
	row["src"] = v.Src
	row["program"] = v.Program
	row["netcardin"] = v.NetCardin
	row["single"] = v.Single
	row["group"] = v.Group
	row["groupUrl"] = v.Groupurl
	row["netcard"] = v.NetCard
	row["vod"] = v.Vod

	row["tsoc"] = v.TSoc
	row["toaac"] = v.ToAac
	row["curGroup"] = v.CurGroup
	row["curSingle"] = v.CurSingle
	row["curVod"] = v.CurVod

	row["toaac"] = v.ToAac

	self.Data["Source"] = row
	self.Data["netcards"] = libs.GetNetCards()
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
	row["id"] = v.ID.Hex()
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
	channel := new(models.ChannelEntity)

	channel.ChannelID = strings.TrimSpace(self.GetString("channelID"))
	channel.Name = strings.TrimSpace(self.GetString("name"))
	channel.Src = strings.TrimSpace(self.GetString("src"))
	channel.Vod = strings.TrimSpace(self.GetString("vod"))
	channel.Single = strings.TrimSpace(self.GetString("single"))
	channel.Group = strings.TrimSpace(self.GetString("group"))
	channel.ToAac = strings.TrimSpace(self.GetString("toaac"))
	channel.TSoc = strings.TrimSpace(self.GetString("tsoc"))
	param := ""

	if id == "" {
		param = "type=add"
		// 检查登录名是否已经存在
		_, err := models.ChannelGetByName(channel.Name)
		if err == nil {
			self.ajaxMsg("频道名称已经存在", MSG_ERR)
		}

	} else {
		param = "type=update"
		channelTemp, err := models.ChannelGetByName(channel.Name)
		if err == nil && channelTemp.ID.Hex() != id {
			self.ajaxMsg("频道名称已经存在", MSG_ERR)
		}
	}

	param = "type=add&id=" + channel.ChannelID
	param += "&name=" + url.QueryEscape(channel.Name)
	param += "&src=" + url.QueryEscape(channel.Src)
	param += "&mode=" + strings.TrimSpace(self.GetString("mode"))
	param += "&vod=" + channel.Vod
	param += "&single=" + channel.Single
	param += "&group=" + channel.Group
	param += "&toaac=" + channel.ToAac

	if strings.TrimSpace(self.GetString("mode")) == "group" {
		param += "&netcardin=" + strings.TrimSpace(self.GetString("netcardin"))
		param += "&program=" + strings.TrimSpace(self.GetString("program"))
	}

	if channel.Vod == "on" {
		param += "&vodTime=36000" //+ strings.TrimSpace(self.GetString("vodTime"))
		param += "&vodpath=" + strings.TrimSpace(self.GetString("vodpath"))
	}

	if channel.Group == "on" {
		param += "&netcard=" + strings.TrimSpace(self.GetString("netcard"))
		param += "&groupurl=" + strings.TrimSpace(self.GetString("groupurl"))
	}

	if channel.TSoc == "on" {
		param += "&tsoc=" + channel.TSoc
		param += "&tsoctime=3600" //+ strings.TrimSpace(self.GetString("tsoctime"))
	}

	//请求api
	_, err := libs.SaveChannelEntity(channel, param)

	// err := models.ChannelAdd(channel);

	if err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *ChannelController) AjaxDel() {
	id := self.GetString("id", "")
	ApiUpdate, _ := models.ChannelGetById(id)
	if ApiUpdate != nil {
		if _, err := libs.DeleteChannel(ApiUpdate); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	}
	self.ajaxMsg("", MSG_OK)
}
