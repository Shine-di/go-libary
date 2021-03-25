/**
 * @author: D-S
 * @date: 2020/8/7 5:47 下午
 */

package ws

import (
	"errors"
	"fmt"
	"github.com/dishine/libary/log"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type Binder struct {
	Mux    sync.Mutex
	ConMap map[string]*websocket.Conn
}

var (
	sUpGrader = websocket.Upgrader{
		//跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type WSService struct {
	W      http.ResponseWriter
	R      *http.Request
	Ip     string // 连接Ip
	Binder *Binder
	Tag    string
}

func (wss *WSService) Start() error {
	con, err := sUpGrader.Upgrade(wss.W, wss.R, nil)
	if err != nil {
		return err
	}
	//保存连接
	if err := wss.saveCon(con); err != nil {
		return err
	}
	// 读取心跳
	wss.readMessage(con)
	return nil
}

func (wss *WSService) readMessage(con *websocket.Conn) {
	for {
		_, message, err := con.ReadMessage()
		if err != nil {
			con.Close()
			wss.Binder.Mux.Lock()
			delete(wss.Binder.ConMap, wss.Ip)
			wss.Binder.Mux.Unlock()
			log.Info(fmt.Sprintf("tag %v ip %v read message err %v ", wss.Tag, wss.Ip, err))
			return
		}
		log.Info(fmt.Sprintf("tag %v ip %v message %v ", wss.Tag, wss.Ip, string(message)))
		if err := con.WriteMessage(websocket.TextMessage, []byte(`pong`)); err != nil {
			wss.Binder.Mux.Lock()
			delete(wss.Binder.ConMap, wss.Ip)
			wss.Binder.Mux.Unlock()
			log.Info(fmt.Sprintf("ip %v send pong err %v ", wss.Ip, err))
			con.Close()
			return
		}
	}
}

func (wss *WSService) saveCon(con *websocket.Conn) error {
	if wss.Binder == nil {
		return errors.New("保存连接错误")
	}
	wss.Binder.Mux.Lock()
	defer wss.Binder.Mux.Unlock()
	if wss.Ip == "" {
		return errors.New("ip is nil")
	}
	c, ok := wss.Binder.ConMap[wss.Ip]
	if ok {
		c.Close()
	}
	log.Info(fmt.Sprintf("%v 连接 %v 保存成功", wss.Tag, wss.Ip))
	wss.Binder.ConMap[wss.Ip] = con
	return nil
}
