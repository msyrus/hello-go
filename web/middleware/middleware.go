package middleware

import (
	"net/http"
)

// Middleware represents http handler middleware
type Middleware func(http.Handler) http.Handler

// Group groups middleware and returns a middlewarethat calls
// the middlewares in sequences they have been passed in parameters
func Group(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}
