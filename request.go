package h3

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const (
	userAgent = "h3/1.0 (github.com/mengdu/h3)"
)

type Request struct {
	Method string
	Url    string
	Header http.Header
	Params url.Values
	Body   io.Reader
	Client *Client
	Dump   bool
}

func NewRequest(method string, path string) *Request {
	return &Request{
		Method: method,
		Url:    path,
		Header: make(http.Header),
		Params: make(url.Values),
	}
}

func (r *Request) Do() (res *Response, err error) {
	c := r.Client
	if c == nil {
		return nil, errors.New("uninitialized client")
	}
	uri, err := buildUrl(c.BaseURL, c.BaseParams, r.Url, r.Params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(r.Method, uri, r.Body)
	if err != nil {
		return nil, err
	}

	request.Header = mergeHeaders(c.BaseHeader, r.Header)

	switch v := r.Body.(type) {
	case ContentBody:
		request.Header.Set("Content-Type", v.Type)
	}

	if request.Header.Get("User-Agent") == "" {
		request.Header.Set("User-Agent", userAgent)
	}

	if c.onBefore != nil {
		if err := c.onBefore(request); err != nil {
			return nil, err
		}
	}

	if r.Dump {
		dump, err := dumpRequest(request, true)
		if err != nil {
			return nil, err
		}
		fmt.Printf("%s\r\n", dump)
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	resp := &Response{
		Response: response,
	}

	if r.Dump {
		dump, err := dumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		fmt.Printf("%s", dump)
	}

	if c.onAfter != nil {
		if err := c.onAfter(resp); err != nil {
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

func buildUrl(base string, basePs url.Values, path string, currentPs url.Values) (string, error) {
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

func color(str string, start int32, end int32, enable bool) string {
	if !enable {
		return str
	}
	return fmt.Sprintf("\u001b[%dm%s\u001b[%dm", start, str, end)
}

func headerString(h http.Header, enableColor bool) string {
	keys := []string{}
	for k := range h {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	strs := []string{}
	for _, k := range keys {
		strs = append(strs, fmt.Sprintf("%s: %s", color(k, 1, 0, enableColor), strings.Join(h.Values(k), ", ")))
	}
	return strings.Join(strs, "\r\n")
}

func dumpRequest(req *http.Request, enableColor bool) (string, error) {
	body := ""

	if req.Body != nil && req.Body != http.NoBody && !strings.HasPrefix(req.Header.Get("Content-Type"), "multipart/form-data") {
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(req.Body); err != nil {
			return "", err
		}
		if err := req.Body.Close(); err != nil {
			return "", nil
		}
		// recover body steam
		req.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))
		body = buf.String()
	}
	strs := []string{
		fmt.Sprintf("%s %s %s",
			color(req.Method, 33, 0, enableColor),
			req.URL.String(),
			fmt.Sprintf(color("HTTP/%d.%d", 36, 0, enableColor), req.ProtoMajor, req.ProtoMinor),
		),
		headerString(req.Header, enableColor),
	}
	if body != "" {
		strs = append(strs, "\r\n"+body+"\r\n")
	} else {
		strs = append(strs, "\r\n")
	}
	return strings.Join(strs, "\r\n"), nil
}

func dumpResponse(res *Response, enableColor bool) (string, error) {
	status := color(res.Status, 32, 0, enableColor)
	if res.StatusCode != 200 {
		status = color(res.Status, 31, 0, enableColor)
	}
	body := ""
	if res.ContentLength != 0 && res.Body != nil && res.Body != http.NoBody {
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(res.Body); err != nil {
			return "", err
		}
		if err := res.Body.Close(); err != nil {
			return "", nil
		}
		// recover body steam
		res.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))
		body = buf.String()
	}
	return fmt.Sprintf("%s %s\r\n%s\r\n\r\n%s",
		color(res.Proto, 36, 0, enableColor),
		status,
		headerString(res.Header, enableColor),
		body,
	), nil
}
