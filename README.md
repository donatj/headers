# headers

[![GoDoc](https://godoc.org/github.com/donatj/headers?status.svg)](https://godoc.org/github.com/donatj/headers)
[![CI](https://github.com/donatj/headers/actions/workflows/ci.yml/badge.svg)](https://github.com/donatj/headers/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/donatj/headers)](https://goreportcard.com/report/github.com/donatj/headers)

Simple Go middleware to append headers to responses.

# Example

This example adds an `X-Foo` and `X-Bar` header to all requests.

```golang
package main

import (
	"net/http"

	"github.com/donatj/headers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Hello World"`))
	})

	handler := headers.Handler(mux,
		headers.Header{Key: "X-Foo", Value: "Baz"},
		headers.Header{Key: "X-Bar", Value: "Qux"},
	)

	http.ListenAndServe(":80", handler)
}
```
