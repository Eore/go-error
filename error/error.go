package error

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type Error struct {
	ErrorType Type        `json:"error_type"`
	Code      string      `json:"code"`
	Err       string      `json:"error,omitempty"`
	Input     interface{} `json:"input,omitempty"`
	Output    interface{} `json:"output,omitempty"`
	Location  string      `json:"location,omitempty"`
}

type Type string

const (
	Info Type = "INFO"
	Warn Type = "WARNING"
	Err  Type = "ERROR"
)

func NewError(errType Type, code string) Error {
	return Error{
		ErrorType: errType,
		Code:      code,
	}
}

func (e Error) WithError(err error) Error {
	e.Err = err.Error()
	return e
}

func (e Error) WithInput(input interface{}) Error {
	e.Input = input
	return e
}

func (e Error) WithOutput(output interface{}) Error {
	e.Output = output
	return e
}

func (e Error) ToJSON() []byte {
	_, file, line, _ := runtime.Caller(1)
	e.Location = fmt.Sprintf("%s:%d", file, line)
	byteData, _ := json.Marshal(e)
	return byteData
}
