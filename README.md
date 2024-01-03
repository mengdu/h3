# H3

A http client for golang.

```go
func main() {
  client := h3.New()
	client.BaseURL = "https://httpbin.org"

  req := h3.NewRequest("POST", "/anything")
  // application/json
	form := h3.Json{}
	if err := form.Set(map[string]any{
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

	fmt.Println(res.Sprint())
}
```