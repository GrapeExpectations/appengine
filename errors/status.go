package errors

import (
	"fmt"
)

type ErrorStatus struct {
	Code int
  Msg string
}

func (e *ErrorStatus) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}

func New(c int, m string) *ErrorStatus {
	return &ErrorStatus{c,m}
}
