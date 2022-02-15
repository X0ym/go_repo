package main

import "fmt"

/**
自定义类型
	使用 type 关键字定义自定义类型
	方式一：type MyInt int 使用基本类型定义新类型
	方式二：使用结构体 struct

类型别名
	type TypeAlias = type
	之前的 rune和byte就是类型别名
		type byte = unit8
		type rune = int32
结构体

（1）定义结构体
	使用 type 和 struct 关键字定义结构体
	使用 var 声明结构体 或短变量声明
（2）临时结构体
	在定义一些临时数据的场景下还可以使用匿名结构体
（3）指针类型结构体
	使用关键字 new 对结构体进行初始，得到结构体的地址
	var p = new(person)

	指针结构体初始化
		1. .访问成员初始化
		2. &取结构体地址，键值对初始化
		3. &取结构体地址，值列标配初始化
（4）结构体初始化
	没有初始化的结构体，其成员变量都是对应类型的零值
	1. 使用键值对初始化
	2. 使用 . 访问成员初始化
	3. 值列表初始化 注意：必须初始化所有字段，且初始化顺序一致

（5）构造函数
	结构体没有构造函数，struct 是值类型，值拷贝的开销大，可使用指针类型的结构体

（6）方法
	方法是一种作用域特定类型的函数 这种特定类型变量叫做 接受者 Receiver。
	方法的定义
		func (recv Receiver) MethodName(参数列表) 返回值列表 { ... }

（7）结构体中的匿名字段
	结构体中允许其成员字段在声明的时候没有字段名而只有类型
	默认会采用类型名作为字段名
	结构体要求字段名称必须唯一，因此结构体中同种类型的匿名字段只能有一个


*/

// 使用 type 关键字定义结构体
type person struct {
	name string // name , city string
	age  int8
	city string
}

// 结构体 person 的构造函数
// 键值对方式
func newPerson1(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

// 值列表方式
func newPerson(name, city string, age int8) *person {
	return &person{
		name,
		age,
		city,
	}
}

type User struct {
	Name string
	Age  int8
	City string
}

func newUser(name, city string, age int8) *User {
	return &User{
		Name: name,
		Age:  age,
		City: city,
	}
}

func (u User) Dream() {
	fmt.Printf("User:%s on going", u.Name)
}

func structTest1() {

	// 结构体实例化
	// 方式一：var 声明结构体变量，直接使用 . 初始化
	var tom person
	fmt.Printf("tom=%#v\n", tom) // 未初始化
	tom.name = "刹车娜扎"
	tom.city = "北京"
	tom.age = 18
	fmt.Printf("tom=%v\n", tom)
	fmt.Printf("tom=%#v\n", tom)

	// 方式二 ：匿名结构体
	var user1 struct {
		name string
		age  int8
	}
	user1.name = "yiming"
	user1.age = 24
	fmt.Printf("user1=%v\n", user1)
	fmt.Printf("user1=%#v\n", user1)

	// 方式三 指针类型结构体
	var jerry = new(person)
	// 初始化方式一：直接对结构体指针使用 . 访问结构体的成员或字段
	jerry.name = "biubiu"
	jerry.city = "上海"
	jerry.age = 24

	// 初始化方式二： 指针类型， & 取结构体地址进行键值对初始化
	p1 := &person{
		name: "name1",
		city: "city1",
		age:  20,
	}
	fmt.Printf("person name1=%#v \n", p1)

	// 结构体初始化方式
	fmt.Println("------- 结构体初始化方式 -------")
	// 方式一：.
	// 方式二：键值对
	user21 := person{
		name: "name2",
		age:  21,
	}
	fmt.Printf("user21=%v\n", user21)
	fmt.Printf("user21=%#v\n", user21)
	// 方式三：使用值得列表初始化
	user22 := person{
		"name2",
		22,
		"武汉",
	}
	fmt.Printf("user22=%v\n", user22)
	fmt.Printf("user22=%#v\n", user22)

	// 测试方法 接收类型 User
	u := newUser("Go", "city", 10)
	u.Dream()

}
