package mysql

import (
	"fmt"
	"github.com/dishine/libary/log"
	"github.com/dishine/libary/node"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"sync"
)

type Config struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
	Open     int64  `json:"open" yaml:"open"`
	Idle     int64  `json:"idle" yaml:"idle"`
}

var (
	conn              *gorm.DB
	lock              = &sync.Mutex{}
	onceLoadDb        sync.Once
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

func Conn() *gorm.DB {
	if conn == nil {
		panic("获取mysql连接为空")
	}

	if err := conn.DB().Ping(); err != nil {
		panic(fmt.Sprintf("获取mysql连接 Ping出错 - [%s]", err.Error()))
	}
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
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		db.SingularTable(true)
		conn = db
	}
	if err := conn.DB().Ping(); err != nil {
		panic(fmt.Sprintf("连接mysql Ping出错 - [%s]", err.Error()))
	}
}

func NewLoad(config *Config) *Load {
	return &Load{
		config: config,
	}
}

type Load struct {
	config *Config
}

func (m *Load) GetOrder() node.Order {
	return node.Before
}
func (m *Load) GetOptionFunc() node.Func {
	return m.Connect
}

// 从yaml 中获取 配置内容
func (m *Load) Connect() error {
	config := m.config

	host := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", config.User, config.Password, config.Host, config.Port, config.Database)
	if db, err := gorm.Open("mysql", host); err != nil {
		panic(fmt.Sprintf("连接mysql出错 - [%s]", err.Error()))
		return err
	} else {
		log.Info("load mysql success", zap.String("conn", host))
		db.DB().SetMaxIdleConns(5)
		db.DB().SetMaxOpenConns(20)
		db.SingularTable(true)
		conn = db
	}
	if err := conn.DB().Ping(); err != nil {
		return err
	}
	return nil
}
