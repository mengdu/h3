package main

import (
	"fmt"
	"os"

	"github.com/mengdu/h3"
)

func main() {
	client := h3.New()

	req := client.Req("POST", "https://httpbin.org/anything")
	req.Dump = true

	// form-data
	form := h3.FormData{}
	form.Add("a", "123")
	form.Add("b", "456")
	form.Add("c", "789")
	file, err := os.Open("./README.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	form.AddFile("file", "demo.txt", file)
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
