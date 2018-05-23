package handler

import (
	"context"
	"net/http"
)

func (h *Handler) Wrap(fn func(context.Context, http.ResponseWriter, *http.Request, HandlerFunc) error) *Handler {

	handlerFn := h.handler
	h.handler = func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		return fn(ctx, w, r, handlerFn)
	}

	return h

}
