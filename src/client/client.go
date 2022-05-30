package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// 大请求测试
func httpPost2(urlStr string, num int, body int) {
	data := getCode(body)
	//fmt.Printf("key=%s\n", data)
	for {
		if num <= 0 {
			break
		}

		t1 := time.Now()
		fmt.Println("开始发送")
		resp, err := http.PostForm(urlStr, url.Values{"key": {data}})
		if err != nil {
			fmt.Println("发送失败. err:", err)
			//log.Fatalf("http post err: ", err)
		}

		fmt.Println(resp.StatusCode)
		resdata, err := io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		times := time.Since(t1)
		fmt.Println("ContentLength:", resp.ContentLength)
		fmt.Println("总耗时: ", times)
		fmt.Println("body: ", string(resdata))
		fmt.Println("res body len:", len(resdata))
		time.Sleep(100 * time.Millisecond)
		num--
	}

}

func httpGet(url string, num int) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("create request failed")
	}
	req.Header.Add("key", getCode(12*1024))
	//fmt.Println(req.Header.Get("key"))
	resp, err1 := client.Do(req)
	if err1 != nil {
		fmt.Println("do get failed")
	}
	fmt.Println(resp.StatusCode)

	//fmt.Println(len(GetCode(3 * 1024)))
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	resp.Body.Close()
}

func whileTest2(url string, times int, num int, body int) {
	iptablesTest2(url, times, num, body)
}

func iptablesTest2(url string, times int, num int, body int) {
	for {
		if times <= 0 {
			break
		}
		fmt.Println("goroutine1:", times)
		go httpPost2(url, num, body)
		//go httpGet(url, num)
		time.Sleep(10 * time.Millisecond)
		times--
	}
}
func main() {
	//url := flag.String("url", "http://10.177.125.17:8001/proxytest_header", "http url")
	url := flag.String("url", "http://10.177.123.78:8001/proxytest_back", "http url")
	//url := flag.String("url", "http://127.0.0.1:8001/proxytest_chunk", "http url")
	//url := flag.String("url", "http://10.177.125.17:8001/proxytest_chunk", "http url")
	//url := flag.String("url", "http://127.0.0.1:2046/proxytest_back", "http url")
	//url := flag.String("url", "http://10.177.125.17:8001/proxytest_notReadBody", "http url")

	times := flag.Int("times", 1, "execute times")
	num := flag.Int("num", 1, "repeat num")
	body := flag.Int("body", 20*1024*1024, "body length in byte")
	flag.Parse()
	fmt.Println("url:", *url)
	fmt.Println("goroutine:", *times)
	fmt.Println("per goroutine times:", *num)
	fmt.Println("body len: ", *body)

	whileTest2(*url, *times, *num, *body)
	time.Sleep(100 * time.Minute)
}

func getByte2(n int) []byte {
	b := make([]byte, n)
	rand.Read(b[:])
	return b
}

func getCode(codeLen int) string {
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
