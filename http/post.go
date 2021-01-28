/**
 * @author: D-S
 * @date: 2020/8/8 11:26 上午
 */

package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type POST struct {
	URL    string
	Header map[string]string
	Params map[string]interface{}
	// Content-Type 默认是json格式
	Body []byte
	// Content-Type application/x-www-form-urlencoded; charset=UTF-8
	FormData map[string]interface{}
	Proxy    string
}

var client = new(http.Client)

func (r *POST) Do() ([]byte, error) {
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
	var body io.Reader
	// body 是json
	if r.Body != nil {
		body = bytes.NewReader(r.Body)
		if r.Header != nil {
			r.Header["Content-Type"] = "application/json"
		} else {
			r.Header = map[string]string{
				"Content-Type": "application/json",
			}
		}
	}
	if r.FormData != nil {
		s := make([]string, 0)
		for k, v := range r.Params {
			s = append(s, fmt.Sprintf("%v=%v", k, v))
		}
		body = strings.NewReader(strings.Join(s, "&"))
		if r.Header != nil {
			r.Header["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
		} else {
			r.Header = map[string]string{
				"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
			}
		}
	}
	req, _ := http.NewRequest("POST", r.URL, body)
	for k, v := range r.Header {
		req.Header.Add(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
