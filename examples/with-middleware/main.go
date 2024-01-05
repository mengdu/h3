package main

import (
	"fmt"
	"net/http"

	"github.com/mengdu/h3"
)

func main() {
	client := h3.New()
	client.SetDump(true)
	client.OnBefore(func(req *http.Request) error {
		fmt.Printf("1: %s %s\n", req.Method, req.URL.String())
		return nil
	})
	client.OnBefore(func(req *http.Request) error {
		fmt.Printf("2: %s %s\n", req.Method, req.URL.String())
		return nil
	})
	client.OnBefore(func(req *http.Request) error {
		fmt.Printf("3: %s %s\n", req.Method, req.URL.String())
		return nil
	})

	client.OnAfter(func(res *h3.Response) error {
		fmt.Printf("1: %s\n", res.Status)
		return nil
	})
	client.OnAfter(func(res *h3.Response) error {
		fmt.Printf("2: %s\n", res.Status)
		return nil
	})
	client.OnAfter(func(res *h3.Response) error {
		fmt.Printf("3: %s\n", res.Status)
		return nil
	})

	_, err := client.Req("GET", "https://httpbin.org/get").Do()
	if err != nil {
		panic(err)
	}
}
