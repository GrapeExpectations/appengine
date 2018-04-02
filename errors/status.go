package errors

import (
	"errors"
	"fmt"
)

type StatusError struct {
	Code int
	Msg  string
	err  error
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("<%d> %s: %v", e.Code, e.Msg, e.err)
}

func (e *StatusError) SetCode(c int) *StatusError {
	if e.Code <= 0 {
		e.Code = c
	}

	return e
}

func (e *StatusError) SetMsg(m string) *StatusError {
	if m == "" || e.Msg == "" {
		e.Msg = m
	}

	return e
}

func New(c int, m string) *StatusError {
	return &StatusError{
		Code: c,
		Msg:  m,
		err:  errors.New(m),
	}
}

func With(err error) *StatusError {
	switch err := err.(type) {
	case *StatusError:
		return err
	default:
		return &StatusError{
			err: err,
		}
	}
}
