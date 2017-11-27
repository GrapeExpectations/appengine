package request

import (
  "context"
  "net/http"
)

func (h *Handler) ContentType(contentType string) *Handler {
  fn := h.handler
  h.handler = func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", contentType)
    fn(ctx, w, r)
  }

  return h
}
