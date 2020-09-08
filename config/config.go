/**
 * @author: D-S
 * @date: 2020/8/25 2:22 下午
 */

package config

import (
	"github.com/Shine-di/go-libary/mongo"
	"github.com/Shine-di/go-libary/mysql"
	"github.com/Shine-di/go-libary/oss"
	"github.com/Shine-di/go-libary/redis"
	"os"
)

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
)

var (
	// redis
	REDIS_ADDRESS  string
	REDIS_PASSWORD = "Js123!@#"
	// mysql
	MYSQL_HOST      string
	MYSQL_USER_NAME = "root"
	MYSQL_PASSWORD  = "Js123!@#"
	MYSQL_DATABASE  = "spark"

	MONGO_ADDRESS  string
	MONGO_USER     = "root"
	MONGO_DATABASE = "admin"
	MONGO_PASSWORD = "Js123!@#"
	//oss
	OSS_ADDRESS string
)

func GetEnv() string {
	return os.Getenv("ENVIRON")
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
		panic("启动环境错误")

	}
}

func GetRedisConfig() *redis.Config {
	return &redis.Config{
		Host:     REDIS_ADDRESS,
		Port:     "",
		User:     "",
		Password: REDIS_PASSWORD,
		Database: "",
	}
}

func GetMySqlConfig() *mysql.Config {
	return &mysql.Config{
		Host:     MYSQL_HOST,
		Port:     "",
		User:     MYSQL_USER_NAME,
		Password: MYSQL_PASSWORD,
		Database: MYSQL_DATABASE,
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
