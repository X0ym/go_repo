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
	os.MkdirAll("/opt/httpserver/log", 0755)
	path := path.Join("/opt/httpserver/log", fileName)
	//打开文件，并且设置了文件打开的模式
	logFile, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	//设置输出方式为：文件
	log.SetOutput(io.MultiWriter(logFile))
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[UPSTREAM] receive request %s\n", r.URL.String())
	log.Println("Method:", r.Method)
	log.Println("RemoteAddr:", r.RemoteAddr)
	log.Println("Content-Length: ", r.ContentLength)

	var data []byte
	var err error
	data, err = ioutil.ReadAll(r.Body)
	log.Println("body len=", len(data))
	if err != nil {
		log.Println("server recv failed. err:", err)
	}
	if len(data) <= 1024 {
		log.Println(string(data))
	}

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
	log.Println("[UPSTREAM] receive request ", r.URL.String())
	log.Println("Method:", r.Method)
	log.Println("RemoteAddr:", r.RemoteAddr)
	log.Println("Content-Length: ", r.ContentLength)

	now := time.Now()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("server recv body failed")
	}
	length := len(data)
	log.Println("body 长度: ", length)
	if len(data) <= 1024 {
		log.Println(string(data))
	}

	w.Header().Set("Content-Length", strconv.Itoa(length))
	w.Write(data)
	time := time.Since(now)
	log.Printf("读取body 并发送响应总耗时:%s\n", time.String())
}

func ServeHTTP3(w http.ResponseWriter, r *http.Request) {
	log.Printf("[UPSTREAM]receive request %s\n", r.URL)
	log.Println("Method:", r.Method)
	log.Println("RemoteAddr:", r.RemoteAddr)
	log.Println("Content-Length: ", r.ContentLength)
	chunkSize := r.Header.Get("chunkSize")
	log.Println("recv req Header - ChunkSize=", chunkSize)
	//var size int
	size, err := strconv.Atoi(chunkSize)
	if err != nil {
		size = 1024
	}
	log.Println("Response chunkSize = ", size)

	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "chunk[%02d] %s", i, GetCode(size))
		flusher.Flush()
		time.Sleep(time.Millisecond * 1)
	}
}

func ServeHTTP4(w http.ResponseWriter, r *http.Request) {
	log.Printf("[UPSTREAM]receive request %s\n", r.URL.String())

	//var err error
	//if r.Method == "POST" {
	//	data, err = ioutil.ReadAll(r.Body)
	//	//fmt.Fprintf(w, "request:%s\n", string(data))
	//	if err != nil {
	//		log.Println("server recv failed. err:", err)
	//	}
	//}
	log.Println("not read req body")

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
	http.HandleFunc("/proxytest_header", ServeHTTP)
	http.HandleFunc("/proxytest_back", ServeHTTP2)
	http.HandleFunc("/proxytest_chunk", ServeHTTP3)
	http.HandleFunc("/proxytest_notReadBody", ServeHTTP4)
	http.ListenAndServe(":8001", nil)
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
