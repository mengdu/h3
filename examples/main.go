package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/mengdu/h3"
)

func req() {
	// proxyURL, _ := url.Parse("http://127.0.0.1:7890")
	proxyURL, _ := url.Parse("socks5://127.0.0.1:7890")
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		// DisableCompression: true,
	}

	client := h3.New()
	client.BaseURL = "https://httpbin.org?a=1"
	client.BaseHeader.Add("x-tag", "test")
	client.BaseParams.Add("token", "123456")
	client.SetTransport(transport)
	client.SetTimeout(5 * time.Second)
	client.OnBefore(func(req *http.Request) error {
		fmt.Printf("> %s %s\n", req.Method, req.URL.String())
		return nil
	})
	// client.OnBefore(func(req *http.Request) error {
	// 	fmt.Println(2)
	// 	return nil
	// })
	// client.OnBefore(func(req *http.Request) error {
	// 	fmt.Println(3)
	// 	return nil
	// })
	client.OnAfter(func(res *h3.Response) error {
		fmt.Printf("< %s %s %s\n", res.Request.Method, res.Request.URL.String(), res.Status)
		return nil
	})

	req := client.Req("POST", "/anything?a=-1&b=123&c=456")
	// req := client.Req("POST", "/delay/6?a=-1&b=123&c=456")
	// req := client.Req("GET", "/image/png?a=-1&b=123&c=456")
	req.Dump = true
	// req := client.Req("GET", "https://api.github.com")
	req.Header.Add("x-test", "1234567")
	req.Header.Add("x-tag", "hello")
	req.Header.Add("x-tag", "123123")
	// req.Header.Add("User-Agent", "go-h3/1.0")
	req.Params.Add("b", "234")
	req.Params.Add("token", "654321")

	// application/json
	form := h3.Json{}
	if err := form.Set(map[string]interface{}{
		"a": 1,
		"b": []string{"1", "2", "3"},
	}); err != nil {
		panic(err)
	}
	req.Body = form.Form()

	// // application/x-www-form-urlencoded
	// form := h3.Urlencoded{}
	// form.Add("a", "234")
	// form.Set("b", "234")
	// req.Body = form.Form()

	// // form-data
	// form := h3.FormData{}
	// form.Add("k", "123")
	// form.Add("k", "456")
	// file, err := os.Open("./README.md")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// form.AddFile("file", "demo.txt", file)
	// req.Body = form.Form()

	res, err := req.Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Status, res.ContentLength)
	data := map[string]interface{}{}
	if err := res.Json(&data); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", data)
}

func main() {
	req()
	// u, err := url.Parse("https://localhost:3000/?a=1&b=2")
	// fmt.Println(err, u)
	// p, err := url.JoinPath("https://localhost:3000/v1?a=234", "/api", "/a/b/c")
	// fmt.Println(err, p)
	// fmt.Println(baseUrl("https://localhost:3000/v1?a=234", "/api/test?a=0&b=123&c=456&d=中文"))
}
