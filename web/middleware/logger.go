package middleware

import (
	"net/http"

	"github.com/msyrus/hello-go/log"
)

// Logger returns a request logging middleware
func Logger(lgr log.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				lgr.Println(r.Method, r.RequestURI)
			}()

			next.ServeHTTP(w, r)
		})
	}
}
