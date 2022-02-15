package main

import "fmt"

/**
interface 是一种类型，interface 是一组 method 的集合，
声明接口类型变量这样就可以接收所有实现了该接口的实例

值接受者和指针接收者的区别
值接收者实现接口后，接口变量既可以接收值类型也可以接收指针类型

指针类型接收者实现接口后，


*/

type Cat struct{}

type Dog struct{}

// Say 实现 Sayer 接口
func (c Cat) Say() string {
	return "Cat 喵喵喵"
}

// Say 实现 Sayer 接口
func (d Dog) Say() string {
	return "Dog 汪汪汪"
}

// Sayer 接口 定义 Say()方法
type Sayer interface {
	Say() string
}

type Mover interface {
	Move()
}

func interfaceTest() {
	//cat := Cat{}
	//fmt.Println(cat.Say())
	//dog := Dog{}
	//fmt.Println(dog.Say())

	var animal1 Sayer
	c := Cat{}
	d := Dog{}
	// Cat行为
	animal1 = c
	fmt.Println(animal1.Say())
	// Dog行为
	animal1 = d
	fmt.Println(animal1.Say())

	var animal2 Sayer = Cat{}
	fmt.Println(animal2.Say())

}
