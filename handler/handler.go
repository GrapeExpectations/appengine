package handler

import (
	"context"
	"net/http"

	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type Handler struct {
	handler func(context.Context, http.ResponseWriter, *http.Request) error
	path    string
}

func NewHandler(path string, fn func(context.Context, http.ResponseWriter, *http.Request) error) *Handler {
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
				if err.Msg != "" {
					log.Debugf(ctx, "%s", err)
				}
				http.Error(w, http.StatusText(err.Code), err.Code)
			default:
				log.Errorf(ctx, "Error: %v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
	}
}
