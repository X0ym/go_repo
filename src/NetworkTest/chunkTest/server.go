package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fmt.Println("Requested path is", string(ctx.Path()))
		body := ctx.Request.Body()
		fmt.Println(string(body))
		fmt.Fprintf(ctx, "resp body: Requested path is %q", string(ctx.Path()))
	}

	// 路由
	router := fasthttprouter.New()
	router.POST("/report", requestHandler)
	if err := fasthttp.ListenAndServe(":8001", router.Handler); err != nil {
		log.Fatalf("fasthttp ListenAndServe failed")
		return
	}
}
