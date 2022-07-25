package main

import (
	"context"
	"os/exec"
)

func main() {
	cmd := exec.CommandContext(context.Background(), "/bin/bash", "-c", "echo test")
	err := cmd.Run()
	if err != nil {
		return
	}
}
