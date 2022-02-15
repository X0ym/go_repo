package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "Requested path is %q", ctx.Path())
	}

	// 路由
	router := fasthttprouter.New()
	router.GET("/", requestHandler)
	if err := fasthttp.ListenAndServe(":8001", router.Handler); err != nil {
		log.Fatalf("fasthttp ListenAndServe failed")
		return
	}
	//s := &fasthttp.Server{
	//	Handler: requestHandler,
	//	Name: "fasthttp server test",
	//}
	//if err := s.ListenAndServe("127.0.0.1:8002"); err != nil {
	//	log.Fatalf("error in ListenAndServe: %s", err)
	//}

}
