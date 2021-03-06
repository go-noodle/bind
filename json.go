package bind

import (
	"encoding/json"
	"net/http"

	"github.com/go-noodle/noodle"
)

// JSON constructs middleware that parses request body according to provided model
// and injects parsed object into context
func JSON(model interface{}, opts ...Option) noodle.Middleware {
	return Generic(model, jsonC, opts...)
}

func jsonC(r *http.Request) Decoder {
	return json.NewDecoder(r.Body)
}
