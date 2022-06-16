package main

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	cmd := exec.Command("/bin/bash", "-c", "sh /Users/xieym/MyWorkspace/mycode/GolandProjects/go_awesomeProject/src/cmdTest/test.sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("获取stdout失败，", err)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println("执行Start。err:", err)
	}

	go func() {
		for {
			fmt.Println("等待context超时")
			select {
			case <-ctx.Done():
				fmt.Println("context 超时，kill 进程组")
				syscall.Kill(cmd.Process.Pid, syscall.SIGKILL)
				return
			}
		}
	}()

	fmt.Println("开始读取输出")
	for {
		tmp := make([]byte, 1024)
		n, err := stdOut.Read(tmp)
		fmt.Print(string(tmp[:n]))
		if err != nil {
			if err != io.EOF {
				fmt.Println("实时读取输出失败，", err)
				break
			} else {
				fmt.Println("运行结束")
				break
			}
		}
	}
}
