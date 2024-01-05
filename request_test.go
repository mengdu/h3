package h3

import (
	"net/url"
	"testing"
)

func TestBaseUrl(t *testing.T) {
	str, err := buildUrl("https://localhost:8080/v1?a=1", nil, "/api/a/b?a=-1&b=123", nil)
	if err != nil {
		t.Error(err)
	}
	expected := "https://localhost:8080/v1/api/a/b?a=-1&b=123"
	if str != expected {
		t.Errorf("expected: %s, got: %s", expected, str)
	}
}

func TestBaseUrlOverwrite(t *testing.T) {
	base := url.Values{}
	base.Add("a", "1")
	base.Add("b", "2")
	base.Add("c", "3")
	current := url.Values{}
	current.Add("b", "-22")
	current.Add("d", "4")
	current.Add("e", "5")
	current.Add("f", "6")
	str, err := buildUrl("https://localhost:8080/v1?a=0&g=7", base, "/api/a/b?a=-1&b=-2&h=8", current)
	if err != nil {
		t.Error(err)
	}
	expected := "https://localhost:8080/v1/api/a/b?a=-1&b=-22&c=3&d=4&e=5&f=6&g=7&h=8"
	if str != expected {
		t.Errorf("expected: %s, got: %s", expected, str)
	}
}
