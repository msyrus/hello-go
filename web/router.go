package web

import (
	"fmt"
	"net/http"

	"github.com/msyrus/hello-go/service"
	"github.com/msyrus/hello-go/web/resp"
)

// NewRouter returns new web Router
func NewRouter(gSvc *service.Greeting) http.Handler {
	mux := http.NewServeMux()

	gCtl := NewGreetController(gSvc)
	mux.Handle("/greetings", http.HandlerFunc(gCtl.GreetDefault))

	return mux
}

// ServeData serves data and meta with http status code 2xx
func ServeData(w http.ResponseWriter, r *http.Request, code int, data interface{}, meta *resp.Pager) {
	if code < 200 || code > 299 {
		panic(fmt.Errorf("serve data with %d", code))
	}
	re := resp.Response{
		Code: code,
		Data: data,
		Meta: meta,
	}
	resp.Render(w, r, re)
}
