package bind

// Option is applied to decoded value and decode error
type Option func(val interface{}, err error) error

// Validator is able to validate itself
type Validator interface {
	Validate() error
}

// Validate decoded value
func Validate() Option {
	return func(val interface{}, err error) error {
		if err != nil {
			return err
		}
		if v, ok := val.(Validator); ok {
			if err := v.Validate(); err != nil {
				return ValidationError{err}
			}
		}
		return nil
	}
}
