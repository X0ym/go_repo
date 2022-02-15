package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//fileInfos, err := ioutil.ReadDir("/opt/web")
	//if err != nil {
	//	return
	//}
	//for _, info := range fileInfos {
	//	if info.IsDir() {
	//		fmt.Println("dir : ", info.Name())
	//	} else {
	//		fmt.Println("文件名", info.Name())
	//	}
	//}
	//
	//_, err3 := os.Lstat("/opt/service/arch_uni_platform/config/sso.propertie")
	//if os.IsNotExist(err3) {
	//	fmt.Println("不存在")
	//}
	//
	//fmt.Println(err3)
	////fmt.Println(info.Name())
	////fmt.Println(info.Size())

	for i := 0 ; i < 10 ; i ++ {
		root, err := getPath()
		if err != nil {
			return
		}
		fmt.Println("getPath() : ", root)
	}
	CleanHttpBody()
}

var path1 = ""

func getPath() (string, error) {
	if path1 == "" {
		fmt.Println("初始化目录")
		dir, err := ioutil.TempDir("", "proxy-")
		if err != nil {
			return "", err
		}
		path1 = dir
		fmt.Println("创建path:", path1)
		return path1, nil
	}
	fmt.Println("直接返回目录:", path1)
	return path1, nil
}

func CleanHttpBody() error {
	// 删除对应的目录即可
	if path1 != "" {
		fmt.Println("path:", path1)
		err := os.Remove(path1)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	fmt.Println("delete success")
	return nil
}