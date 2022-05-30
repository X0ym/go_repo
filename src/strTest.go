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

}
