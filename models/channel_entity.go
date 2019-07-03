package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type ChannelEntity struct {
	ID        bson.ObjectId `bson:"_id"`
	ChannelID string        `bson:"id"`
	Name      string        `bson:"name"`
	Src       string        `bson:"src"`
	NetCardin string        `bson:"netcardin"`
	Program   int           `bson:"program"`
	Group     string        `bson:"group"`
	Single    string        `bson:"single"`
	Vod       string        `bson:"vod"`
	TSoc      string        `bson:"tsoc"`
	ToAac     string        `bson:"toaac"`
	CurGroup  string        `bson:"cur_group"`
	CurSingle string        `bson:"cur_single"`
	CurVod    string        `bson:"cur_vod"`
	CurTsoc   string        `bson:"cur_tsoc"`
	DeMux     int           `bson:"demux"`
}

func (a *ChannelEntity) TableName() string {
	return "channel"
}

func ChannelGetList(page, pageSize int, filters ...interface{}) ([]*ChannelEntity, int64) {
	offset := (page - 1) * pageSize
	list := make([]*ChannelEntity, 0)

	con := GetDataBase().C("channel")
	query := con.Find(nil)

	total, _ := query.Count()

	query.Skip(offset).Limit(pageSize).All(&list)

	return list, int64(total)
}

func ChannelGetById(id string) (*ChannelEntity, error) {
	objectId := bson.ObjectIdHex(id)
	r := new(ChannelEntity)
	con := GetDataBase().C(r.TableName())
	err := con.Find(bson.M{"_id": objectId}).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func ChannelGetByName(name string) (*ChannelEntity, error) {
	r := new(ChannelEntity)
	err := GetDataBase().C(r.TableName()).Find(bson.M{"name": name}).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func checkChannelId(a *ChannelEntity) error {
	if strings.ToLower(a.ChannelID) == "all" {
		return errors.New("节目ID不是是all")
	}
	return nil
}

func ChannelAdd(a *ChannelEntity) error {
	if err := checkChannelId(a); err != nil {
		return err
	}
	return GetDataBase().C(a.TableName()).Insert(a)
}

func (a *ChannelEntity) Update(fields ...string) error {
	if err := checkChannelId(a); err != nil {
		return err
	}
	err := GetDataBase().C(a.TableName()).UpdateId(a.ID, a)
	if err != nil {
		return err
	}
	return nil
}

func (a *ChannelEntity) Delete() error {
	return GetDataBase().C(a.TableName()).RemoveId(a.ID)
}
