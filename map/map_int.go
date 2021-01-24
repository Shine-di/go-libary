/**
 * @author: D-S
 * @date: 2020/8/25 11:13 下午
 */

package _map

import (
	"encoding/json"
	"github.com/dishine/libary/log"
	"github.com/dishine/libary/redis"
	"github.com/dishine/libary/util"
	"go.uber.org/zap"
	"sync"
)

const (
	KEY = "ENUM"
)

type IntMap struct {
	Map      map[int64]bool
	lock     sync.RWMutex
	RedisKey redis.PrefixEnum
}

func (m *IntMap) Set(key int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.Map[key] = true
}

func (m *IntMap) Get(key int64) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.Map[key]
}

func (m *IntMap) GetAll() []int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	result := make([]int64, 0)
	for k := range m.Map {
		result = append(result, k)
	}
	return result
}

func (m *IntMap) Delete(key int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.Map, key)
}

func (m *IntMap) GetAllFromRedis() []int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	redisData := make([]int64, 0)
	data, err := redis.GetRedis().GetValue(m.RedisKey, KEY)
	if err != nil {
		log.Warn("redis获取数据错误", zap.Error(err))
	} else {
		if err := json.Unmarshal([]byte(data), &redisData); err != nil {
			log.Warn("redis获取解析数据错误", zap.Error(err))
		} else {
			for _, e := range redisData {
				m.Map[e] = true
			}
		}
	}
	result := make([]int64, 0)
	for k := range m.Map {
		result = append(result, k)
	}
	return result
}

func (m *IntMap) SaveToRedis() {
	m.lock.Lock()
	defer m.lock.Unlock()
	result := make([]int64, 0)
	for k := range m.Map {
		result = append(result, k)
	}
	b, _ := json.Marshal(util.Deduplication(result))
	if err := redis.GetRedis().SetValue(m.RedisKey, KEY, string(b), 0); err != nil {
		log.Warn("数据保存redis失败", zap.Error(err))
		return
	}
	log.Info("数据保存redis成功")
}

func (m *IntMap) SYNC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	data, err := redis.GetRedis().GetValue(m.RedisKey, KEY)
	if err != nil {
		log.Warn("redis获取数据错误", zap.Error(err))
		return
	}
	result := make([]int64, 0)
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		log.Warn("redis获取解析数据错误", zap.Error(err))
	} else {
		for _, e := range result {
			m.Map[e] = true
		}
	}
}

type IntMaps struct {
	Map      map[int64][]int64
	lock     sync.RWMutex
	RedisKey redis.PrefixEnum
}

func (m *IntMaps) SetOne(key, value int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.Map[key] = append(m.Map[key], value)
}

func (m *IntMaps) GetOne(key int64) []int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.Map[key]
}

func (m *IntMaps) GetAll() map[int64][]int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.Map
}

func (m *IntMaps) Delete(key int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.Map, key)
}

func (m *IntMaps) SaveToRedis() {
	m.lock.Lock()
	defer m.lock.Unlock()
	result := make(map[int64][]int64, 0)
	for key, value := range m.Map {
		data := util.Deduplication(value)
		result[key] = append(result[key], data...)
	}
	b, _ := json.Marshal(result)
	if err := redis.GetRedis().SetValue(m.RedisKey, KEY, string(b), 0); err != nil {
		log.Warn("数据保存redis失败", zap.Error(err))
		return
	}
	log.Info("数据保存redis成功")
}

func (m *IntMaps) GetAllFromRedis() map[int64][]int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	redisData := make(map[int64][]int64, 0)
	data, err := redis.GetRedis().GetValue(m.RedisKey, KEY)
	if err != nil {
		log.Warn("redis获取数据错误", zap.Error(err))
	} else {
		if err := json.Unmarshal([]byte(data), &redisData); err != nil {
			log.Warn("redis获取解析数据错误", zap.Error(err))
		} else {
			for k, v := range redisData {
				m.Map[k] = v
			}
		}
	}
	result := make(map[int64][]int64, 0)
	for k, v := range m.Map {
		result[k] = v
	}
	return result
}

func (m *IntMaps) SYNC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	data, err := redis.GetRedis().GetValue(m.RedisKey, KEY)
	if err != nil {
		log.Warn("redis获取数据错误", zap.Error(err))
		return
	}
	result := make(map[int64][]int64, 0)
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		log.Warn("redis获取解析数据错误", zap.Error(err))
	} else {
		for k, v := range result {
			m.Map[k] = v
		}
	}
}
