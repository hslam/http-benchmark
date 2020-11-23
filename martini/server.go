package main

import (
	"flag"
	"github.com/go-martini/martini"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":8080", "-addr=:8080")
	flag.Parse()
}

func main() {
	m := Classic()
	m.Get("/", func() string {
		return "Hello World"
	})
	m.RunOnAddr(":8080")
}

func Classic() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return &martini.ClassicMartini{m, r}
}
