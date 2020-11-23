package main

import (
	"flag"
	"github.com/hslam/rum"
	"net/http"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":8080", "-addr=:8080")
	flag.Parse()
}

func main() {
	m := rum.New()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	m.Run(addr)
}
