package h3

import (
	"fmt"
	"testing"
)

func TestBaseUrl(t *testing.T) {
	fmt.Println(buildUrl("https://localhost:8080/v1?a=1", nil, "/api/a/b?a=-1&b=123", nil))
}
