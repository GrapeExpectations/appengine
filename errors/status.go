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

func New(c int, m string, a ...interface{}) *ErrorStatus {
	return &ErrorStatus{c, fmt.Sprintf(m, a)}
}

func With(err error, c int, m string, a ...interface{}) *ErrorStatus {
	msg := fmt.Sprintf(m, a)
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
