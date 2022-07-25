package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go Monitor(ctx)

	time.Sleep(20 * time.Second)
}

func Monitor(ctx context.Context) {
	times := 1
	for {
		fmt.Println("monitor", times)
		times++
		time.Sleep(1 * time.Second)
	}
}
