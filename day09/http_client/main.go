package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//resp,err := http.Get("http://0.0.0.0:9090/golang/?name=刘洋君&age=33")

	data := url.Values{}
	urlObj, _ := url.Parse("http://0.0.0.0:9090/golang/")
	data.Set("name", "刘洋君")
	data.Set("age", "33")
	urlStr := data.Encode() // url encode之后的url
	urlObj.RawQuery = urlStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("get url failed,err:", err)
		return
	}
	defer resp.Body.Close()
	// 从resp中把服务器返回数据读出来
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read url failed,err:", err)
		return
	}
	fmt.Println(string(b))
}
