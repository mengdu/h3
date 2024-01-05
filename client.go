package h3

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	BaseURL    string
	BaseHeader http.Header
	BaseParams url.Values

	client     *http.Client
	onBefore   func(req *http.Request) error
	onAfter    func(res *Response) error
	enableDump bool
}

func New() *Client {
	return &Client{
		BaseHeader: make(http.Header),
		BaseParams: make(url.Values),
		client:     &http.Client{},
	}
}

func (c *Client) SetTransport(transport http.RoundTripper) {
	c.client.Transport = transport
}

func (c *Client) SetTimeout(timeout time.Duration) {
	c.client.Timeout = timeout
}

func (c *Client) SetJar(jar http.CookieJar) {
	c.client.Jar = jar
}

func (c *Client) SetDump(enable bool) {
	c.enableDump = enable
}

func (c *Client) OnBefore(fn func(req *http.Request) error) {
	if fn == nil {
		return
	}
	old := c.onBefore
	c.onBefore = func(req *http.Request) error {
		if old != nil {
			if err := old(req); err != nil {
				return err
			}
		}
		return fn(req)
	}
}

func (c *Client) OnAfter(fn func(res *Response) error) {
	if fn == nil {
		return
	}
	old := c.onAfter
	c.onAfter = func(res *Response) error {
		if old != nil {
			if err := old(res); err != nil {
				return err
			}
		}
		return fn(res)
	}
}

func (c *Client) Req(method string, path string) *Request {
	req := NewRequest(method, path)
	req.Client = c
	req.Dump = c.enableDump
	return req
}

func (c Client) Get(path string) *Request {
	return c.Req("GET", path)
}

func (c Client) Post(path string, body io.Reader) *Request {
	req := c.Req("POST", path)
	req.Body = body
	return req
}

func (c Client) Put(path string, body io.Reader) *Request {
	req := c.Req("PUT", path)
	req.Body = body
	return req
}

func (c Client) Patch(path string, body io.Reader) *Request {
	req := c.Req("PATCH", path)
	req.Body = body
	return req
}

func (c Client) Delete(path string) *Request {
	return c.Req("DELETE", path)
}

func (c Client) Head(path string) *Request {
	return c.Req("HEAD", path)
}
