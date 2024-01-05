package main

import (
	"fmt"

	"github.com/mengdu/h3"
)

func main() {
	client := h3.New()
	client.BaseHeader.Add("x-tag", "hello")
	client.BaseHeader.Add("x-tag", "test")
	client.BaseHeader.Add("x-tag", "demo")
	client.BaseHeader.Add("x-test", "123")

	req := client.Req("GET", "https://httpbin.org/get")
	req.Header.Add("x-test", "456") // Overwrite base header
	req.Header.Add("x-demo", "hello")

	res, err := req.Do()
	if err != nil {
		panic(err)
	}
	fmt.Println(res.String())
}
