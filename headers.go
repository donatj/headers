package headers

import "net/http"

// Header represents the key and value of a given header
type Header struct {
	Key, Value string
}

type handler struct {
	h http.Handler

	headers []Header
}

// Handler provides simple header response injection middleware.
func Handler(h http.Handler, headers ...Header) http.Handler {
	return &handler{
		h:       h,
		headers: headers,
	}
}

func (xh *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, header := range xh.headers {
		w.Header().Set(header.Key, header.Value)
	}

	xh.h.ServeHTTP(w, r)
}
