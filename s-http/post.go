/**
 * @author: D-S
 * @date: 2020/8/8 11:26 上午
 */

package s_http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type POST struct {
	URL    string
	Header map[string]string
	Params map[string]interface{}
	Body   string
	Proxy  string
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
	req, _ := http.NewRequest("POST", r.URL, strings.NewReader(r.Body))
	req.Header.Set("Content-Type", "application/json")
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
