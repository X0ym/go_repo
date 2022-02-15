package main

import "fmt"

/**
	标识符：由任意数量的字符，数字或下划线组成，区分大小写，且必须以字符或下划线开头
	声明：变量 var , 常量 const , 类型 type , 函数 func

	变量声明
		通过 var 创建一个变量，变量的声明格式：
		方式一：标准声明 var name type = expression
		方式二：批量声明
				(1)	var (
						name1 string
						name2 int
						name3 bool
						name4 float32
						)
				(2) var i, j, k int // int int int
		注意：类型或表达式可省略一个，但不能都省
			 Go语言中的变量需要声明后才能使用
			 同一作用域内不支持重复声明
 			 Go语言的变量声明后必须使用
	变量初始化
		变量声明后，Go会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值
		变量的默认值
			 整型和浮点型变量的默认值为0
			 string 默认值为空字符串
			 bool 默认为false
			 切片、函数、指针变量 默认为nil
		初始化方式
			1. 使用表达式赋初值 var name type = expression
			2. 直接赋值
				声明忽略类型，根据初始值确定类型  var b, f, s = true, 2.3, "four" --- 可初始化多个值
				声明不忽略类型 并赋值对应类型的初始值 var name type = value
			3. 函数返回值
				单返回值 var name = function(x)
				多返回值 var f,err = function(x)
	短变量声明
		在函数内部使用 := 方式声明并初始化变量
		类似var声明，短变量声明也可以进行多变量声明并初始化 如：i, j := 0, 1
		注意：一定要在函数内使用
	类型推导
		变量初始化时可省略类型，编译器根据表达式的值推导变量的类型并完成初始化
	匿名变量 （go支持多返回值）
		在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线_表示


	常量
		常量定义时必须初始化
		关键字 const

	实体的作用域


*/

// Var1 常量 包可见范围
const Var1 = 20.0

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func foo() (int, string) {
	return 10, "干"
}

func NameTest() {
	var var2 = 10 // 可见范围 函数内
	var total = Var1 + var2

	name := "Go"
	age := 18
	fmt.Println("name: " + name)
	fmt.Println("age: ", age)
	//var x int
	x, _ := foo()
	_, y := foo()
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(total)

	fmt.Println(KB, MB, GB, TB, PB)
}

func nameTest2() {
	var 我 int = 3
	fmt.Println(我)
}
