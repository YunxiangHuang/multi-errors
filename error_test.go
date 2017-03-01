package multierror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWithError(t *testing.T) {
	singleError := errors.New("1")
	mError := New()
	mError.Add(singleError)

	// New with single error.
	dstError := NewWithError(singleError)
	assert.Equal(t, 1, len(dstError.errs))

	// New with multi-error.
	dstError = NewWithError(mError)
	assert.Equal(t, mError.errs, dstError.errs)

	dstError.Add(errors.New("new error"))
	assert.NotEqual(t, mError.errs, dstError.errs)
}

func TestDefaultFormatter(t *testing.T) {
	mError := New()
	mError.Add(errors.New("1"))
	mError.Add(errors.New("2"))

	expect := `1, 2`
	assert.Equal(t, expect, mError.Error())
}

func TestSetterFormatter(t *testing.T) {
	mError := New()
	mError.Add(errors.New("1"))
	mError.Add(errors.New("2"))

	fmt := func([]error) string {
		return "ERROR"
	}
	mError.SetFormatter(fmt)
	assert.Equal(t, "ERROR", mError.Error())
}

func TestErrorsErrorOrNil(t *testing.T) {
	mError := New()
	assert.Nil(t, mError.ErrorOrNil())

	mError.Add(errors.New("error"))
	assert.NotNil(t, mError.ErrorOrNil())
}

func TestErrorsAdd(t *testing.T) {
	mError := New()
	assert.Empty(t, mError.errs)

	mError.Add(errors.New("error"))
	assert.NotEmpty(t, mError.errs)
	assert.Equal(t, "error", mError.Error())
}

func TestErrorsAddNil(t *testing.T) {
	mError := New()
	assert.Empty(t, mError.errs)

	mError.Add(nil)
	assert.Empty(t, mError.errs)
}

func TestErrorsRawError(t *testing.T) {
	mError := New()
	assert.Equal(t, mError.errs, mError.RawError())

	mError.Add(errors.New("error"))
	assert.Equal(t, mError.errs, mError.RawError())
}
