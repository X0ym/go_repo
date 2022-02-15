package main

func packageTest() {

}

/**
go package
	同一文件夹下只要文件都属于同一个package，文件中定义的函数可以自有调用
	go build 时，需要指定包，而不能指定包含main方法的文件 (编译时需要将其链接, 否则会找不到定义的函数)

	每个包定义一个不同的命名空间作为它的标识符，保证不与程序的其他部分冲突
	包

10.2 导入路径
	导入路径：每个包通过唯一的字符串进行标识
	对于共享或公开的包，导入路径需要全局唯一。为了避免冲突，除了标准库中的包之外，其他包的导入路径应该以互联网域名作为路径的开始，以方便找包

10.3 包的声明
	每一个Go源文件的开头都需要进行包声明。主要的目的是当该包被其他包导入时，作为其默认的标识符

10.4 导入声明
	方式一
	import "fmt"
	import "os"

	方式二
	import (
		"fmt"
		"os"
	)

	如果需要把两个相同名字的包导入到第三个包，需要进行 重命名导入。
	即导入声明中必须至少为其中一个指定替代名字来避免冲突 (仅在当前文件中生效)
	import (
		"crypto/rand"
		mrand "math/rand"
	)

10.5 空导入
	导入的包不使用，会产生一个编译错误
	如果必须导入包，只是为了对包级别的变量执行初始化表达式求值，并执行其 init 函数
	如 import _ "image/png"

10.6 包及其命名
	Go对包及其成员的命名习惯
	
*/
