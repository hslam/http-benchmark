package main

import (
	"flag"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":8080", "-addr=:8080")
	flag.Parse()
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", func(ctx *fasthttp.RequestCtx) {
		ctx.Write([]byte("Hello World"))
	})
	fasthttp.ListenAndServe(addr, router.Handler)
}
