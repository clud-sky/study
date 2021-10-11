package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//str := "hello"
	str, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	w.Write([]byte(str))
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 对于GET请求，参数都放在url上（query param），请求体中是没有数据的
	fmt.Println(r.URL) // r.url.query
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/hello/", f1)
	http.HandleFunc("/golang/", f2)
	http.ListenAndServe("0.0.0.0:9090", nil)
}
