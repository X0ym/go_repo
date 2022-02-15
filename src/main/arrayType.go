package main

import "fmt"

/**
数组是相同元素的固定长度的序列，必须指定长度，不初始化则自动赋值类型默认值
数组声明
	var arr [len]type   指定长度  数组类型包括类型和数组长度，如 [3]int 和 [4]int 是不同的类型
数组初始化
	var arr [len]type
	arr = [len]type {val1, ... ,valn} 初始化列表的长度 <= len , 后面的值为类型的默认值
数组声明并初始化
	根据初始化列表初始化
	var arr [len] type = [len]type {val1, val2, ...}
	var arr = [len]type {val1, val2, ...}
	根据初始化列表指定数据类型和数组大小并初始化
	var arr = [...]type {val1, val2, ...} 或 var arr = []type {val1, val2, ...}
	指定索引值
	var arr = [...]type {inx1 : val1, inx2 : val2, ...} 或 var arr = []type {inx1 : val1, inx2 : val2, ...}

数组的遍历
	for
	for range
		for i,v := range arr { ... }
		for i := range arr { ... }
		for _,v := range arr { ... }

数组是值类型，赋值和传参会复制整个数组



*/

func arrayTest() {
	// 数组声明
	var arr1 [3]int
	var arr2 [3]int
	// 数组初始化
	arr1 = [3]int{1, 2, 3}
	arr2 = [3]int{1, 2}

	var arr3 = [3]int{1, 2}

	var arr4 = [...]int{1, 2}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	var arr5 = [...]int{0, 1, 1, 0, 0, 5}

	fmt.Println("------数组的遍历---------")
	// 数组的遍历
	for i := 0; i < len(arr5); i++ {
		fmt.Printf("i : %d \n", arr5[i])
	}

	// range 忽略index
	for _, v := range arr5 {
		fmt.Printf("%d \n", v)
	}
	// range 忽略value 或者写成 i:_
	for i := range arr5 {
		fmt.Printf("%d : %d \n", i, arr5[i])
	}
	// range index，value
	for i, v := range arr5 {
		fmt.Printf("%d : %d \n", i, v)
	}

	// 数组是值类型

}
