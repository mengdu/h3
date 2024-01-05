package main

import (
	"fmt"

	"github.com/mengdu/h3"
)

func main() {
	client := h3.New()
	client.BaseParams.Add("a", "1")
	client.BaseParams.Add("b", "2")

	req := client.Req("GET", "https://httpbin.org/get?c=3&d=4")
	req.Params.Add("e", "5")
	req.Params.Add("a", "6") // Overwrite base params
	req.Params.Add("d", "7") // Overwrite url params

	res, err := req.Do()
	if err != nil {
		panic(err)
	}
	fmt.Println(res.String())
}
