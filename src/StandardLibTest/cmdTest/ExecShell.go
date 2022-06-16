package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "-c", "sh /Users/xieym/MyWorkspace/mycode/GolandProjects/go_awesomeProject/src/cmdTest/test.sh")
	bytes, err := cmd.Output()
	if err != nil {
		fmt.Println(cmd, "执行shell返回异常，", err)
	}
	cmd.Wait()
	fmt.Println("执行完成。输出：", string(bytes))
}
