package main

import (
	"encoding/json"
	"fmt"
)

/**
嵌套结构体

(1) 结构体嵌套

(2) 结构体的"继承"
通过嵌套匿名结构体实现继承

(3) 结构体字段的可见性
结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）

(4) 结构体与JSON序列化



*/

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

type Student struct {
	ID     int
	Gender string
	Name   string
}

type Class struct {
	Title    string
	Students []*Student
}

func structTest2() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}

	cc := Circle{
		Point: Point{
			X: 5,
			Y: 5,
		},
		Radius: 10,
	}
	fmt.Printf("%#v\n", cc)

	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "male",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}

	// JSON序列化 ： 结构体 -> JSON格式字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
	} else {
		fmt.Printf("json:%s\n", data)
	}

	// JSON反序列化
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed")
	} else {
		fmt.Printf("%#v\n", c1)
	}

	var w Wheel
	w.X = 10
	w.Y = 10
	w.Radius = 5
	w.Spokes = 20
	fmt.Printf("%#v\n", w)

	w.Circle.Radius = 20
	w.Circle.X = 15
	w.Circle.Point.Y = 15
	fmt.Printf("%#v\n", w)

	pp := person{
		name: "dd",
		age:  16,
		city: "beijing",
	}
	fmt.Println(pp)
}
