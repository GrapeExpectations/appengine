package errors

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/appengine/log"
)

type StatusError struct {
	Code int
	Msg  string
	err  error
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("<%d> %s: %v", e.Code, e.Msg, e.err)
}

func (e *StatusError) GetCode() int {
	if e.Code > 0 {
		return e.Code
	}

	if e.err != nil {
		switch err := e.err.(type) {
		case *StatusError:
			return err.GetCode()
		}
	}

	return http.StatusInternalServerError
}

func (e *StatusError) Log(ctx context.Context) {
	log.Errorf(ctx, "{code: %d, message: \"%v\"}", e.Code, e.Msg)

	if e.err != nil {
		switch err := e.err.(type) {
		case *StatusError:
			err.Log(ctx)
		default:
			log.Errorf(ctx, "{cause: \"%v\"}", err)
		}
	}
}

func (e *StatusError) SetCode(c int) *StatusError {
	e.Code = c
	return e
}

func (e *StatusError) SetMsg(m string) *StatusError {
	e.Msg = m
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

func Wrap(m string, err error) *StatusError {
	return &StatusError{
		Code: 0,
		Msg:  m,
		err:  err,
	}
}
