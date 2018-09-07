package errors

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/appengine/log"
)

type StatusError struct {
	Code   int     `json:"code"`
	Detail Message `json:"detail"`
	err    error
}

func (e *StatusError) Error() string {
	return e.json()
}

func (e *StatusError) json() string {
	b, err := json.Marshal(e)
	if err != nil {
		return string(b)
	}

	return fmt.Sprintf("{code: %d, detail: {package: \"%s\", function:\"%s\", message: \"%s\"}}", e.Code, e.Detail.Pkg, e.Detail.Fn, e.Detail.Msg)
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
	log.Errorf(ctx, e.json())

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

func New(c int, m Message) *StatusError {
	return &StatusError{
		Code:   c,
		Detail: m,
		err:    errors.New(m.Msg),
	}
}

func Wrap(err error, m Message) *StatusError {
	return &StatusError{
		Code:   0,
		Detail: m,
		err:    err,
	}
}
