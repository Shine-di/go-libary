/**
 * @author: D-S
 * @date: 2020/7/25 6:50 下午
 */

package http

import (
	"crypto/tls"
	"fmt"
	"github.com/dishine/libary/log"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type GET struct {
	URL      string
	Header   map[string]string
	Params   map[string]string
	UseProxy bool
	Proxy    string
	Token    string
}

func (r *GET) Do() ([]byte, error) {
	if r.Token != "" {
		l, _ := url.Parse(r.URL)
		q := l.Query()
		q.Add("token", r.Token)
		l.RawQuery = q.Encode()
		r.URL = l.String()
	}
	client := new(http.Client)
	if r.Params != nil {
		s := make([]string, 0)
		for k, v := range r.Params {
			s = append(s, fmt.Sprintf("%v=%v", k, v))
		}
		r.URL = r.URL + "?" + strings.Join(s, "&")
	}
	if r.UseProxy && r.Proxy == "" {
		proxyIp := getProxy()
		if proxyIp != "" {
			log.Info("使用代理", zap.Any("ip", proxyIp))
			u, _ := url.Parse(proxyIp)
			client.Transport = &http.Transport{
				Proxy:           http.ProxyURL(u),
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
		}
	}
	if r.Proxy != "" {
		log.Info("使用代理", zap.Any("ip", r.Proxy))
		u, _ := url.Parse(r.Proxy)
		client.Transport = &http.Transport{
			Proxy:           http.ProxyURL(u),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	client.Timeout = 10 * time.Second
	if client.Transport == nil {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	req, _ := http.NewRequest("GET", r.URL, nil)
	for k, v := range r.Header {
		req.Header.Add(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
