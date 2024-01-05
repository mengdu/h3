package h3

import "io"

var DefaultClient = New()

func Req(method string, path string) *Request {
	return DefaultClient.Req(method, path)
}

func Get(path string) *Request {
	return DefaultClient.Req("GET", path)
}

func Post(path string, body io.Reader) *Request {
	req := DefaultClient.Req("POST", path)
	req.Body = body
	return req
}

func Put(path string, body io.Reader) *Request {
	req := DefaultClient.Req("PUT", path)
	req.Body = body
	return req
}

func Patch(path string, body io.Reader) *Request {
	req := DefaultClient.Req("PATCH", path)
	req.Body = body
	return req
}

func Delete(path string) *Request {
	return DefaultClient.Req("DELETE", path)
}

func Head(path string) *Request {
	return DefaultClient.Req("HEAD", path)
}
