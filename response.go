package go_response

import (
	"encoding/json"
	"net/http"
)

// Response ...
type Response struct {
	Vars     map[string]interface{}
	Var      interface{}
	HTTPCode int
}

// RenderJSON will write the current response in JSON form.
func (v *Response) RenderJSON(w http.ResponseWriter, r *http.Request) error {
	var b []byte
	var err error

	if v.Var == nil {
		b, err = json.Marshal(v.Vars)
	} else {
		b, err = json.Marshal(v.Var)
	}

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(v.HTTPCode)
	w.Write(b)

	return nil
}

// New creates a new Response instance and returns it.
func New() *Response {
	v := Response{}
	v.Vars = make(map[string]interface{})
	v.Var = nil
	v.HTTPCode = 200
	return &v
}
