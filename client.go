package h3

import (
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	BaseURL    string
	BaseHeader http.Header
	BaseParams url.Values
	client     *http.Client
	Before     func(req *http.Request) error
	After      func(res *Response) error
}

const (
	userAgent = "Golang H3/1.0"
)

func New() *Client {
	return &Client{
		BaseHeader: make(http.Header),
		BaseParams: make(url.Values),
		client:     &http.Client{},
	}
}

func (c *Client) Transport(transport http.RoundTripper) {
	c.client.Transport = transport
}

func (c *Client) Timeout(timeout time.Duration) {
	c.client.Timeout = timeout
}

func (c *Client) Jar(jar http.CookieJar) {
	c.client.Jar = jar
}

func (c *Client) Do(req *Request) (res *Response, err error) {
	uri, err := baseUrl(c.BaseURL, c.BaseParams, req.Url, req.Params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(req.Method, uri, req.Body)
	if err != nil {
		return nil, err
	}

	request.Header = mergeHeaders(c.BaseHeader, req.Header)

	switch v := req.Body.(type) {
	case ContentBody:
		request.Header.Set("Content-Type", v.Type)
	}

	if request.Header.Get("User-Agent") == "" {
		request.Header.Set("User-Agent", userAgent)
	}

	if c.Before != nil {
		if err := c.Before(request); err != nil {
			return nil, err
		}
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	resp := &Response{
		Response: response,
	}

	if c.After != nil {
		if err := c.After(resp); err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func mergeHeaders(headers ...http.Header) http.Header {
	h := make(http.Header)
	for _, header := range headers {
		for key, values := range header {
			if len(values) > 1 {
				h.Set(key, values[0])
				for i := 1; i < len(values); i++ {
					h.Add(key, values[i])
				}
			} else {
				h.Set(key, values[0])
			}
		}
	}

	return h
}

func baseUrl(base string, basePs url.Values, path string, currentPs url.Values) (string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	qs := url.Values{}
	for k := range basePs {
		qs[k] = basePs[k]
	}

	if u.Scheme == "" {
		u2, err := url.Parse(base)
		if err != nil {
			return "", err
		}
		u2 = u2.JoinPath(u.Path)
		q2 := u2.Query()
		for k := range q2 {
			qs[k] = q2[k]
		}

		q := u.Query()
		for k := range q {
			qs[k] = q[k]
		}
		for k := range currentPs {
			qs[k] = currentPs[k]
		}
		u2.RawQuery = qs.Encode()
		return u2.String(), nil
	} else {
		q := u.Query()
		for k := range q {
			qs[k] = q[k]
		}
		for k := range currentPs {
			qs[k] = currentPs[k]
		}
		u.RawQuery = qs.Encode()
		return u.String(), nil
	}
}
