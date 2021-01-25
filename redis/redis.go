package redis

import (
	"errors"
	"fmt"
	"github.com/dishine/libary/log"
	"github.com/dishine/libary/node"
	"go.uber.org/zap"
	"time"

	"github.com/go-redis/redis"
)

type Config struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Database int    `json:"database" yaml:"database"`
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
func (m *Load) Connect() error {
	config := m.config
	client := redis.NewClient(&redis.Options{
		Addr:        config.Host + ":" + config.Port,
		Password:    config.Password,
		DB:          config.Database,
		DialTimeout: time.Second * 3,
		PoolSize:    10,
	})

	redisc = &Redis{
		Client: client,
	}

	if err := client.Ping().Err(); err != nil {
		return err
	}
	log.Info("load redis success", zap.String("conn", config.Host+":"+config.Port))
	return nil
}

var redisc *Redis

func GetRedis() *Redis {
	return redisc
}

type Redis struct {
	Client *redis.Client
}

func InitRedis(config *Config) {

	client := redis.NewClient(&redis.Options{
		Addr:        config.Host,
		Password:    "",
		DB:          config.Database,
		DialTimeout: time.Second * 3,
		PoolSize:    10,
	})

	redisc = &Redis{
		Client: client,
	}

	if err := client.Ping().Err(); err != nil {
		panic(err.Error())
	}
	log.Info("load redis success", zap.String("conn", config.Host))
}

// Get returns the value saved under a given key.
func (r *Redis) GetValue(prefix PrefixEnum, key string) (string, error) {
	k := getKey(prefix, key)
	return r.Client.Get(k).Result()
}

func (r *Redis) GetValueBytes(prefix PrefixEnum, key string) ([]byte, error) {
	k := getKey(prefix, key)
	return r.Client.Get(k).Bytes()
}

func (r *Redis) MGetValue(prefix PrefixEnum, keys []string) ([]interface{}, error) {
	ks := make([]string, 0, len(keys))
	for _, value := range keys {
		ks = append(ks, getKey(prefix, value))
	}
	return r.Client.MGet(ks...).Result()
}

// Set saves an arbitrary value under a specific key.
func (r *Redis) SetValue(prefix PrefixEnum, key string, value interface{}, expire time.Duration) error {
	k := getKey(prefix, key)
	return r.Client.Set(k, value, expire).Err()
}

func (r *Redis) SetValueNX(prefix PrefixEnum, key string, value interface{}, expire time.Duration) error {
	k := getKey(prefix, key)

	bool, err := r.Client.SetNX(k, value, expire).Result()

	if err != nil {
		return err
	}
	if bool == false {
		return errors.New("setnx false")
	}
	return nil
}

func (r *Redis) Incr(prefix PrefixEnum, key string) (int64, error) {
	k := getKey(prefix, key)
	return r.Client.Incr(k).Result()
}

func (r *Redis) IncrWithExpire(prefix PrefixEnum, key string, expire time.Duration) error {
	k := getKey(prefix, key)
	err := r.Client.Incr(k).Err()
	if err != nil && expire > 0 {
		r.Client.Expire(k, expire)
	}
	return err
}

// Delete removes a specific key and its value from the Redis server.
func (r *Redis) Delete(prefix PrefixEnum, key string) error {
	k := getKey(prefix, key)
	return r.Client.Del(k).Err()
}

func (r *Redis) HExists(prefix PrefixEnum, key, field string) (bool, error) {
	k := getKey(prefix, key)
	return r.Client.HExists(k, field).Result()
}

func (r *Redis) HGet(prefix PrefixEnum, key, field string) (string, error) {
	k := getKey(prefix, key)
	return r.Client.HGet(k, field).Result()
}

func (r *Redis) HGetALLMap(prefix PrefixEnum, key string) (map[string]string, error) {
	k := getKey(prefix, key)
	return r.Client.HGetAll(k).Result()
}

