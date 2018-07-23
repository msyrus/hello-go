package middleware

import (
	"net"
	"net/http"
	"runtime/debug"
)

// Recover middleware recover panic from API handler
// This should be the first middleware in middleware stack
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				switch err := err.(type) {
				case *net.OpError:
					break
				case error:
					debug.PrintStack()
					http.Error(w, err.Error(), http.StatusInternalServerError)
				case string:
					debug.PrintStack()
					http.Error(w, err, http.StatusInternalServerError)
				default:
					debug.PrintStack()
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}
