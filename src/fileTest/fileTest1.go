package main

import (
	"fmt"
	"os"
)

func main() {
	/**
	1. 创建文件
		注意：文件的权限问题
	*/
	//file, err := os.Create("/Users/xieym/testFile.txt") // 有权限
	//if err != nil {
	//	fmt.Println("创建文件失败：", err)
	//}
	//defer file.Close()

	//file, err := os.Create("/opt/yiming/testCreateFile.txt") // 无权限
	//if err != nil {
	//	fmt.Println("创建文件失败. err:", err)
	//}
	//defer file.Close()
	/*
		result:
			创建文件失败. err: open /opt/yiming/testCreateFile.txt: permission denied
	*/

	/**
	2. 打开文件

	*/

	// 只读方式打开
	//file, err := os.Open("/Users/xieym/testFile.txt")
	//file, err := os.OpenFile("/Users/xieym/testFile.txt", os.O_RDWR|os.O_APPEND, 0666)
	//if err != nil {
	//	return
	//}
	//if err != nil {
	//	fmt.Println("打开文件失败. err:", err)
	//}
	//defer file.Close()
	//// Open函数打开文件，写入文件时报错 bad file descriptor
	//_, err = file.WriteString("this is data1.")
	//if err != nil {
	//	fmt.Println("写入文件失败，err:", err)
	//}
	//file.Sync()

	// 打开文件后，不能同时 Read 和 Write

	file1, err := os.Open("/Users/xieym/testFile.txt")
	buf := make([]byte, 1024)
	n, err := file1.Read(buf)
	if err != nil {
		fmt.Println("读取文件失败，err:", err)
	}
	fmt.Println("读取字节数：", n)
	fmt.Println("读取内容: ", string(buf[:n]))

	//err = os.Remove("/Users/xieym/testFile.txt")
	//if err != nil {
	//	fmt.Println("删除 testFile 失败, err=", err)
	//}
	//
	//err = os.Remove("/Users/xieym/emptyDir")
	//if err != nil {
	//	fmt.Println("Remove 删除空目录失败")
	//}
	//
	err = os.Remove("/Users/xieym/yiming")
	if err != nil {
		fmt.Println("Remove 删除非空目录失败, err=", err)
	}

	err = os.RemoveAll("/Users/xieym/yiming")
	if err != nil {
		fmt.Println("删除目录 yiming 失败 err=", err)
	}

}
