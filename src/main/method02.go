package main

import "fmt"

func (u *User) getName() (name string) {
	name = u.Name
	return name
}

func (u *User) setName(newName string) {
	u.Name = newName
}

func methodTest02() {
	u := User{
		Name: "go",
		Age:  10,
		City: "beijing",
	}
	fmt.Println(u.getName())
	u.setName("shenzhen")
	fmt.Println(u.getName())
}
