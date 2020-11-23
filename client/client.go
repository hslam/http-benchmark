package main

import (
	"flag"
	"github.com/hslam/stats"
	"io"
	"io/ioutil"
	"net/http"
)

var addr string
var clients int
var total int
var parallel int
var bar bool

func init() {
	flag.StringVar(&addr, "addr", ":8080", "-addr=:8080")
	flag.IntVar(&total, "total", 100000, "-total=100000")
	flag.IntVar(&clients, "clients", 1, "-clients=1")
	flag.IntVar(&parallel, "parallel", 1, "-parallel=1")
	flag.BoolVar(&bar, "bar", true, "bar: -bar=true")
	flag.Parse()
	stats.SetBar(bar)
}
func main() {
	if clients < 1 || total < 1 {
		return
	}
	var wrkClients = make([]stats.Client, clients)
	for i := 0; i < clients; i++ {
		var conn = &WrkClient{}
		conn.client = &http.Client{
			Transport: &http.Transport{
				MaxConnsPerHost: parallel + 1,
			},
		}
		conn.url = "http://" + addr
		conn.method = "GET"
		wrkClients[i] = conn
	}
	stats.StartPrint(parallel, total, wrkClients)
}

type WrkClient struct {
	client *http.Client
	url    string
	method string
}

func (c *WrkClient) Call() (int64, int64, bool) {
	req, _ := http.NewRequest(c.method, c.url, nil)
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, 0, false
	}
	written, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return 0, 0, false
	}
	return 0, written, true
}
