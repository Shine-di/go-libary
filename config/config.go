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

func Init(data interface{}) {
	log.Info("启动环境", zap.String("ENVIRON", GetEnv()))

	log.Info("启动位置", zap.String("LOCATION", GetLocation()))
	name := ""
	if IsLocal() {
		name = "config/" + GetEnv() + "-" + "local" + ".yaml"
	} else {
		name = "config/" + GetEnv() + ".yaml"
	}
	if err := util.ParseYaml(name, data); err != nil {
		panic(fmt.Sprintf("解析配置文件错误: %v", err.Error()))
	}
}
