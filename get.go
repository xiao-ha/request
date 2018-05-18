package request

import (
	"net/http"
	"compress/gzip"
	"io/ioutil"
	"net/url"
)


func (r *Request)FullUri(path string) string{
	if r.URL != nil{
		p , err := url.Parse(path)
		if err == nil{
			return r.URL.ResolveReference(p).String()
		}
	}
	return path
}


func (r *Request)Get(uri string, rel string) string {

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil{
		return ""
	}

	for k, v := range r.head {
		req.Header.Set(k, v)
	}

	if len(r.authorization) > 0{
		req.Header.Set("Authorization", r.authorization)
	}

	r.URL = req.URL

	if len(rel) > 0{
		req.Header.Set("Referer", rel)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	switch resp.Header.Get("Content-Encoding"){
	case "gzip":
		r, _ := gzip.NewReader(resp.Body)
		defer r.Close()
		body, err := ioutil.ReadAll(r)
		if err != nil {
			return ""
		}
		return string(body)
	default:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ""
		}
		return string(body)
	}
	return ""
}
