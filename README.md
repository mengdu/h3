# H3

A simple encapsulation of Golang HTTP client.

```go
package demo

import "fmt"

func main() {
	client := h3.New()
	client.BaseURL = "https://httpbin.org"

	req := client.Req("POST", "/anything")
	// application/json
	form := h3.Json{}
	if err := form.Set(map[string]interface{}{
		"a": 1,
		"b": []string{"1", "2", "3"},
	}); err != nil {
		panic(err)
	}
	req.Body = form.Form()

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{}
	if err := res.Json(&data); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", data)
}
```

[More examples](./examples/)
