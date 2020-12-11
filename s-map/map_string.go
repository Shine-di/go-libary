/**
 * @author: D-S
 * @date: 2020/11/29 下午9:39
 */

package s_map

import (
	"encoding/json"
	"github.com/Shine-di/go-libary/log"
	"github.com/Shine-di/go-libary/redis"
	"github.com/Shine-di/go-libary/util"
	"go.uber.org/zap"
	"sync"
)

type StrMap struct {
	Map      map[string]bool
	lock     sync.RWMutex
	RedisKey redis.PrefixEnum
}

func (m *StrMap) Set(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.Map[key] = true
}

func (m *StrMap) Get(key string) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.Map[key]
}

func (m *StrMap) GetAll() []string {
	m.lock.Lock()
	defer m.lock.Unlock()
	result := make([]string, 0)
	for k := range m.Map {
		result = append(result, k)
	}
	return result
}

func (m *StrMap) Delete(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.Map, key)
}

func (m *StrMap) SaveToRedis() {
	m.lock.Lock()
	defer m.lock.Unlock()
	result := make([]string, 0)
	for k := range m.Map {
		result = append(result, k)
	}
	b, _ := json.Marshal(util.DeduplicationStr(result))
	if err := redis.GetRedis().SetValue(m.RedisKey, KEY, string(b), 0); err != nil {
		log.Warn("数据保存redis失败", zap.Error(err))
		return
	}
	log.Info("数据保存redis成功")
}

func (m *StrMap) GetAllFromRedis() []string {
	m.lock.Lock()
	defer m.lock.Unlock()
	redisData := make([]string, 0)
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
	result := make([]string, 0)
	for e := range m.Map {
		result = append(result, e)
	}
	return result
}

func (m *StrMap) SYNC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	result := make([]string, 0)
	data, err := redis.GetRedis().GetValue(m.RedisKey, KEY)
	if err != nil {
		log.Warn("redis获取数据错误", zap.Error(err))
		return
	}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		log.Warn("redis获取解析数据错误", zap.Error(err))
	}
	for _, e := range result {
		m.Map[e] = true
	}
}

type StrMaps struct {
	Map      map[string][]string
	lock     sync.RWMutex
	RedisKey redis.PrefixEnum
}

func (m *StrMaps) SetOne(key, value string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.Map[key] = append(m.Map[key], value)
}

func (m *StrMaps) GetOne(key string) []string {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.Map[key]
}

func (m *StrMaps) GetAll() map[string][]string {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.Map
}

func (m *StrMaps) Delete(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.Map, key)
}

func (m *StrMaps) SaveToRedis() {
	m.lock.Lock()
	defer m.lock.Unlock()

	result := make(map[string][]string, 0)
	for key, value := range m.Map {
		data := util.DeduplicationStr(value)
		result[key] = append(result[key], data...)
	}
	b, _ := json.Marshal(result)
	if err := redis.GetRedis().SetValue(m.RedisKey, KEY, string(b), 0); err != nil {
		log.Warn("数据保存redis失败", zap.Error(err))
		return
	}
	log.Info("数据保存redis成功")
}

func (m *StrMaps) GetAllFromRedis() map[string][]string {
	m.lock.Lock()
	defer m.lock.Unlock()
	redisData := make(map[string][]string, 0)
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
	result := make(map[string][]string, 0)
	for k, v := range m.Map {
		result[k] = v
	}
	return result
}

func (m *StrMaps) SYNC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	result := make(map[string][]string, 0)
	data, err := redis.GetRedis().GetValue(m.RedisKey, KEY)
	if err != nil {
		log.Warn("redis获取数据错误", zap.Error(err))
		return
	}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		log.Warn("redis获取解析数据错误", zap.Error(err))
	} else {
		for k, v := range result {
			m.Map[k] = v
		}
	}
}
