/**
 * @author: D-S
 * @date: 2020/8/25 11:13 下午
 */

package s_map

import (
	"encoding/json"
	"github.com/Shine-di/go-libary/log"
	"github.com/Shine-di/go-libary/redis"
	"go.uber.org/zap"
	"sync"
)

const (
	KEY = "ENUM"
)
type IntMap struct {
	Map  map[int64]bool
	lock sync.RWMutex
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

func(m *IntMap)SaveToRedis(){
	m.lock.Lock()
	defer m.lock.Unlock()
	result := make([]int64, 0)
	for k := range m.Map {
		result = append(result, k)
	}
	 if err := redis.GetRedis().SetValue(m.RedisKey,KEY,result,0);err != nil {
	 	log.Warn("数据保存redis失败",zap.Error(err))
	 	return
	 }
	 log.Info("数据保存redis成功")
}

func(m *IntMap)GetAllFromRedis()[]int64{
	m.lock.Lock()
	defer m.lock.Unlock()
	result := make([]int64,0)
	data,err :=redis.GetRedis().GetValue(m.RedisKey,KEY)
	if err != nil {
		log.Warn("redis获取数据错误",zap.Error(err))
		return result
	}

	if err :=json.Unmarshal([]byte(data), &result);err != nil {
		log.Warn("redis获取解析数据错误",zap.Error(err))
	}
	return result
}
type IntMaps struct {
	Map  map[int64][]int64
	lock sync.RWMutex
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




