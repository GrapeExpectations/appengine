package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine"
)

type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request) error

type Handler struct {
	handler HandlerFunc
	path    string
}

type Logger interface {
	Debug(ctx context.Context, format string, args ...interface{})
	Error(ctx context.Context, format string, args ...interface{})
}

type defaultLogger struct{}

func (l *defaultLogger) Debug(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *defaultLogger) Error(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func NewHandler(path string, fn HandlerFunc) *Handler {
	return &Handler{
		handler: fn,
		path:    path,
	}
}

func (h *Handler) Route(logger Logger) (string, func(http.ResponseWriter, *http.Request)) {
	return h.path, func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		if logger == nil {
			logger = &defaultLogger{}
		}

		if err := h.handler(ctx, w, r); err != nil {
			switch err := err.(type) {
			case *errors.StatusError:
				code := err.GetCode()

				if code >= 400 && code < 500 {
					err.Log(ctx, func(s string) {
						logger.Debug(ctx, "%s\n", s)
					})
				} else {
					err.Log(ctx, func(s string) {
						logger.Error(ctx, "%s\n", s)
					})
				}

				http.Error(w, http.StatusText(code), code)
			default:
				logger.Error(ctx, "Error: %v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
	}
}
