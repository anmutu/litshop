package literr

import (
	"encoding/json"
	"litshop/src/lvm/local"
	"litshop/src/pkg/d"
)

type LitError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *LitError) Error() string {
	return e.Message
}

func (e *LitError) Json() ([]byte, error) {
	return json.Marshal(d.H{
		"code":    e.Code,
		"message": e.Message,
	})
}

func New(code int, message string) *LitError {
	return &LitError{
		Code:    code,
		Message: message,
	}
}

func NewWithCode(code int) *LitError {
	return &LitError{
		Code:    code,
		Message: local.Trans(trans(code)),
	}
}
