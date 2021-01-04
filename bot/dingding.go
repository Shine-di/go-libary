/**
 * @author: D-S
 * @date: 2021/1/4 3:33 下午
 */

package bot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/Shine-di/go-libary/log"
	"github.com/royeo/dingrobot"
	"go.uber.org/zap"
	"net/url"
	"time"
)

type DingBot struct {
	WebHook string `json:"web_hook"`
	Sign    string `json:"sign"`
}

func NewDingBot(webHook, sign string) *DingBot {
	return &DingBot{
		WebHook: webHook,
		Sign:    sign,
	}
}

func (d *DingBot) SendTest(content string) {
	time.Sleep(500 * time.Millisecond)
	ts, sign := d.sign(d.Sign)
	webhook := fmt.Sprintf("%s&timestamp=%d&sign=%s", d.WebHook, ts, sign)
	robot := dingrobot.NewRobot(webhook)
	err := robot.SendText(content, []string{}, false)
	if err != nil {
		log.Error("钉钉消息发送失败", zap.Error(err))
	}
}
func (d *DingBot) sign(sign string) (int64, string) {
	ts := time.Now().Unix() * 1000
	s := fmt.Sprintf("%d\n%s", ts, sign)
	h := hmac.New(sha256.New, []byte(sign))
	h.Write([]byte(s))
	sign = url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	return ts, sign
}
