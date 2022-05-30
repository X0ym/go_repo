package utils

import (
	"fmt"
	"os"
)

var (
	path    = ""
	subtype = ""
)

func init() {
	os.Setenv("WCloud_Mesh_SubType", "1")
	//param1 := "echo export WCloud_Mesh_SubType=\"1\" >> /etc/profile"
	//cmd1 := exec.Command("/bin/bash", "-c", param1)
	//cmd1.Run()
	//param2 := "source /etc/profile"
	//cmd2 := exec.Command("/bin/bash", "-c", param2)
	//cmd2.Run()
	//cmd := exec.Command("/bin/bash", "-c", "source /Users/xieym/MyWorkspace/mycode/GolandProjects/go_awesomeProject/src/initTest/setEnv.sh")
	//output, err := cmd.Output()
	//if err != nil {
	//	fmt.Println("执行出错")
	//	fmt.Println(err)
	//}
	//cmd.Wait()
	//fmt.Println(string(output))

	//command := exec.Command("/bin/bash", "-c", "source /etc/profile")
	//command.Run()

	initEnv()
}

func initEnv() {
	path = os.Getenv("PATH")
	fmt.Println("path=", path)
	subtype = os.Getenv("WCloud_Mesh_SubType")
	if subtype == "" {
		fmt.Println("未设置成功")
	}
	fmt.Println("subType=", subtype)
}

func GetPath() string {
	return path
}

func GetSubType() string {
	return subtype
}
