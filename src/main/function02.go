package main

import "fmt"

/**
函数声明

匿名函数

变长函数
	在参数列表最后的类型名称前使用 "..."表明这是一个变长函数

*/

// 匿名函数示例
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

// 至少给一个参数
func max(num1 int, nums ...int) int {
	maxNum := num1
	if len(nums) < 1 {
		return maxNum
	} else {
		for _, num := range nums {
			if num > maxNum {
				maxNum = num
			}
		}
		return maxNum
	}

}

func functionTest02() {
	// 测试匿名函数
	f := squares() // 执行squares()返回匿名函数  f 为匿名函数的函数变量

	fmt.Println(f()) //调用匿名函数，输出 1  其中，x = 1
	fmt.Println(f()) //调用匿名函数，输出 4  其中，x = 2
	fmt.Println(f()) //调用匿名函数，输出 9  其中，x = 3

	fmt.Println(sum())
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3, 4))

	s := []int{1, 2, 3}
	fmt.Println(sum(s...))

	// 测试max 至少需要一个参数
	fmt.Println(max(1))
	fmt.Println(max(1, 2))
	fmt.Println(max(1, 2, 3, 4))

}
