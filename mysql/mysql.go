package mysql

import (
	"fmt"
	"github.com/dishine/libary/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"sync"
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

var (
	conn       *gorm.DB
	lock       = &sync.Mutex{}
	onceLoadDb sync.Once
)

func Conn() *gorm.DB {
	if conn == nil {
		panic("获取mysql连接为空")
	}

	if err := conn.DB().Ping(); err != nil {
		panic(fmt.Sprintf("获取mysql连接 Ping出错 - [%s]", err.Error()))
	}
	//conn.LogMode(true)
	return conn
}
func InitConnect(config *Config) {
	if conn != nil {
		return
	}

	host := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", config.User, config.Password, config.Host, config.Database)
	if db, err := gorm.Open("mysql", host); err != nil {
		panic(fmt.Sprintf("连接mysql出错 - [%s]", err.Error()))
		return
	} else {
		log.Info("load mysql success", zap.String("conn", host))
		db.DB().SetMaxIdleConns(5)
		db.DB().SetMaxOpenConns(20)
		db.SingularTable(true)
		conn = db
	}
	if err := conn.DB().Ping(); err != nil {
		panic(fmt.Sprintf("连接mysql Ping出错 - [%s]", err.Error()))
	}
}
