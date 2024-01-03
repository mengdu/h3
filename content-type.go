package h3

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"os"
)

const (
	json_type       = "application/json"
	urlencoded_type = "application/x-www-form-urlencoded"
)

type ContentBody struct {
	*bytes.Buffer
	Type string
}

func (c ContentBody) Read(p []byte) (n int, err error) {
	if c.Buffer == nil {
		return 0, nil
	}
	return c.Buffer.Read(p)
}

type Urlencoded struct {
	v url.Values
}

func (u *Urlencoded) init() {
	if u.v == nil {
		u.v = make(url.Values)
	}
}

func (u *Urlencoded) Add(key string, value string) {
	u.init()
	u.v.Add(key, value)
}

func (u *Urlencoded) Set(key string, value string) {
	u.init()
	u.v.Set(key, value)
}

func (u *Urlencoded) Get(key string) string {
	if u.v == nil {
		return ""
	}
	return u.v.Get(key)
}

func (u *Urlencoded) Del(key string) {
	if u.v != nil {
		u.v.Del(key)
	}
}

func (u Urlencoded) Form() ContentBody {
	if u.v == nil {
		return ContentBody{
			Type:   urlencoded_type,
			Buffer: bytes.NewBuffer([]byte("")),
		}
	}
	return ContentBody{
		Type:   urlencoded_type,
		Buffer: bytes.NewBuffer([]byte(u.v.Encode())),
	}
}

type Json struct {
	buf *bytes.Buffer
}

func (j *Json) Set(v interface{}) error {
	if s, ok := v.(string); ok {
		j.buf = bytes.NewBuffer([]byte(s))
		return nil
	} else if b, ok := v.([]byte); ok {
		j.buf = bytes.NewBuffer(b)
		return nil
	}

	buf, err := json.Marshal(v)
	if err != nil {
		return err
	}
	j.buf = bytes.NewBuffer(buf)
	return nil
}

func (j Json) Form() ContentBody {
	if j.buf == nil {
		return ContentBody{
			Type:   json_type,
			Buffer: bytes.NewBuffer([]byte("")),
		}
	}
	return ContentBody{
		Type:   json_type,
		Buffer: j.buf,
	}
}

type FormData struct {
	buf    *bytes.Buffer
	writer *multipart.Writer
}

func (f *FormData) init() {
	if f.writer == nil {
		f.buf = &bytes.Buffer{}
		f.writer = multipart.NewWriter(f.buf)
	}
}

func (f *FormData) Add(key string, value string) error {
	f.init()
	return f.writer.WriteField(key, value)
}

func (f *FormData) AddFile(key string, filename string, file *os.File) error {
	f.init()
	part, err := f.writer.CreateFormFile(key, filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	return err
}

func (f FormData) Form() ContentBody {
	f.init()
	f.writer.Close()
	return ContentBody{
		Type:   f.writer.FormDataContentType(),
		Buffer: f.buf,
	}
}
