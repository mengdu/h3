package main

import (
	"time"

	"github.com/mengdu/h3"
)

func main() {
	client := h3.New()
	client.SetTimeout(5 * time.Second)
	client.SetDump(true)

	_, err := client.Get("https://httpbin.org/delay/8").Do()
	if err != nil {
		panic(err)
	}
}