func (r *Redis) HSet(prefix PrefixEnum, key, field, value string, expire time.Duration) error {
	k := getKey(prefix, key)
	err := r.Client.HSet(k, field, value).Err()
	if err != nil && expire > 0 {
		r.Client.Expire(k, expire)
	}
	return err
}

func (r *Redis) HMSetMap(prefix PrefixEnum, key string, data map[string]interface{}, expire time.Duration) error {
	k := getKey(prefix, key)
	err := r.Client.HMSet(k, data).Err()
	if err != nil && expire > 0 {
		r.Client.Expire(k, expire)
	}
	return err
}

func (r *Redis) HIncrBy(prefix PrefixEnum, key, field string, increase int64) (int64, error) {
	k := getKey(prefix, key)
	return r.Client.HIncrBy(k, field, increase).Result()
}

func (r *Redis) HDel(prefix PrefixEnum, key string, field ...string) error {
	k := getKey(prefix, key)
	return r.Client.HDel(k, field...).Err()
}

func (r *Redis) ZAdd(prefix PrefixEnum, key string, score float64, data interface{}) error {
	k := getKey(prefix, key)
	return r.Client.ZAdd(k, redis.Z{Score: score, Member: data}).Err()
}

func (r *Redis) ZMAdd(prefix PrefixEnum, key string, members ...redis.Z) error {
	k := getKey(prefix, key)
	return r.Client.ZAdd(k, members...).Err()
}

func (r *Redis) ZRangeByScore(prefix PrefixEnum, key string, offset, limit int64) ([]string, error) {
	k := getKey(prefix, key)
	return r.Client.ZRangeByScore(k, redis.ZRangeBy{Min: "-inf", Max: "+inf", Offset: offset, Count: limit}).Result()
}

func (r *Redis) ZRange(prefix PrefixEnum, key string, start, stop int64) ([]string, error) {
	k := getKey(prefix, key)
	return r.Client.ZRange(k, start, stop).Result()
}

func (r *Redis) ZCount(prefix PrefixEnum, key, min, max string) (int64, error) {
	k := getKey(prefix, key)
	return r.Client.ZCount(k, min, max).Result()
}

func (r *Redis) ZCountAll(prefix PrefixEnum, key string) (int64, error) {
	k := getKey(prefix, key)
	return r.Client.ZCount(k, "-inf", "+inf").Result()
}

func (r *Redis) ZRem(prefix PrefixEnum, key string, members ...interface{}) error {
	k := getKey(prefix, key)
	return r.Client.ZRem(k, members...).Err()
}

func (r *Redis) SAdd(prefix PrefixEnum, key string, members ...interface{}) error {
	k := getKey(prefix, key)
	return r.Client.SAdd(k, members...).Err()
}

func (r *Redis) Smembers(prefix PrefixEnum, key string) ([]string, error) {
	k := getKey(prefix, key)
	return r.Client.SMembers(k).Result()
}

func (r *Redis) SRem(prefix PrefixEnum, key, elementKey string) (int64, error) {
	k := getKey(prefix, key)
	return r.Client.SRem(k, elementKey).Result()
}

func (r *Redis) RPush(prefix PrefixEnum, key string, values ...interface{}) error {
	k := getKey(prefix, key)
	return r.Client.RPush(k, values...).Err()
}
func (r *Redis) Subscribe(prefix PrefixEnum, key string) *redis.PubSub {
	k := getKey(prefix, key)
	return r.Client.Subscribe(k)
}

func (r *Redis) Publish(prefix PrefixEnum, key string, value interface{}) (int64, error) {
	k := getKey(prefix, key)

	return r.Client.Publish(k, value).Result()
}

func (r *Redis) TTL(prefix PrefixEnum, key string) (time.Duration, error) {
	k := getKey(prefix, key)
	return r.Client.TTL(k).Result()
}

func getKey(prefix PrefixEnum, key string) string {
	return fmt.Sprintf("%s:%s", prefix, key)
}

type PrefixEnum string

const (
	RedisNil = redis.Nil
)
