package error

import (
	"encoding/json"
	"fmt"
)

type Err struct {
	json.Marshaler
	error
	errorType Type
	code      string
	errors    []error
	message   string
	detail    map[string]interface{}
}

type Type string

const (
	Info  Type = "INFO"
	Warn  Type = "WARNING"
	Error Type = "ERROR"
)

func NewError(errType Type, code string) Err {
	return Err{
		errorType: errType,
		code:      code,
	}
}

func (e Err) WithError(err error) Err {
	errCast, ok := err.(Err)
	if ok {
		e.errors = append(e.errors, errCast)
	} else {
		e.errors = append(e.errors, Err{
			code:    "NATIVE_ERROR",
			message: err.Error(),
		})
	}
	return e
}

func (e Err) WithMessage(message string) Err {
	e.message = message
	return e
}

func (e Err) WithDetail(name string, detaildata interface{}) Err {
	e.detail[name] = detaildata
	return e
}

func (e Err) ToMap() map[string]interface{} {
	data := map[string]interface{}{
		"code":    e.code,
		"errors":  e.errors,
		"message": e.message,
		"detail":  e.detail,
	}
	return data
}

func (e Err) Error() string {
	return fmt.Sprintf("[%s] %s", e.code, e.message)
}

func (e Err) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.ToMap())
}
