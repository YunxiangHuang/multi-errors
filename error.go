package multierror

// Errors is an error list, it helps return multi-errors.
type Errors struct {
	errs      []error
	formatter func([]error) string
}

// Add adds error to inner if given error is not nil.
func (e *Errors) Add(err error) {
	if err == nil {
		return
	}

	switch newErr := err.(type) {
	case *Errors:
		e.errs = append(e.errs, newErr.errs...)
	default:
		e.errs = append(e.errs, err)
	}
}

// SetFormatter sets inner formatter.
// The formatter used to formatted the error string.
func (e *Errors) SetFormatter(f func([]error) string) {
	if f != nil {
		e.formatter = f
	}
}

// RawError returns inner errors as a slice.
func (e *Errors) RawError() []error {
	return e.errs
}

// Error implements error interface.
func (e *Errors) Error() string {
	if e.formatter != nil {
		return e.formatter(e.errs)
	}
	return defaultFormatter(e.errs)
}

// ErrorOrNil returns nil if no errors.
func (e *Errors) ErrorOrNil() error {
	if len(e.errs) == 0 {
		return nil
	}
	return e
}

// defaultFormatter combines errors and separated by ", ".
func defaultFormatter(es []error) string {
	var msg, firstFlag = "", true
	for _, err := range es {
		if !firstFlag {
			msg += ", "
		}
		msg += err.Error()
		firstFlag = false
	}
	return msg
}

// New creates an errors.
func New() *Errors {
	return &Errors{
		errs:      make([]error, 0),
		formatter: defaultFormatter,
	}
}

// NewWithError creates an errors and adds the given error.
func NewWithError(err error) *Errors {
	newErr := New()
	newErr.Add(err)
	return newErr
}
