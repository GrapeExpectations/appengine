package errors

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"google.golang.org/appengine/log"
)

type StatusError struct {
	Code   int     `json:"code"`
	Detail message `json:"detail"`
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

	return fmt.Sprintf("{code: %d, detail: {function:\"%s\", file:\"%s\", line:%d, message: \"%s\"}}", e.Code, e.Detail.Fn, e.Detail.File, e.Detail.Line, e.Detail.Msg)
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

func (e *StatusError) Log(ctx context.Context, logFn func(string)) {
	logFn(e.json())

	if e.err != nil {
		switch err := e.err.(type) {
		case *StatusError:
			err.Log(ctx, logFn)
		default:
			log.Errorf(ctx, "{cause: \"%v\"}", err)
		}
	}
}

func (e *StatusError) SetCode(c int) *StatusError {
	e.Code = c
	return e
}

func New(c int, m string) *StatusError {
	detail := message{
		Fn:   "<unknown>",
		File: "<unknown>",
		Line: 0,
		Msg:  m,
	}

	if pc, file, line, ok := runtime.Caller(2); ok {
		detail.File = file
		detail.Line = line

		if fn := runtime.FuncForPC(pc); fn != nil {
			detail.Fn = fn.Name()
		}
	}

	return &StatusError{
		Code:   c,
		Detail: detail,
	}
}

func Wrap(err error, m string) *StatusError {
	detail := message{
		Fn:   "<unknown>",
		File: "<unknown>",
		Line: 0,
		Msg:  m,
	}

	if pc, file, line, ok := runtime.Caller(2); ok {
		detail.File = file
		detail.Line = line

		if fn := runtime.FuncForPC(pc); fn != nil {
			detail.Fn = fn.Name()
		}
	}

	return &StatusError{
		Code:   0,
		Detail: detail,
		err:    err,
	}
}
