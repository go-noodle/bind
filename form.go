package bind

import (
	"net/http"

	"github.com/ajg/form"
	"github.com/go-noodle/noodle"
)

// Form constructs middleware that parses request form according to provided model
// and injects parsed object into context
func Form(model interface{}) noodle.Middleware {
	return Generic(model, formC)
}

func formC(r *http.Request) Decoder {
	return form.NewDecoder(r.Body)
}
