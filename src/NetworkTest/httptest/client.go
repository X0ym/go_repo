package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/*

1. http.Get(url)
2. http.Post(url, contentType string, body io.Reader)
3. http.PostForm(url string, data url.Values)
4. 先获取 Request, 再使用 Client 调用 Do 发送请求
	4.1 获取 Request
		http.NewRequest(method, url string, body io.Reader)
		http.NewRequestWithContext(ctx context.Context, method, url string, body io.Reader)
		http.ReadRequest(b *bufio.Reader)
	4.2 发送请求
		resp := Client.Do(req)

*/
func main() {
	//test1("http://127.0.0.1:8002/test2")
	test1("http://127.0.0.1:2046/proxytest_header")

	//test4("http://10.177.123.78:8001/proxytest_back", 20*1024)
}

func test1(url string) {
	resp1, err := http.Get(url)
	if err != nil {
		fmt.Println("GET 请求失败。 err=", err)
		return
	}
	fmt.Println("resp 状态码：", resp1.StatusCode)
}

func test2() {
	reqBody := []byte(getCode(2000 * 1024))
	reader := bytes.NewReader(reqBody)
	resp2, err := http.Post("http://10.177.123.78:8001/proxytest_back", "text/plain", reader)
	if err != nil {
		fmt.Println("POST 请求失败。 err=", err)
	}
	fmt.Println("reps2 状态码: ", resp2.StatusCode)
	return
}

func test3() {
	var data url.Values
	data.Set("key1", "key1-data1")
	data.Set("key1", "key1-data1")
	resp3, err := http.PostForm("http://127.0.0.1:8001/testThree", data)
	if err != nil {
		fmt.Println("POSTForM 请求失败。 err=", err)
	}
	fmt.Println("reps3 状态码: ", resp3.StatusCode)
}

func test4(url string, bodyLen int) {

	reader := strings.NewReader(getCode(bodyLen))

	client := http.Client{}
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败", err)
	}
	fmt.Println("StatusCode: ", resp.StatusCode)
	fmt.Println("ContentLength: ", resp.ContentLength)
	fmt.Println("TransferEncoding: ", resp.TransferEncoding)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应body失败：", err)
	}
	fmt.Println(string(respBody))
	fmt.Println("body len=", len(respBody))
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
