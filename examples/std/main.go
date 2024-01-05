package main

import (
	"github.com/mengdu/h3"
)

func main() {
	h3.DefaultClient.SetDump(true)

	_, err := h3.Get("https://httpbin.org/get").Do()
	if err != nil {
		panic(err)
	}

	_, err = h3.Post("https://httpbin.org/post", nil).Do()
	if err != nil {
		panic(err)
	}

	_, err = h3.Put("https://httpbin.org/put", nil).Do()
	if err != nil {
		panic(err)
	}

	_, err = h3.Patch("https://httpbin.org/patch", nil).Do()
	if err != nil {
		panic(err)
	}

	_, err = h3.Delete("https://httpbin.org/delete").Do()
	if err != nil {
		panic(err)
	}

	_, err = h3.Head("https://httpbin.org/head").Do()
	if err != nil {
		panic(err)
	}
}
