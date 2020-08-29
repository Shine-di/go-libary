package mongo

import (
	"fmt"
	"go-libary/log"
	"go.uber.org/zap"
	"time"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func Conn() *mgo.Session {
	if session == nil {
		panic("获取mongo连接为空")
	}
	if err := session.Clone().Ping(); err != nil {
		panic(fmt.Sprintf("获取mongo连接 Ping出错 - [%s]", err.Error()))
	}
	return session
}

func InitConnect() {
	dialInfo := mgo.DialInfo{
		Addrs:    []string{config.MONGO_ADDRESS},
		Timeout:  time.Second * 3,
		Database: config.MONGO_DATABASE,
		Username: config.MONGO_USER,
		Password: config.MONGO_PASSWORD,
	}
	ses, err := mgo.DialWithInfo(&dialInfo)
	if err != nil {
		log.Info("mongo 连接错误")
		panic(err.Error())
	}
	ses.SetSocketTimeout(time.Hour)
	//
	////if c.IsDevelop() {
	//	mgo.SetDebug(true)
	//	mgo.SetLogger(new(MongoLog))
	////}
	session = ses
	log.Info("load mongo success", zap.String("conn", config.MONGO_ADDRESS))
}

type MongoLog struct {
}

func (MongoLog) Output(calldepth int, s string) error {
	log.Info(fmt.Sprintf(" %v , %v", calldepth, s))
	return nil
}
