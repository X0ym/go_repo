package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/**
Go基本数据类型
	整型
		主要分为无符号和有符号
		无符号整型
			uint8	无符号 8位整型 (0 到 255)
			uint16	无符号 16位整型 (0 到 65535)
			uint32	无符号 32位整型 (0 到 4294967295)
			uint64	无符号 64位整型 (0 到 18446744073709551615)
		有符号整型
			int8	有符号 8位整型 (-128 到 127)
			int16	有符号 16位整型 (-32768 到 32767)
			int32	有符号 32位整型 (-2147483648 到 2147483647)
			int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)
		特殊整型 使用时考虑不同机器的差异
			uint	32位操作系统上就是uint32，64位操作系统上就是uint64
			int	32位操作系统上就是int32，64位操作系统上就是int64
			uintptr	无符号整型，用于存放一个指针
	数字字面量
		二进制 v := 0b0011010
		八进制 v := 0o377
		十六进制 v := 0x123

	浮点数
		float32
		float64

	复数
		complex64	实数和虚数为32位
		complex128	实数和虚数为64位

	布尔型bool	Go中bool不允许将整型强制转位bool
		true
		false

	字符串string
		原生类型，string内部utf-8编码
		转义字符
			转义符	含义
			\b		退格符
			\r		回车符（返回行首）
			\n		换行符（直接跳到下一行的同列位置）
			\t		制表符
			\'		单引号
			\"		双引号
			\\		反斜杠
		多行字符串 用 ``，原生的字符串字面量的书写形式，其中转移序列不起作用，那么实质内容与字面写法严格一致，包括反斜杠和换行符
		原生的字符串字面量可以展开多行，唯一的特殊处理是删除回车符
	字符 byte 和 rune类型
		byte型或者称 unit8型，代表一个 ASCII字符
		rune型，实际是一个 int32 类型,代表一个 utf-8 类型字符


	类型转换
		Go语言中只有强制类型转换，没有隐式类型转换
		语法：T(表达式)
			T表示要转化的目标类型，表达式包括变量、复杂算子和函数返回值等


*/
func typeTest() {
	var a1 int8 = 10
	var b1 int16 = 11
	var c1 int32 = 475
	var d1 int64 = 321
	fmt.Printf("int8:%d , int16:%d , int32:%d , int64:%d \n", a1, b1, c1, d1)

	// c := Var1 + a

	a2 := 0b0101101
	b2 := 0o377
	c2 := 0x3FFF
	fmt.Printf("二进制L：%b, 八进制：%o, 十六进制：%x \n", a2, b2, c2)
	fmt.Printf("二进制对应十进制：%d, 八进制对应十进制：%d, 十六进制对应十进制：%d \n", a2, b2, c2)

	a3 := 1 + 2i
	b3 := 2 + 1i
	var c3 complex64 = 1 + 3i
	var d3 complex128 = 1283 + 122i
	fmt.Println(a3, b3, c3, d3)

	fmt.Println("--------字符串---------")
	a4 := "hello go"
	b4 := "1213"
	var c4 string = "hello world"
	// d4 多行字符串
	var d4 = `first line\r\n 
second line`

	fmt.Println(a4, b4, c4)
	fmt.Println(d4)
	fmt.Println("-------------")
	fmt.Println(len(a4), len(b4))
	fmt.Println(a4 + b4)
	fmt.Println(strings.Split(c4, " "))
	fmt.Println(a4, "hello")
	fmt.Println(strings.Index(c4, "world"))

	fmt.Println("------字符--------")
	var a5 = 'c'
	b5 := 'z'
	var c5 byte = 'a'
	fmt.Println(a5, " ", b5, " ", c5)
	s := string("hello world")
	traversalString(s)

	s11 := "hello,世界"
	fmt.Println(s11)
	fmt.Println(utf8.DecodeRuneInString(s11))

}

func traversalString(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)\n", s[i], s[i])
	}
}
