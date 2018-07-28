package web

import (
	"net/http"

	"github.com/msyrus/hello-go/service"
)

// GreetController holds necessery fields to serve greet handlers
type GreetController struct {
	svc *service.Greeting
}

// NewGreetController returns a new GreetController
func NewGreetController(svc *service.Greeting) *GreetController {
	return &GreetController{
		svc: svc,
	}
}

// GreetDefault serves default greetings
func (c *GreetController) GreetDefault(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	msg, _ := c.svc.GreetDefault(name)
	ServeData(w, r, http.StatusOK, msg, nil)
}
