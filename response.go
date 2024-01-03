package h3

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Response struct {
	*http.Response
	raw []byte
}

func (r *Response) Raw() ([]byte, error) {
	if r.raw != nil {
		return r.raw, nil
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	r.raw = body
	return body, nil
}

func (r Response) String() (string, error) {
	body, err := r.Raw()
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (r Response) Json(v interface{}) error {
	body, err := r.Raw()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func (r Response) Sprint() string {
	body, err := r.Raw()
	if err != nil {
		return err.Error()
	}
	headers := []string{}
	for k := range r.Request.Header {
		headers = append(headers, fmt.Sprintf("%s: %s", color(k, 1, 0), strings.Join(r.Request.Header.Values(k), ", ")))
	}

	headers2 := []string{}
	for k := range r.Header {
		headers2 = append(headers2, fmt.Sprintf("%s: %s", color(k, 1, 0), strings.Join(r.Header.Values(k), ", ")))
	}
	status := color(r.Status, 32, 0)
	if r.StatusCode != 200 {
		status = color(r.Status, 31, 0)
	}
	return fmt.Sprintf("%s %s\n%s\n\n====\n\n%s %s\n%s\n\n%s",
		color(r.Request.Method, 33, 0),
		r.Request.URL.String(),
		strings.Join(headers, "\n"),
		r.Proto,
		status,
		strings.Join(headers2, "\n"),
		string(body))
}

func color(str string, start int32, end int32) string {
	return fmt.Sprintf("\u001b[%dm%s\u001b[%dm", start, str, end)
}
