package bind

import (
	"net/http"
	"reflect"

	"github.com/go-noodle/noodle"
)

type key struct{}

var bindKey key

type decoderResult struct {
	val interface{}
	err error
}

// Constructor is a generic function modelled after json.NewDecoder
type Constructor func(r *http.Request) Decoder

// Decoder populates target object with data from request body
type Decoder interface {
	Decode(interface{}) error
}

// Generic is a middleware factory for request binding.
// Accepts Constructor and returns binder for model
func Generic(model interface{}, dc Constructor, opts ...Option) noodle.Middleware {
	typeModel := reflect.TypeOf(model)
	if typeModel.Kind() == reflect.Ptr {
		typeModel = typeModel.Elem()
	}
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			val := reflect.New(typeModel).Interface()
			err := dc(r).Decode(val)
			if err != nil {
				err = DecodeError{err}
			}
			for _, opt := range opts {
				if err = opt(val, err); err != nil {
					break
				}
			}
			next(w, noodle.WithValue(r, bindKey, decoderResult{val, err}))
		}
	}
}

// GetData extracts data parsed from upstream Bind operation, discarding
// decoding error.  Deprecated in favor of Get() function.
func GetData(r *http.Request) (res interface{}) {
	res, _ = Get(r)
	return
}

// Get extracts data parsed from upstream Bind operation, along with the decode error.
func Get(r *http.Request) (val interface{}, err error) {
	res, ok := noodle.Value(r, bindKey).(decoderResult)
	if !ok {
		return
	}
	return res.val, res.err
}
