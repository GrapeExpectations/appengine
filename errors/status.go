package errors

import (
	"fmt"
)

type ErrorStatus struct {
	Code int
	Msg  string
}

func (e *ErrorStatus) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}

func New(c int, m string) *ErrorStatus {
	return &ErrorStatus{c, m}
}

func With(err error, c int, m string) *ErrorStatus {
	msg := m
	if m == "" && err != nil {
		msg = err.Error()
	}

	if serr, ok := err.(*ErrorStatus); ok {
		if serr.Code == 0 {
			serr.Code = c
		}
		if serr.Msg == "" {
			serr.Msg = msg
		}
		return serr
	}

	return New(c, msg)
}
