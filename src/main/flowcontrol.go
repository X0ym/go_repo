package main

import "fmt"

/**
if 分支结构 条件判断
for 循环结构
	for 初始化；条件；结束语句
	for 条件  类似 while循环
	for 无限循环
for range (学引用数据类型之后再练习）
	遍历数组、切片、字符串、map 及通道（channel）。 通过for range遍历的返回值有以下规律：
		1. 数组、切片、字符串返回索引和值
		2. map返回键和值
		3. 通道（channel）只返回通道内的值
switch case （待学习）
*/

func flowTest() {
	score := 30
	if score >= 80 {
		fmt.Println("及格")
	} else if score >= 60 {
		fmt.Println("优秀")
	} else {
		fmt.Println("不及格")
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 无限循环
	for {
		fmt.Println("无限循环")
	}
}
