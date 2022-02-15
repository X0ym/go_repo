package main

import "fmt"

/**
函数
	函数定义
		func 函数名(参数)(返回值){
    		函数体
		}
		函数名：命名规则即标识符命名规则，在同一个包内，函数名也称不能重名
		参数：参数变量和参数变量的类型组成，多个参数之间使用,分隔
		返回值：由返回值变量和其变量类型组成，也可以只写返回值的类型；多个返回值必须用()包裹，并用,分隔
		函数体
	参数
		类型简写，函数的参数中如果相邻变量的类型相同，则可以类型指定合在一起,但需要用逗号隔开，如 func intSum (x, y int) int { ... }
		可变参数，指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识, 如 func intSum (x int , y... int) { ... }
			可变参数通常要作为函数的最后一个参数
			本质上，函数的可变参数是通过切片来实现的
	返回值
		多返回值
		定义返回值变量名，可在函数中直接使用，默认为类型的默认值
	函数类型
		定义函数类型；可用关键字 type 定义函数类型，格式如下：
			type funcTypeName func(type_1, ... , type_n) (type_11, ... , type_nn)
		符合函数类型定义的函数都属于该函数类型
		定义函数类型变量
			var name funcTypeName = functionName
	函数作为参数或返回值

	匿名函数
		只能在函数内部定义匿名函数，定义格式如下：
			func ( 参数 ) ( 返回值 ){
    			函数体
			}
		匿名函数由于没有函数名，则匿名函数需要保存在某个变量中，或者作为立即执行函数
		匿名函数多用于回调函数和闭包

	闭包

	defer语句



*/

// 可变参数
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

// 多返回值
func calc1(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 定义返回值变量，在函数中可直接使用 (初始化值？)
func calc2(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 当函数的返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片
func someFunc(x int) []int {
	if x == 0 {
		return nil
	} else {
		return []int{1, 2}
	}
}

// 定义函数类型
type calculation func(int, int) int

// calculation类型函数
func add(x, y int) (sum int) {
	sum = x + y
	return
}

// calculation类型函数
func sub(x, y int) (res int) {
	res = x - y
	return
}

func functionTest() {
	res1 := intSum2()
	res2 := intSum2(10, 20)
	res3 := intSum2(10, 20, 30)
	fmt.Println(res1, res2, res3)

	x, y := 20, 10
	sum1, sub1 := calc1(x, y)
	sum2, sub2 := calc2(x, y)
	fmt.Println("sum1:", sum1, "sub1", sub1)
	fmt.Println("sum2", sum2, "sub2:", sub2)

	// 定义函数类型变量
	var function calculation = add
	i := function(1, 2) // 像 C 语言那样调用
	fmt.Println(i)

	// 匿名函数
	// 方式一：保存在变量
	add1 := func(x, y int) int {
		return x + y
	}
	// 调用匿名函数
	add1(10, 20)

	// 方式二：直接调用
	func(x, y int) {
		fmt.Println(x + y)
	}(11, 21)

}
