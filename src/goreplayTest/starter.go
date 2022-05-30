package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	gorPath := "/Users/xieym/Downloads/goreplay"
	end := " -exit-after 1m  -stats -output-http-stats -output-http-stats-ms 1000"
	//param := gorPath + " -input-raw" + " :8000" + " -output-http" + " http://127.0.0.1:8001"
	//param := gorPath + " --input-raw :8000 --output-http http://127.0.0.1:8001|5"

	// URL 过滤
	//param := gorPath + " -input-raw :8000 -output-http http://127.0.0.1:8001 -http-allow-url [0-9]+" + end
	//param := gorPath + " -input-raw :8000 -output-http http://127.0.0.1:8001 -http-disallow-url [0-9]+" + end

	// Header 过滤
	//param := gorPath + " -input-raw :8000 -output-http http://127.0.0.1:8001 --http-allow-header header_key:[0-9]+"
	//param := gorPath + " -input-raw :8000 -output-http http://127.0.0.1:8001 --http-disallow-header header_key:[0-9]+"

	// 请求方法过滤
	//param := gorPath + " -input-raw :8000 -output-http http://127.0.0.1:8001 --http-disallow-header header_key:[0-9]+ -http-allow-method GET -http-allow-method POST"

	// URL重写
	//param := gorPath + " -input-raw :8000 -output-http http://127.0.0.1:8001 -http-allow-url [0-9]+ -http-rewrite-url '/test([0-9]+):/test_new$1'"
	// " --input-raw :8000 --output-http http://127.0.0.1:8001 --http-allow-url [0-9]+ --http-rewrite-url \"/test([0-9]+):/test_new$1\""  注意不用 -- 不用双引号

	// Header 重写
	//param := gorPath + " -input-raw :8000 -output-http http://127.0.0.1:8001 -http-rewrite-header 'header_key:/test([0-9]+),/test_new$1'"

	// 添加 Header
	//param := gorPath + " -input-raw :8000 -output-http http://127.0.0.1:8001 -http-set-header header_key:value999"

	// 组合参数
	param := gorPath + " -input-raw :8000 -output-http http://127.0.0.1:8001 -http-allow-url [0-9]+ -http-allow-url [A-Z]" + end

	cmd := exec.Command("/bin/bash", "-c", param)
	// 设置子进程拥有独立的进程组id，即子进程的 pid
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	fmt.Println(cmd.String())
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("获取stdout失败，", err)
	}
	cmd.Stderr = cmd.Stdout

	// ========================= 执行命令 =========================
	cmd.Start() // Start() 方法不等待命令执行结束

	// ==================== 校验进程是否存在 ======================
	err = syscall.Kill(cmd.Process.Pid, 0)
	if err == nil {
		// 进程存在
		fmt.Println("进程存在，pid=", cmd.Process.Pid)
	} else {
		// 进程不存在

	}

	running, err := CheckProRunning("goreplay -input-raw :8000")
	if err != nil {
		fmt.Println("进程不存在：", err)
	}
	fmt.Println("进程pid:", running)

	pid, err := GetPid("goreplay -input-raw :8000")
	if err != nil {
		fmt.Println("进程不存在：", err)
	}
	fmt.Println("进程存在？\n", pid)
	// ===========================================================

	// ============= 读取输出 ==============
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

	cmd.Wait() // Wait() 方法 等待命令执行结束

	// 杀死进程 kill 进程组
	//defer func() {
	//	err = syscall.Kill(cmd.Process.Pid, syscall.SIGKILL)
	//	if err != nil {
	//		// 杀死进程失败
	//		fmt.Println("kill process failed. pid=", cmd.Process.Pid)
	//	}
	//}()

}

// CheckProRunning 根据进程名判断进程是否运行
func CheckProRunning(serverName string) (bool, error) {
	cmd := `ps ux | awk '/` + serverName + `/ && !/awk/ {print $2}'`
	pid, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("查询进程信息失败：", err)
		return false, err
	}
	// Output 中 Command 运行成功后返回pid 可能存在多个
	fmt.Println(string(pid))
	return string(pid) != "", err
}

// GetPid 根据进程名称获取进程ID 可能存在多个
func GetPid(serverName string) (string, error) {
	cmd := `ps ux | awk '/` + serverName + `/ && !/awk/ {print $2}'`
	pid, err := exec.Command("/bin/sh", "-c", cmd).Output()
	fmt.Println(string(pid))
	return strings.TrimSpace(string(pid)), err
}
