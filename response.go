package h3

import (
	"encoding/json"
	"io"
	"net/http"
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
