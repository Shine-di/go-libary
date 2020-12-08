/**
 * @author: D-S
 * @date: 2020/7/25 6:50 下午
 */

package s_http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

	if r.UseProxy && r.Proxy == "" {
		proxyIp := getProxy()
		if proxyIp != "" {
			u, _ := url.Parse(proxyIp)
			client.Transport = &http.Transport{
				Proxy: http.ProxyURL(u),
			}
		}
	}
	if r.Proxy != "" {
		u, _ := url.Parse(r.Proxy)
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(u),
		}
	}
	if r.Params != nil {
		s := make([]string, 0)
		for k, v := range r.Params {
			s = append(s, fmt.Sprintf("%v=%v", k, v))
		}
		r.URL = r.URL + "?" + strings.Join(s, "&")
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
