/**
 * @author: D-S
 * @date: 2020/12/8 下午6:07
 */

package http

import "sync"

var (
	lock  = sync.RWMutex{}
	proxy = []string{}

	index = 9
)

func InitProxy(ips []string) {
	lock.Lock()
	proxy = ips
	lock.Unlock()
}

func getProxy() string {
	if len(proxy) == 0 {
		return ""
	}
	defer lock.Unlock()
	lock.Lock()
	k := index % len(proxy)
	index = index + 1
	return proxy[k]
}
