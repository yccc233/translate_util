package httptra

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
func PostHttp(url string, data interface{}, contentType string) string {

	// 超时时间：10秒
	client := &http.Client{Timeout: 10 * time.Second}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(result)
}


