package handler

import (
	"context"
	"net/http"
)

func (h *Handler) CORS() *Handler {
	fn := h.handler
	h.handler = func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Headers", "Authorization,Content-Type")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "OPTIONS,GET,PUT,POST,DELETE")
		if r.Method == "OPTIONS" {
			return
		}

		fn(ctx, w, r)
	}

	return h
}
