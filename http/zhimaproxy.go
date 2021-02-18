/**
 * @author: D-S
 * @date: 2021/2/18 6:58 下午
 */

package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dishine/libary/cst"
	"github.com/dishine/libary/log"
	"go.uber.org/zap"
	"sync"
	"time"
)

type zhiMaResp struct {
	Code    int64     `json:"code"`
	Data    []ProxyIp `json:"data"`
	Msg     string    `json:"msg"`
	Success bool      `json:"success"`
}
type ProxyIp struct {
	IP         string `json:"ip"`
	Port       int64  `json:"port"`
	ExpireTime string `json:"expire_time"` //2021-02-18 19:11:02
}

var (
	ZhiMa         *ZhiMaProxy
	useZhiMaProxy = true
)

func InItZhiMaProxy(addr string, useProxy ...bool) error {
	if len(useProxy) > 0 {
		useZhiMaProxy = useProxy[0]
		if useProxy[0] {
			ZhiMa = NewZhiMaProxy(addr)
		}
	} else {
		ZhiMa = NewZhiMaProxy(addr)
	}
	return nil
}

//http://webapi.http.zhimacangku.com/getip?num=10&type=2&pro=0&city=0&yys=0&port=11&time=1&ts=1&ys=0&cs=0&lb=1&sb=0&pb=4&mr=1&regions=110000,130000,140000,210000,220000,230000,310000,320000,330000,340000,350000,360000,370000,410000,420000,430000,440000,450000,500000,510000,520000,530000,610000,640000

type ZhiMaProxy struct {
	zhiMaUrl string
	lock     sync.RWMutex
	index    int
	ips      map[string]int64 // key为代理ip value 为过期时间
}

func NewZhiMaProxy(addr string) *ZhiMaProxy {
	zhiMa := &ZhiMaProxy{
		zhiMaUrl: addr,
		lock:     sync.RWMutex{},
		index:    11,
		ips:      map[string]int64{},
	}
	if useZhiMaProxy {
		_ = zhiMa.LoadIp()
	}
	return zhiMa
}

func (z *ZhiMaProxy) LoadIp() error {
	get := GET{
		URL:      z.zhiMaUrl,
		Header:   nil,
		Params:   nil,
		UseProxy: false,
		Proxy:    "",
		Token:    "",
	}
	resp, err := get.Do()
	if err != nil {
		return err
	}
	zhiMaResp := new(zhiMaResp)
	if err := json.Unmarshal(resp, zhiMaResp); err != nil {
		return err
	}
	if !zhiMaResp.Success {
		return errors.New(zhiMaResp.Msg)
	}
	z.lock.Lock()
	defer z.lock.Unlock()
	for _, e := range zhiMaResp.Data {
		t, _ := time.ParseInLocation(cst.TIME_STAMP, e.ExpireTime, time.Local)
		t.Unix()
		proxyIp := "http://" + e.IP + ":" + fmt.Sprint(e.Port)
		_, ok := z.ips[proxyIp]
		if ok {
			z.ips[proxyIp] = t.Unix()
		} else {
			z.ips[proxyIp] = t.Unix()
		}
	}
	return nil
}
func (z *ZhiMaProxy) checkIp() {
	if !useZhiMaProxy {
		return
	}
	z.lock.Lock()
	if len(z.ips) == 0 {
		z.lock.Unlock()
		if err := z.LoadIp(); err != nil {
			log.Info("获取新的ip错误", zap.Error(err))
			return
		}
	}
	defer z.lock.Unlock()
	t := time.Now().Unix()
	for k, v := range z.ips {
		if v < t {
			delete(z.ips, k)
		}
	}
}
func (z *ZhiMaProxy) GetIp() string {
	if !useZhiMaProxy {
		return ""
	}
	z.lock.Lock()
	if len(z.ips) == 0 {
		z.lock.Unlock()
		if err := z.LoadIp(); err != nil {
			log.Info("获取新的ip错误", zap.Error(err))
			return ""
		}
		return ""
	}
	defer z.lock.Unlock()
	t := time.Now().Unix()
	for k, v := range z.ips {
		if v < t {
			delete(z.ips, k)
			continue
		}
		return k
	}
	return ""
}
