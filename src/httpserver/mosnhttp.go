package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/**

一个简单的HTTP服务

*/

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("timestest env recv request. time:" + time.Now().String())
	fmt.Println(r.Method)
	fmt.Println(r.Header)
	fmt.Println(r.UserAgent())
	fmt.Println(r.URL)
	fmt.Println()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("server read failed. err : ", err)
	}

	if body != nil {
		fmt.Println(string(body))
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
	fmt.Fprintf(w, "URL: %#v\n", r.URL)
	fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
	fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
	fmt.Fprintf(w, "TLS: %#v\n", r.TLS)
	fmt.Fprintf(w, "\nHeaders:\n")
	fmt.Fprintf(w, "key's len:%d\n", len(r.Header.Get("key")))

	// 休眠代表服务端处理时间
	time.Sleep(time.Minute * 0)
}

func main() {
	port := flag.String("port", "8001", "http port")
	flag.Parse()
	fmt.Println("port:", *port)
	http.HandleFunc("/proxytest_header", ServeHTTP)
	http.HandleFunc("/9", ServeHTTP)
	http.HandleFunc("/test99", ServeHTTP)
	http.HandleFunc("/test_new99", ServeHTTP)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		fmt.Println("server listen failed. err : ", err)
	}
}
