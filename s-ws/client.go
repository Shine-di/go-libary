/**
 * @author: D-S
 * @date: 2020/7/25 5:03 下午
 */

package s_ws

import (
	"fmt"
	"github.com/Shine-di/go-libary/log"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"time"
)

var (
	stop = make(chan bool, 0)
)

// stop chan 必须有缓存的
type WS struct {
	URL              string
	Stop             chan bool
	Message          chan []byte
	Duration         time.Duration
	Header           http.Header
	Con              *websocket.Conn
	HeartbeatMessage string
	Proxy            string
}

func (e *WS) Start() {
	//e.Proxy = "http://8.210.99.228:59073"
	dialer := new(websocket.Dialer)
	if e.Proxy != "" {
		proxy, err := url.Parse(e.Proxy)
		if err != nil {
			fmt.Println(err.Error())
		}
		dialer.Proxy = http.ProxyURL(proxy)
	}
	conn, _, err := dialer.Dial(e.URL, e.Header)
	if err != nil {
		log.Warn("ws error", zap.Any("连接", e.URL), zap.Error(err))
		e.Stop <- true
		return
	}
	log.Info("连接成功")
	e.Con = conn
	go e.Heartbeat(conn)
	go e.ReadMessage(conn)
	return
}
func (e *WS) Heartbeat(conn *websocket.Conn) {
	m := "ping"
	if e.HeartbeatMessage != "" {
		m = e.HeartbeatMessage
	}
	for {
		err := conn.WriteMessage(websocket.TextMessage, []byte(m))
		log.Info("--心跳--", zap.Any("连接", e.URL))
		if err != nil {
			e.Stop <- true
			return
		}
		<-time.After(e.Duration)
	}
}
func (e *WS) ReadMessage(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			e.Stop <- true
			return
		}
		e.Message <- message
	}
}
