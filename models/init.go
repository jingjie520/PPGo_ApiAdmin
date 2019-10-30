/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 00:18:02
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-16 17:26:48
***********************************************/

package models

import (
	"net/url"
	"streamConsole/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
)

// SerialValid 序列号是否有效

var SerialValid bool = false

func Init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	// fmt.Println(dsn)

	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(Auth), new(Role), new(RoleAuth), new(Admin), new(Serial))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	InitIptv()

	InitMongo()

	GetMgo()

}

var IptvUrl string = "" //iptv地址
func InitIptv() {
	iptvIp := beego.AppConfig.String("iptv.ip")
	iptvPort := beego.AppConfig.String("iptv.port")
	IptvUrl = iptvIp + ":" + iptvPort
}

var (
	MongodbAddr   string = "" //mongodb数据库地址
	MongodbName   string = "" //mongodb数据名称
	MongodbUser   string = "" //mongodb用户名
	MongodbPasswd string = "" //mongodb密码

	mgoSession  *mgo.Session
	mgoDatabase *mgo.Database
)

func InitMongo() {
	MongodbAddr = beego.AppConfig.String("mongo.host")
	MongodbName = beego.AppConfig.String("mongo.name")
	MongodbUser = beego.AppConfig.String("mongo.user")
	MongodbPasswd = beego.AppConfig.String("mongo.pass")

}

func GetMgo() *mgo.Session {

	utils.ConsoleLogs.Info(MongodbAddr)
	utils.ConsoleLogs.Info(MongodbName)
	utils.ConsoleLogs.Info(MongodbUser)
	utils.ConsoleLogs.Info(MongodbPasswd)

	if mgoSession == nil {
		var err error

		if MongodbUser == "" || MongodbPasswd == "" {
			mgoSession, err = mgo.Dial(MongodbAddr)
		} else {
			dialInfo := &mgo.DialInfo{
				Addrs:     []string{MongodbAddr},
				Direct:    false,
				Timeout:   time.Second * 30,
				Database:  MongodbName,
				Source:    MongodbName,
				Username:  MongodbUser,
				Password:  MongodbPasswd,
				PoolLimit: 4096, // Session.SetPoolLimit
			}

			mgoSession, err = mgo.DialWithInfo(dialInfo)
		}

		if err != nil {
			return nil
		}
		//使用指定数据库
		mgoDatabase = mgoSession.DB(MongodbName)

	}

	return mgoSession.Clone()
}
func GetDataBase() *mgo.Database {
	return mgoDatabase
}

func GetErrNotFound() error {
	return mgo.ErrNotFound
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
