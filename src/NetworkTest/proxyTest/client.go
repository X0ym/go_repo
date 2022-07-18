package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	str := flag.String("str", "test2", "功能选择")
	url := flag.String("url", "http://www.baidu.com", "set url")
	j := flag.Int64("bodyLen", 20*1024, "flag int64 body len")
	s := flag.Int64("chunkSize", 1024, "flag int64 chunk size")

	flag.Parse()

	fmt.Println("str=", *str)
	if *str == "test1" {
		fmt.Println("req body len=", *j)
		fmt.Println("do test1")
		test1(*url, *j)
	}
	if *str == "test2" {
		fmt.Println("chunk size=", *s)
		fmt.Println("do test2")
		fmt.Println("url: ", *url)
		test2(*url, *s)
	}

}

func test1(url string, bodyLen int64) {
	req1 := &fasthttp.Request{}
	req1.SetRequestURI(url)
	resp1 := &fasthttp.Response{}

	req1.SetBody([]byte(getCode(bodyLen)))
	req1.Header.Set("bodyLen", "20480")

	if err := fasthttp.Do(req1, resp1); err != nil {
		fmt.Println(err)
		fmt.Println("请求失败")
	}

	fmt.Println("Response header:")
	resp1.Header.VisitAll(func(key, value []byte) {
		fmt.Printf("key:%s value:%s\n", string(key), string(value))
	})

	fmt.Println("Response body:")
	body := resp1.Body()
	fmt.Println(string(body))
	fmt.Println("body len=", len(body))
}

func test2(url string, chunkSize int64) {
	pr, pw := io.Pipe()
	// 开协程写入大量数据
	go func() {
		for i := 0; i < 10; i++ {
			code := getCode(chunkSize)
			pw.Write([]byte(fmt.Sprintf("line:%d %s\r\n", i, code)))
		}
		pw.Close()
	}()
	// 传递Reader
	client := http.Client{}
	req, err := http.NewRequest("POST", url, pr)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "text/pain")
	req.Header.Set("chunkSize", strconv.FormatInt(chunkSize, 10))
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

func getCode(codeLen int64) string {
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

func readBodyChunked(r *bufio.Reader, maxBodySize int, dst []byte) ([]byte, error) {
	if len(dst) > 0 {
		panic("BUG: expected zero-length buffer")
	}

	strCRLFLen := len(strCRLF)
	for {
		chunkSize, err := parseChunkSize(r)
		if err != nil {
			return dst, err
		}
		if maxBodySize > 0 && len(dst)+chunkSize > maxBodySize {
			return dst, ErrBodyTooLarge
		}
		dst, err = appendBodyFixedSize(r, dst, chunkSize+strCRLFLen)
		if err != nil {
			return dst, err
		}
		if !bytes.Equal(dst[len(dst)-strCRLFLen:], strCRLF) {
			return dst, ErrBrokenChunk{
				error: fmt.Errorf("cannot find crlf at the end of chunk"),
			}
		}
		dst = dst[:len(dst)-strCRLFLen]
		if chunkSize == 0 {
			return dst, nil
		}
	}
}

func appendBodyFixedSize(r *bufio.Reader, dst []byte, n int) ([]byte, error) {
	if n == 0 {
		return dst, nil
	}

	offset := len(dst)
	dstLen := offset + n
	if cap(dst) < dstLen {
		b := make([]byte, round2(dstLen))
		copy(b, dst)
		dst = b
	}
	dst = dst[:dstLen]

	for {
		nn, err := r.Read(dst[offset:])
		if nn <= 0 {
			if err != nil {
				if err == io.EOF {
					err = io.ErrUnexpectedEOF
				}
				return dst[:offset], err
			}
			panic(fmt.Sprintf("BUG: bufio.Read() returned (%d, nil)", nn))
		}
		offset += nn
		if offset == dstLen {
			return dst, nil
		}
	}
}

func round2(n int) int {
	if n <= 0 {
		return 0
	}
	n--
	x := uint(0)
	for n > 0 {
		n >>= 1
		x++
	}
	return 1 << x
}

func parseChunkSize(r *bufio.Reader) (int, error) {
	n, err := readHexInt(r)
	if err != nil {
		return -1, err
	}
	for {
		c, err := r.ReadByte()
		if err != nil {
			return -1, ErrBrokenChunk{
				error: fmt.Errorf("cannot read '\r' char at the end of chunk size: %s", err),
			}
		}
		// Skip any trailing whitespace after chunk size.
		if c == ' ' {
			continue
		}
		if c != '\r' {
			return -1, ErrBrokenChunk{
				error: fmt.Errorf("unexpected char %q at the end of chunk size. Expected %q", c, '\r'),
			}
		}
		break
	}
	c, err := r.ReadByte()
	if err != nil {
		return -1, ErrBrokenChunk{
			error: fmt.Errorf("cannot read '\n' char at the end of chunk size: %s", err),
		}
	}
	if c != '\n' {
		return -1, ErrBrokenChunk{
			error: fmt.Errorf("unexpected char %q at the end of chunk size. Expected %q", c, '\n'),
		}
	}
	return n, nil
}

func readHexInt(r *bufio.Reader) (int, error) {
	n := 0
	i := 0
	var k int
	for {
		c, err := r.ReadByte()
		if err != nil {
			if err == io.EOF && i > 0 {
				return n, nil
			}
			return -1, err
		}
		k = int(hex2intTable[c])
		if k == 16 {
			if i == 0 {
				return -1, errEmptyHexNum
			}
			if err := r.UnreadByte(); err != nil {
				return -1, err
			}
			return n, nil
		}
		if i >= maxHexIntChars {
			return -1, errTooLargeHexNum
		}
		n = (n << 4) | k
		i++
	}
}

var (
	errEmptyHexNum    = errors.New("empty hex number")
	errTooLargeHexNum = errors.New("too large hex number")
	ErrBodyTooLarge   = errors.New("body size exceeds the given limit")
)

var strCRLF = []byte("\r\n")

const (
	hex2intTable   = "\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x00\x01\x02\x03\x04\x05\x06\a\b\t\x10\x10\x10\x10\x10\x10\x10\n\v\f\r\x0e\x0f\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\n\v\f\r\x0e\x0f\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10"
	maxHexIntChars = 15
	upperhex       = "0123456789ABCDEF"
	lowerhex       = "0123456789abcdef"
)

// ErrBrokenChunk is returned when server receives a broken chunked body (Transfer-Encoding: chunked).
type ErrBrokenChunk struct {
	error
}
