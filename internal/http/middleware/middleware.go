// internal/http/middleware/route.go
package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

func HandleFunc(mux *http.ServeMux, pattern string, fn http.HandlerFunc, mws ...Middleware) {
	var h http.Handler = fn
	if len(mws) > 0 {
		h = Chain(h, mws...)
	}
	mux.Handle(pattern, h)
}
