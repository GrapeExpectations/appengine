package handler

import (
	"context"
	"net/http"
)

func (h *Handler) ContentType(contentType string) *Handler {
	fn := h.handler
	h.handler = func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-type", contentType)
		return fn(ctx, w, r)
	}

	return h
}
