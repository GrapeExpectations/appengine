package request

import (
  "context"
  "google.golang.org/appengine"
  "net/http"
)

type Handler struct {
  handler func(context.Context, http.ResponseWriter, *http.Request)
  path string
}

func NewHandler(path string, fn func(context.Context, http.ResponseWriter, *http.Request)) *Handler {
  return &Handler{
    handler: fn,
    path: path,
  }
}

func (h *Handler) Route() (string, func(http.ResponseWriter, *http.Request)) {
  return h.path, func(w http.ResponseWriter, r *http.Request) {
    ctx := appengine.NewContext(r)
    h.handler(ctx, w, r)
  }
}
