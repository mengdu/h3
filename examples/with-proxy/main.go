package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/mengdu/h3"
)

func main() {
	proxyURL, _ := url.Parse("socks5://127.0.0.1:7890")
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := h3.New()
	client.SetTransport(transport)

	res, err := client.Get("https://httpbin.org/get").Do()
	if err != nil {
		panic(err)
	}
	fmt.Println(res.String())
}
