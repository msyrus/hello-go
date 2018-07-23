package resp

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Response holds the api response object
type Response struct {
	Code   int         `json:"-"`
	Data   interface{} `json:"data,omitempty"`
	Errors []Error     `json:"errors,omitempty"`
	Meta   *Pager      `json:"meta,omitempty"`
}

// Render renders Response
func Render(w http.ResponseWriter, r *http.Request, re Response) {
	if re.Code == 0 {
		panic(errors.New("response status not defined"))
	}
	RenderJSON(w, re.Code, re)
}

// RenderJSON renders data into json response body
// it panic if if failes to encode json or write to response
func RenderJSON(w http.ResponseWriter, code int, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(body); err != nil {
		panic(err)
	}
}
