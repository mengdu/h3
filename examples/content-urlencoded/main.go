package main

import (
	"fmt"

	"github.com/mengdu/h3"
)

func main() {
	client := h3.New()

	req := client.Req("POST", "https://httpbin.org/anything")
	req.Dump = true

	form := h3.Urlencoded{}
	form.Add("a", "123")
	form.Add("a", "456")
	form.Add("b", "789")
	req.Body = form.Form()

	res, err := req.Do()
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Status)

	result := map[string]interface{}{}
	if err := res.Json(&result); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", result["form"])
}
