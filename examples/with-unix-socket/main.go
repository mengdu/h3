package main

import (
	"net"
	"net/http"

	"github.com/mengdu/h3"
)

func main() {
	address := "./demo.sock"
	conn, err := net.Dial("unix", address)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return conn, nil
		},
	}

	client := h3.New()
	client.SetTransport(transport)
	client.BaseURL = "http://unix"
	client.SetDump(true)

	_, err = client.Req("GET", "/any?a=1").Do()
	if err != nil {
		panic(err)
	}
}
