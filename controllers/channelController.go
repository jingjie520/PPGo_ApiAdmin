package controllers

import (
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

//存储资源
func (self *ChannelController) AjaxSave() {
	id := self.GetString("id", "")
	if id == "" {
		Api := new(models.ChannelEntity)

		Api.Name = strings.TrimSpace(self.GetString("name"))

		// 检查登录名是否已经存在
		_, err := models.ChannelGetByName(Api.Name)

		if err == nil {
			self.ajaxMsg("频道名称已经存在", MSG_ERR)
		}

		if err := models.ChannelAdd(Api); err != nil {
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
