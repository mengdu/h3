package main

import (
	"fmt"

	"github.com/mengdu/h3"
)

func main() {
	client := h3.New()
	// client.SetDump(true)

	req := client.Req("POST", "https://httpbin.org/anything")
	req.Dump = true

	form := h3.Json{}
	data := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": []interface{}{true, nil, "str", 10},
	}
	if err := form.Set(data); err != nil {
		panic(err)
	}
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
	fmt.Printf("%#v\n", result["json"])
}
