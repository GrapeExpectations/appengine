package handler

import (
	"context"
	"net/http"
)

func (h *Handler) Wrap(fn func(context.Context, http.ResponseWriter, *http.Request,
	func(context.Context, http.ResponseWriter, *http.Request))) *Handler {

	handlerFn := h.handler
  h.handler = func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		fn(ctx, w, r, handlerFn)
  }

  return h

}
