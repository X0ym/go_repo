package main

import (
	"flag"
	"fmt"
	"os"
)

func main () {
	// os.Args
	// 命令行参数解析为字符串切片，第一个参数是执行文件名称
	if len(os.Args) > 0 {
		for i , org := range os.Args {
			fmt.Printf("参数 %d : %s\n", i, org)
		}
	}

	// flag
	// 使用方式一：flag.Xxx(name string, value string, usage string)
	str := flag.String("str", "strName", "flag string usage message")
	i := flag.Int("int", 10, "flag int usage message")
	j := flag.Int64("int64", 1001, "flag int64 usage message")
	k := flag.Bool("bool", false, "flag bool usage message")
	l := flag.Float64("float64", 100.1, "flag float64 usage message")
	duration := flag.Duration("duration", 1001, "flag duration usage message")

	// 方式二: flag.XxxVar(p *string, name string, value string, usage string)
	var strNew string
	flag.StringVar(&strNew, "strNew", "默认string值", "string flag 帮助信息")

	flag.Parse()

	fmt.Println("string类型: ", *str)
	fmt.Println("int类型: ", *i)
	fmt.Println("int64类型: ", *j)
	fmt.Println("bool类型: ", *k)
	fmt.Println("float64类型: ", *l)
	fmt.Println("Duration类型(int64): ", *duration)

	fmt.Println("strNew: ", strNew)

	// 支持的命令行参数形式
	/*
		-flag value
		--flag value
		-flag=value
		--flag=value
		其中需要特别注意的是 bool 类型的flag参数必须使用等号的方式指定

	 */

}
