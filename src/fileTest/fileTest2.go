package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {

	info, err := os.Lstat("/Users/xieym/test.txt")
	if !os.IsNotExist(err) {
		fmt.Println("存在该文件. name: ", info.Name())
		fmt.Println("删除")
		os.Remove("/Users/xieym/test.txt")
	}
	if os.IsNotExist(err) {
		fmt.Println("不存在")
		//return
	}

	file, err11 := os.OpenFile("/Users/xieym/test.txt", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err11 != nil {
		fmt.Println("打开文件失败")
		fmt.Println(err11)
	}

	data := "This is test data."
	n1, err12 := file.Write([]byte(data))
	if err12 != nil {
		fmt.Println("写入错误")
	}
	fmt.Println("写入字节数:", n1)

	err2 := os.Remove(file.Name())
	if err2 != nil {
		fmt.Println("delete failed. err:", err2)
	}
	fmt.Println("删除成功")

	err1 := file.Close()
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("删除文件后，关闭成功")

	dir := os.TempDir()
	fmt.Println("os tempDir: ", dir)
	root := path.Join(dir, "mesh", "bigBody")
	//root := dir + "mesh" + string(os.PathSeparator) + "bigBody"
	err4 := os.MkdirAll(root, 0700)
	if err4 != nil {
		fmt.Println("创建根目录失败")
		fmt.Println(err4)
	}

	tempDir, err := ioutil.TempDir(root, "proxy-")
	if err != nil {
		fmt.Println("创建临时目录失败")
	}
	fmt.Println("创建临时目录成功. Path:", tempDir)

	//f, err14 := os.Open("/Users/xieym/test.txt")
	//if err14 != nil {
	//	fmt.Println("打开文件失败")
	//	fmt.Println(err14)
	//}
	//fmt.Println("成功打开文件")
	//read := make([]byte, 1024)
	//nn, err13 := f.Read(read)
	//if err13 != nil {
	//	fmt.Println("读取错误")
	//}
	//fmt.Println(string(read[:nn]))
	//if strings.EqualFold(string(read[:nn]), data) {
	//	fmt.Println("写入与读取正常")
	//}
	//
	//f.Close()
}
