package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	//s := new(service.ServiceA)
	//rpc.Register(s)
	//listener, err := net.Listen("tcp", ":9091")
	//if err != nil {
	//	fmt.Println("监听失败。 err:", err)
	//}
	//for {
	//
	//}
	LogDirectory := "/opt/syslog/persistlog/bootloader"
	LOG_DIRECTORY_PREFIX := "/opt/syslog/persistlog"
	agentType := "HTTP"
	if agentType == "HTTP" {
		LogDirectory = filepath.Join(LOG_DIRECTORY_PREFIX, "httpmesh", "bootloader")
	} else {
		LogDirectory = filepath.Join(LOG_DIRECTORY_PREFIX, "scfmesh", "bootloader")
	}
	fmt.Println(LogDirectory)
}
