package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file1, err1 := os.Open("/opt/service/arch_uni_platform/config/log4j.xml")
	if err1 != nil {
		fmt.Println("打开文件1失败")
	}
	defer file1.Close()

	file2, err2 := os.OpenFile("/opt/service/arch_uni_platform/config/sso.properties", os.O_RDWR, 0777)
	if err2 != nil {
		fmt.Println("打开文件2失败")
	}

	fmt.Println("done")

	name := file2.Name()
	fmt.Println("文件名：", name)
	fd := file2.Fd()
	fmt.Println("文件2的FD：", fd)
	//err := file2.Chmod(os.ModeAppend)
	//if err != nil {
	//	return
	//}

	/**

	 */
	buf1 := make([]byte, 4096)
	n, err := file2.Read(buf1)
	if err != nil {
		return
	}
	fmt.Println("读取字节数：", n)
	fmt.Println(string(buf1[:n]))

	buf2 := make([]byte, 4096)
	nn, err2 := file2.ReadAt(buf2, 2)
	if err2 != nil {
		if err2 == io.EOF {
			fmt.Println("读取到文件2结束，读取字节数：", nn)
		}
		//fmt.Println("读取文件2发送错误")
	}
	fmt.Println(string(buf2[:nn]))

	defer file2.Close()
	/**
	写入文件
	*/
	//file3, err3 := os.OpenFile("/Users/xieym/test.txt", os.O_RDWR | os.O_CREATE |os.O_APPEND, 0777)
	//if err3 != nil {
	//	fmt.Println(err3)
	//}
	//defer file3.Close()
	//nnn, err := file3.ReadFrom(file2)
	//if err != nil {
	//	fmt.Println("ReadFrom failed")
	//}
	//fmt.Println("ReadForm bytes :", nnn)

	// bufio 下的一些文件读写操作
}
