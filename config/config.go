/**
 * @author: D-S
 * @date: 2020/8/25 2:22 下午
 */

package config

import (
	"fmt"
	"github.com/dishine/libary/log"
	"github.com/dishine/libary/mongo"
	"github.com/dishine/libary/mysql"
	"github.com/dishine/libary/oss"
	"github.com/dishine/libary/redis"
	"github.com/dishine/libary/util"
	"go.uber.org/zap"
	"os"
)

func SetEnv(env string) {
	_ = os.Setenv("ENVIRON", env)
}
func GetEnv() string {
	return os.Getenv("ENVIRON")
}

func SetLocation(location string) {
	_ = os.Setenv("LOCATION", location)
}
func GetLocation() string {
	return os.Getenv("LOCATION")
}

func IsLocal() bool {
	if GetLocation() == "local" {
		return true
	}
	return false
}

var (
	Config *Yaml
)

type Yaml struct {
	Mysql *mysql.Config `json:"mysql" yaml:"mysql"`
	Mongo *mongo.Config `json:"mongo" yaml:"mongo"`
	Redis *redis.Config `json:"redis" yaml:"redis"`
}

// 环境加后缀
func init() {
	log.Info("启动环境", zap.String("ENVIRON", GetEnv()))

	log.Info("启动位置", zap.String("LOCATION", GetLocation()))
	name := ""
	if IsLocal() {
		name = "config/" + GetEnv() + "-" + "local" + ".yaml"
	} else {
		name = "config/" + GetEnv() + ".yaml"
	}
	config := new(Yaml)
	if err := util.ParseYaml(name, config); err != nil {
		panic(fmt.Sprintf("解析配置文件错误: %v", err.Error()))
	}
	Config = config
}

func init() {
	switch GetEnv() {
	case RELEASE:
		REDIS_ADDRESS = HOST_RELEASE + ":63379"
		MYSQL_HOST = HOST_RELEASE + ":33336"
		MONGO_ADDRESS = HOST_RELEASE + ":27020"
		OSS_ADDRESS = "http://" + OSS_RELEASE
	case LOCAL:
		REDIS_ADDRESS = HOST_LOCAL + ":63379"
		MYSQL_HOST = HOST_LOCAL + ":33336"
		MONGO_ADDRESS = HOST_LOCAL + ":27020"
		OSS_ADDRESS = "http://" + OSS_LOCAL
	default:
		//panic("启动环境错误")

	}
}

const (
	RELEASE = "release"
	LOCAL   = "local"

	HOST_RELEASE = "172.24.172.191"
	HOST_LOCAL   = "47.112.121.179"

	OSS_RELEASE = "oss-cn-shenzhen-internal.aliyuncs.com"
	OSS_LOCAL   = "oss-cn-shenzhen.aliyuncs.com"

	AccessKeyId     = "LTAI4G7891QXgoxYevdLyVEn"
	AccessKeySecret = "e6mN4TAbzyYk6ZtqzmlZR9y8p8r6gJ"
	Bucket          = "ufinance"

	MONGO_USER     = "root"
	MONGO_DATABASE = "admin"
	MONGO_PASSWORD = "Js123!@#"

	MYSQL_USER_NAME = "root"
	MYSQL_PASSWORD  = "Js123!@#"
	MYSQL_DATABASE  = "spark"

	REDIS_PASSWORD = "Js123!@#"
)

var (
	// redis
	REDIS_ADDRESS string

	// mysql
	MYSQL_HOST string

	MONGO_ADDRESS string

	//oss
	OSS_ADDRESS string
)

func GetRedisConfig(database int) *redis.Config {
	return &redis.Config{
		Host:     REDIS_ADDRESS,
		Port:     "",
		User:     "",
		Password: REDIS_PASSWORD,
		Database: database,
	}
}

func GetMySqlConfig(database string) *mysql.Config {
	return &mysql.Config{
		Host:     MYSQL_HOST,
		Port:     "",
		User:     MYSQL_USER_NAME,
		Password: MYSQL_PASSWORD,
		Database: database,
	}
}

func GetMongoConfig() *mongo.Config {
	return &mongo.Config{
		Host:     MONGO_ADDRESS,
		Port:     "",
		User:     MONGO_USER,
		Password: MONGO_PASSWORD,
		Database: MONGO_DATABASE,
	}
}

func GetOssConfig() *oss.Config {
	return &oss.Config{
		EndPoint:        OSS_ADDRESS,
		Bucket:          Bucket,
		Timeout:         0,
		AccessKeyId:     AccessKeyId,
		SecretAccessKey: AccessKeySecret,
		BaseUrl:         OSS_LOCAL,
	}
}
