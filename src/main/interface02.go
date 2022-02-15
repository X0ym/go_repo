package main

import "fmt"

/**
接收类型为指针类型
	接口类型的变量只能接收指针类型的结构体，否则会编译错误

接口嵌套
	接口与接口之间可以通过嵌套创造出新接口
type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type animal interface {
	Sayer
	Mover
}

空接口
	空接口的概念
		空接口没有定义任何方法，因此任何类型都实现了空接口
		于是，空接口类型的变量可以存储任意类型的变量
	空接口的应用
		1. 空接口作为函数的参数，于是函数可以接收任意参数

*/

// ---- 验证指针类型接收-----

type People interface {
	Speak(string) string
}

type Student01 struct{}

// Speak 方法接收者为指针类型
func (stu *Student01) Speak(str string) (res string) {
	if str == "go" {
		res = "yes"
	} else {
		res = "no"
	}
	return
}

// ----------------

// ---------- 接口嵌套 ----------

type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) Say() string {
	return fmt.Sprintf("cat %s excute Say...\n", c.name)
}

func (c cat) Move() {
	fmt.Printf("cat %s excute Move...\n", c.name)
}

// --------------

// 验证空接口作为函数参数 ---------

// 空接口作为map的值 1

func InterfaceTest02() {
	var peo People = &Student01{}
	fmt.Println(peo.Speak("go"))

	var a animal = cat{name: "tom"}
	fmt.Println(a.Say())
	a.Move()

	var x interface{}
	s1 := "biubiu"
	x = s1
	fmt.Printf("type:%T value:%v\n", x, x)

	s2 := 100
	x = s2
	fmt.Printf("type:%T value:%v\n", x, x)

	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "jerry"
	studentInfo["age"] = 18
	studentInfo["isMale"] = true
	fmt.Println(studentInfo)
}
