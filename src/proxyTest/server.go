package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func init() {
	//设置前缀
	log.SetPrefix("[INFO]")

	//设置要打印的内容：日期，时间，长文件名
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	fileName := "httpserver.log"
	os.MkdirAll("/opt/yiming/httpserver/log", 0755)
	path := path.Join("/opt/yiming/httpserver/log", fileName)
	//打开文件，并且设置了文件打开的模式
	logFile, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	//设置输出方式为：文件
	log.SetOutput(io.MultiWriter(logFile))
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[UPSTREAM]receive request %s\n", r.URL)
	fmt.Printf("[UPSTREAM]receive request %s\n", r.URL)
	var data []byte
	var err error
	if r.Method == "POST" {
		data, err = ioutil.ReadAll(r.Body)
		log.Println(w, "request body:%s\n", string(data))
		if err != nil {
			log.Println("server recv failed. err:", err)
		}
	}

	//log.Println("server read body len:", len(data))
	//log.Println("Method:", r.Method)
	//log.Println("Host:", r.Host)
	//log.Println("RemoteAddr:", r.RemoteAddr)
	//log.Println("URL:", r.URL)
	//log.Println()

	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
	fmt.Fprintf(w, "URL: %#v\n", r.URL)
	fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
	fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
	fmt.Fprintf(w, "TLS: %#v\n", r.TLS)
	fmt.Fprintf(w, "succ\n")
	fmt.Fprintf(w, "\nHeaders:\n")
	r.Header.Write(w)

}

func ServeHTTP2(w http.ResponseWriter, r *http.Request) {
	log.Printf("[UPSTREAM]receive request %s\n", r.URL.String())
	log.Printf("Content-Length: %v", r.ContentLength)

	now := time.Now()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("server recv body failed")
	}
	length := len(data)
	log.Println("body 长度: ", length)

	w.Header().Set("Content-Length", strconv.Itoa(length))
	w.Write(data)
	time := time.Since(now)
	log.Printf("读取body 并发送响应总耗时:%s", time.String())
}

func ServeHTTP3(w http.ResponseWriter, r *http.Request) {
	log.Printf("[UPSTREAM]receive request %s\n", r.URL)
	log.Println("Content-Length: ", r.ContentLength)
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("server recv failed")
	}
	log.Println("body 长度: ", len(data))

	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	for i := 0; i < 200; i++ {
		fmt.Fprintf(w, "chunk [%02d] data: %v %s", i, time.Now(), GetCode(10))
		log.Println("chunk", i)
		flusher.Flush()
		time.Sleep(time.Millisecond * 1)
	}
}

func ServeHTTP4(w http.ResponseWriter, r *http.Request) {
	log.Printf("[UPSTREAM]receive request %s\n", r.URL)

	var data []byte
	//var err error
	//if r.Method == "POST" {
	//	data, err = ioutil.ReadAll(r.Body)
	//	//fmt.Fprintf(w, "request:%s\n", string(data))
	//	if err != nil {
	//		log.Println("server recv failed. err:", err)
	//	}
	//}
	fmt.Fprintf(w, "server read body len: %d\n", len(data))
	log.Println("server read body len:", len(data))

	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
	fmt.Fprintf(w, "URL: %#v\n", r.URL)
	fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
	fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
	fmt.Fprintf(w, "TLS: %#v\n", r.TLS)
	fmt.Fprintf(w, "succ\n")
	fmt.Fprintf(w, "\nHeaders:\n")
	r.Header.Write(w)

}

func main() {
	http.HandleFunc("/fd/check", ServeHTTP)
	http.HandleFunc("/proxytest_header", ServeHTTP)
	http.HandleFunc("/test99", ServeHTTP)
	http.HandleFunc("/proxytest_back", ServeHTTP2)
	http.HandleFunc("/proxytest_chunk", ServeHTTP3)
	http.HandleFunc("/proxytest_notReadBody", ServeHTTP4)
	http.ListenAndServe(":8000", nil)
}

func GetCode(codeLen int) string {
	// 1. 定义原始字符串
	rawStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	// 2. 定义一个buf，并且将buf交给bytes往buf中写数据
	buf := make([]byte, 0, codeLen)
	b := bytes.NewBuffer(buf)
	//随机从中获取
	rand.Seed(time.Now().UnixNano())
	for rawStrLen := len(rawStr); codeLen > 0; codeLen-- {
		randNum := rand.Intn(rawStrLen)
		b.WriteByte(rawStr[randNum])
	}
	return b.String()
}
