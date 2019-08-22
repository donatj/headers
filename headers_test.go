package headers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInvalidHeaders(t *testing.T) {
	req, _ := http.NewRequest("GET", "localhost", bytes.NewReader([]byte{}))
	rec := httptest.NewRecorder()

	x := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Foo", "Bar")
		w.Write([]byte("baz"))
	})

	xhs := Handler(x, Header{Key: "Zoop", Value: "Zam"})
	xhs.ServeHTTP(rec, req)

	res := rec.Result()

	val, _ := res.Header["Zoop"]
	if val[0] != "Zam" {
		t.Errorf("expected header value '%v'; got '%v'", "Zam", val[0])
	}

	body, _ := ioutil.ReadAll(res.Body)
	sbody := strings.TrimSpace(string(body))
	if sbody != "baz" {
		t.Errorf("expected message '%v'; got '%v'", "baz", sbody)
	}
}
