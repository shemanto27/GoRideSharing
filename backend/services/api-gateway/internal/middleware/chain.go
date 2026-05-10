package middleware

import "net/http"

// custom function type 
// Middleware = any function
// that takes handler
// and returns handler
type Middleware func(http.Handler) http.Handler

// Variadic Parameters (middlewares ...Middleware)
// Go internally converts into slice.
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {

	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}