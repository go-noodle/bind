package bind

// DecodeError wraps any error occured when decoding value
type DecodeError struct {
	error
}

// Cause satisfies errors.Causer interface
func (d *DecodeError) Cause() error {
	return d.error
}

// ValidationError wraps validation errors
type ValidationError struct {
	error
}

// Cause satisfies errors.Causer interface
func (d *ValidationError) Cause() error {
	return d.error
}
