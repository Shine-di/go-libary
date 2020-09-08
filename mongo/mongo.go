package mongo

import (
	"fmt"
	"github.com/Shine-di/go-libary/log"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"time"
)

var session *mgo.Session

var (
	Nil = mgo.ErrNotFound
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Open     int64  `json:"open"`
	Idle     int64  `json:"idle"`
}

func Conn() *mgo.Session {
	if session == nil {
		panic("获取mongo连接为空")
	}
	if err := session.Clone().Ping(); err != nil {
		panic(fmt.Sprintf("获取mongo连接 Ping出错 - [%s]", err.Error()))
	}
	return session
}

func InitConnect(config *Config) {
	dialInfo := mgo.DialInfo{
		Addrs:    []string{config.Host},
		Timeout:  time.Second * 3,
		Database: config.Database,
		Username: config.User,
		Password: config.Password,
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
	log.Info("load mongo success", zap.String("conn", config.Host))
}

type MongoLog struct {
}

func (MongoLog) Output(calldepth int, s string) error {
	log.Info(fmt.Sprintf(" %v , %v", calldepth, s))
	return nil
}
