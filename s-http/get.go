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
)

type GET struct {
	URL    string
	Header map[string]string
	Proxy  string
	Token  string
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
	if r.Proxy != "" {
		u, _ := url.Parse(r.Proxy)
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(u),
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
