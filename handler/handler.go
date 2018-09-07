package handler

import (
	"context"
	"net/http"

	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request) error

type Handler struct {
	handler HandlerFunc
	path    string
}

func NewHandler(path string, fn HandlerFunc) *Handler {
	return &Handler{
		handler: fn,
		path:    path,
	}
}

func (h *Handler) Route() (string, func(http.ResponseWriter, *http.Request)) {
	return h.path, func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		if err := h.handler(ctx, w, r); err != nil {
			switch err := err.(type) {
			case *errors.StatusError:
				code := err.GetCode()

				if code >= 400 && code < 500 {
					err.Log(ctx, func(s string) {
						log.Debugf(ctx, "%s", s)
					})
				} else {
					err.Log(ctx, func(s string) {
						log.Errorf(ctx, "%s", s)
					})
				}

				http.Error(w, http.StatusText(code), code)
			default:
				log.Errorf(ctx, "Error: %v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
	}
}
