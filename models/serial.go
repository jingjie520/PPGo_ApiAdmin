/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-24 11:48:17
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
)

type Serial struct {
	Id           int
	HardwareCode string
	SerialCode   string
	ValidTime    int64
	Remark       string
	CreateId     int
	UpdateId     int
	CreateTime   int64
	UpdateTime   int64
}

type SerialData struct {
	ValidTime    string `json:"validTime"`
	HardwareCode string `json:"serial"`
}

func (a *Serial) TableName() string {
	return TableName("set_serial")
}

func SerialGetById(id int) (*Serial, error) {
	r := new(Serial)
	err := orm.NewOrm().QueryTable(TableName("set_serial")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Serial) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
