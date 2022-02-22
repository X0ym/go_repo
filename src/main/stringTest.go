package main

import "fmt"

func main() {
	str := "go 语言"
	// len() 函数 得到string的字节数
	fmt.Println("len(str) = ", len(str))
	fmt.Println("len(\"语言\") = ", len("语言"))

	// str[i] 直接取第i个字符，0 <= i < len(str)
	fmt.Println(str[0], str[5])
	// 第i个字节不一定是第i个字符

	// 支持 + 运算符拼接字符串
	fmt.Println("str1" + str)

}
