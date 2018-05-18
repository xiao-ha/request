package request

import (
	"net/http"
	"compress/gzip"
	"io/ioutil"
	"bytes"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const(
	X_WWW_FORM_URLENCODED = iota
	JSON

)

func (r *Request)Post(uri string, data interface{} , rel string , postType int) string {
	var body string

	switch postType{
	case X_WWW_FORM_URLENCODED:
		dataMap := data.(map[string]string)
		for k,v := range dataMap{
			body += k + "=" + v + "&"
		}
		body = body[:len(body)-2]
	case JSON:
		dataJson , err := json.Marshal(&data)
		if err != nil{
			return ""
		}
		body = string(dataJson)
	}


	req, err := http.NewRequest("POST", uri, bytes.NewReader([]byte(body)))
	if err != nil{
		return ""
	}
	switch postType{
	case X_WWW_FORM_URLENCODED:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	case JSON:
		req.Header.Set("Content-Type", "application/json")
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