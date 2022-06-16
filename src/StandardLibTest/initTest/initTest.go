package main

import (
	"fmt"
	"go_awesomeProject/src/StandardLibTest/initTest/utils"
	"os"
	"os/exec"
	"time"
)

func main() {

	path := utils.GetPath()
	fmt.Println(path)
	fmt.Println("subType=", utils.GetSubType())

	time.Sleep(1 * time.Second)
	go func() {
		fmt.Println("子进程获取环境变量")
		cmd := exec.Command("/bin/bash", "-c", "sh /Users/xieym/MyWorkspace/mycode/GolandProjects/go_awesomeProject/src/initTest/getEnv.sh")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("获取出现异常")
			fmt.Println(err)
		}
		fmt.Println("output=", string(output))

		getenv := os.Getenv("WCloud_Mesh_SubType")
		fmt.Println("子进程获取：", getenv)
	}()

	time.Sleep(5 * time.Second)
}
