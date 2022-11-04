package apierror

import (
	"encoding/json"
	"errors"
	"fmt"
)

const errorDataField = "error"

func Convert(err error) *Error {
	if err == nil {
		return nil
	}

	var e *Error

	if !errors.As(err, &e) {
		e = NewUnknown(err, nil).(*Error)
	}

	return e
}

func Wrap(err error, code string, level Level, data map[string]any) error {
	if data == nil {
		data = make(map[string]any)
	}

	return &Error{code: code, level: level, data: data, error: err}
}

func New(code string, level Level, data map[string]any) error {
	if data == nil {
		data = make(map[string]any)
	}

	return &Error{code: code, level: level, data: data}
}

func NewUser(code string, data map[string]any) error {
	return New(code, LevelUser, data)
}

func NewSystem(code string, data map[string]any) error {
	return New(code, LevelSystem, data)
}

func NewUnknown(err error, data map[string]any) error {
	return Wrap(err, CodeUnknown, LevelSystem, data)
}

type Error struct {
	error error
	code  string
	level Level
	data  map[string]any
}

type errorJSON struct {
	Code  string         `json:"code"`
	Level Level          `json:"level"`
	Data  map[string]any `json:"data"`
}

func (e *Error) MarshalJSON() ([]byte, error) {
	data := e.data

	if e.error != nil {
		var err *Error
		if errors.As(e.error, &err) {
			data[errorDataField] = err
		} else {
			data[errorDataField] = e.error.Error()
		}
	}

	return json.Marshal(errorJSON{Code: e.code, Level: e.level, Data: data})
}

func (e *Error) UnmarshalJSON(data []byte) error {
	var err errorJSON

	if err := json.Unmarshal(data, &err); err != nil {
		return err
	}

	e.code = err.Code
	e.level = err.Level
	if err.Data == nil {
		err.Data = make(map[string]any)
	}
	e.data = err.Data

	return nil
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) Level() Level {
	return e.level
}

func (e *Error) Data() map[string]any {
	return e.data
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) Unwrap() error {
	return e.error
}

func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}
	return t.code == e.code && t.level == e.level
}

func (e *Error) String() string {
	marshaledErr, err := e.MarshalJSON()
	if err != nil {
		return fmt.Sprintf("error: %s, code: %s, level: %s, data: %v", e.error, e.code, e.level, e.data)
	}

	return string(marshaledErr)
}

func (e *Error) IsUser() bool {
	return e.level == LevelUser
}

func (e *Error) IsSystem() bool {
	return e.level == LevelSystem
}
