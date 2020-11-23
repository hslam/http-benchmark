package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":8080", "-addr=:8080")
	flag.Parse()
}

func main() {
	gin.Default()
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})
	router.Run(addr)
}
