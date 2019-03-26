package bind

import (
	"net/http"

	"github.com/go-noodle/noodle"
	"github.com/gorilla/schema"
)

// DefaultDecoder for URL parameters
var DefaultDecoder = schema.NewDecoder()

// Query constructs middleware that parses request query  according to provided model
// and injects parsed object into context
func Query(model interface{}, opts ...Option) noodle.Middleware {
	return Generic(model, schemaC, opts...)
}

type schemaDecoder http.Request

func (r *schemaDecoder) Decode(dest interface{}) error {
	return DefaultDecoder.Decode(dest, r.URL.Query())
}

func schemaC(r *http.Request) Decoder {
	return (*schemaDecoder)(r)
}
