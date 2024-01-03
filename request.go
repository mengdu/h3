package h3

import (
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Method string
	Url    string
	Header http.Header
	Params url.Values
	Body   io.Reader
}

func NewRequest(method string, path string) *Request {
	return &Request{
		Method: method,
		Url:    path,
		Header: make(http.Header),
		Params: make(url.Values),
	}
}
