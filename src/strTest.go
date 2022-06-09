package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	env := os.Getenv("WCloud_Mesh_SubType")
	num, err := strconv.Atoi(env)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

	compare := strings.Compare("0", "unknown")
	fmt.Println("compare:", compare)

	b := false
	if b && func1() {
		fmt.Println("条件测试通过")
	}

}

func func1() bool {
	fmt.Println("函数执行")
	return false
}
