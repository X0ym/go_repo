package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	pr, pw := io.Pipe()
	// 开协程写入大量数据
	go func() {
		for i := 0; i < 10; i++ {
			code := GetCode(10)
			pw.Write([]byte(fmt.Sprintf("line:%d %s\r\n", i, code)))
		}
		pw.Close()
	}()
	// 传递Reader
	resp, err := http.Post("http://localhost:2046/report", "text/pain", pr)
	if err != nil {
		fmt.Println("请求失败")
	}
	fmt.Println("StatusCode: ", resp.StatusCode)
	fmt.Println("ContentLength: ", resp.ContentLength)
	fmt.Println("TransferEncoding: ", resp.TransferEncoding)

	if resp.ContentLength == -1 {
		bufReader := bufio.NewReader(resp.Body)
		var body []byte
		body, err = readBodyChunked(bufReader, 0, body)
		if err != nil {
			fmt.Println("read chunk failed. err: ", err)
		}
		fmt.Println(body)
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取body失败：", err)
		}
		fmt.Println(string(body))
	}
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

func writeChunk(dst []byte, src []byte) []byte {
	n := len(src)
	dst = writeHexInt(dst, n)
	dst = append(dst, strCRLF...)
	dst = append(dst, src...)
	dst = append(dst, strCRLF...)
	return dst
}

const (
	maxHexIntChars = 15
	upperhex       = "0123456789ABCDEF"
	lowerhex       = "0123456789abcdef"
)

var hexIntBufPool sync.Pool
var strCRLF = []byte("\r\n")

func writeHexInt(dst []byte, n int) []byte {
	if n < 0 {
		panic("BUG: int must be positive")
	}

	v := hexIntBufPool.Get()
	if v == nil {
		v = make([]byte, maxHexIntChars+1)
	}
	buf := v.([]byte)
	i := len(buf) - 1
	for {
		buf[i] = lowerhex[n&0xf]
		n >>= 4
		if n == 0 {
			break
		}
		i--
	}
	dst = append(dst, buf[i:]...)
	hexIntBufPool.Put(v)
	return dst
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

const (
	hex2intTable = "\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x00\x01\x02\x03\x04\x05\x06\a\b\t\x10\x10\x10\x10\x10\x10\x10\n\v\f\r\x0e\x0f\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\n\v\f\r\x0e\x0f\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10"
)

// ErrBrokenChunk is returned when server receives a broken chunked body (Transfer-Encoding: chunked).
type ErrBrokenChunk struct {
	error
}
