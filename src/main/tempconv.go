package main

import "fmt"

type Celsius float32
type Fahrenheit float32

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func tempTest() {
	c := Celsius(20)
	f := Fahrenheit(68)

	var1 := CToF(c)
	var2 := FToC(f)

	fmt.Println(var1, var2)
}
