/**
 * @author: D-S
 * @date: 2021/1/11 12:40 下午
 */

package token

import (
	"fmt"
	"github.com/dishine/libary/log"
	"github.com/dishine/libary/response"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type Limit struct {
	Max  int
	m    map[string]int
	lock sync.Mutex
}

func (l *Limit) clean(t time.Duration) {
	for {
		log.Info("开始执行定时任务 - Clean")
		l.lock.Lock()
		l.m = make(map[string]int)
		l.lock.Unlock()
		<-time.After(t)
	}
}
func NewLimit(max int, t time.Duration) *Limit {
	l := &Limit{
		Max:  max,
		m:    map[string]int{},
		lock: sync.Mutex{},
	}
	go l.clean(t)
	return l
}

func (l *Limit) Check(ip string) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	c, ok := l.m[ip]
	if !ok {
		l.m[ip] = 1
	} else {
		if c >= l.Max {
			return false
		}
		l.m[ip] += 1
	}
	return true
}

func (l *Limit) RateLimitMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		if !l.Check(token) {
			c.Set(response.ERROR, response.Error("too many requests"))
			log.Info(fmt.Sprintf("TOKEN [%s] 已限流", token))
			c.Abort()
			return
		}
		c.Next()
	}
}
