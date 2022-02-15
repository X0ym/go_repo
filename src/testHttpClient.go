package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// 大请求测试
func httpPost(urlStr string) {
	resp, err := http.PostForm(urlStr, url.Values{"key": {string(getByte(50 * 1024))}})
	if err != nil {
		log.Fatalf("http post err: ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("", err)
	}
	//fmt.Println("body:" , body)
	fmt.Println(string(body))
}

func whileTest(url string, times int, num int) {
	iptablesTest(url, times)
}

func iptablesTest(url string, times int) {
	for {
		times--

		go httpPost(url)
		time.Sleep(10 * time.Millisecond)
		if times <= 0 {
			break
		}
	}
}
func main() {
	//url := flag.String("url", "http://10.177.54.220:8001/http_proxy_test", "http url")
	url := flag.String("url", "http://127.0.0.1:2046/", "http url")
	times := flag.Int("times", 1, "execute times")
	num := flag.Int("num", 1, "repeat num")
	flag.Parse()
	fmt.Println("url:", *url)
	whileTest(*url, *times, *num)

	// 等待协程执行完
	time.Sleep(1000 * time.Minute)
}

func getByte(n int) []byte {
	b := make([]byte, n)
	rand.Read(b[:])
	return b
}
