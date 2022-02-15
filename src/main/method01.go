package main

import "fmt"

/**


 */

type Num int
type num int
type str string

func (s str) toString() {
	fmt.Println(s)
}

func (x num) method() {
	fmt.Println("private method:", x)
}

func (x num) Method() {
	fmt.Println("public method:", x)
}

type newNum struct {
	name string
	age  int
}

type Animal struct {
	name string
}

func NewAnimal() *Animal {
	return &Animal{}
}

func (a *Animal) SetName(name string) {
	a.name = name
}

type dog struct {
	Animal
	FeatureA string
}

func (d dog) Say() {
	fmt.Println("dog say", d.name)
}

func (n newNum) newMethod() {

}

func methodTest01() {
	var a num = 1
	a.method()

	p := NewAnimal()
	p.SetName("动物")

	d := dog{Animal: *p}
	d.Say()
	d.SetName("Dog")
	d.Say()

	var s str
	s = "biubiu"
	s.toString()
}
