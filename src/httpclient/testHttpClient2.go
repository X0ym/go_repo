package main

import (
	"bytes"
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
func httpPost2(urlStr string, num int) {
	data := GetCode(1 * 1024 * 1024)
	for {
		if num <= 0 {
			break
		}
		resp, err := http.PostForm(urlStr, url.Values{"key": {data}})
		if err != nil {
			log.Fatalf("http post err: ", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("", err)
		}
		if body != nil {
			fmt.Println(string(body))
		}
		resp.Body.Close()
		num--
	}

}

func httpGet(url string, num int) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("create request failed")
	}
	req.Header.Add("key", GetCode(8*1024))

	resp, err1 := client.Do(req)
	if err1 != nil {
		fmt.Println("do get failed")
	}

	//fmt.Println(len(GetCode(3 * 1024)))
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	resp.Body.Close()
}

func whileTest2(url string, times int, num int) {
	iptablesTest2(url, times, num)
}

func iptablesTest2(url string, times int, num int) {
	for {
		if times <= 0 {
			break
		}
		fmt.Println(times)
		//go httpPost2(url, num)
		go httpGet(url, num)
		time.Sleep(10 * time.Millisecond)
		times--
	}
}
func main() {

	url := flag.String("url", "http://10.145.35.25:2046/http_proxy_test", "http url")
	times := flag.Int("times", 100, "execute times")
	num := flag.Int("num", 10, "repeat num")
	flag.Parse()
	fmt.Println("url:", *url)
	fmt.Println("goroutine:", *times)
	fmt.Println("per goroutine times:", *num)

	whileTest2(*url, *times, *num)
	time.Sleep(1000 * time.Minute)
}

func getByte2(n int) []byte {
	b := make([]byte, n)
	rand.Read(b[:])
	return b
}

// GetCode 获取一个随机用户唯一编号
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
