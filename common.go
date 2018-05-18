package request

import (
	"net/url"
	"net/http"
	"crypto/tls"
	"encoding/base64"
)

type Request struct {
	URL *url.URL
	host string
	head map[string]string
	authorization string
	client *http.Client
}


func Build(userAgent string , username string, password string) *Request{

	var r Request

	if len(username) > 0  && len(password) > 0{
		r.authorization = "Basic " + base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	}

	r.head = map[string]string{
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Accept-Encoding": "gzip, deflate",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"User-Agent":      "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.90 Safari/537.36",
	}

	if len(userAgent) > 0{
		r.head["User-Agent"] = userAgent
	}

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	r.client = &http.Client{Transport: tr}

	return &r
}