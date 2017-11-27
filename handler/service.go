package handler

import (
  "context"
  "google.golang.org/appengine"
  "net/http"
)

func (h *Handler) ServiceRequest() *Handler {
  fn := h.handler
  h.handler = func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
    dev := appengine.IsDevAppServer()

    requestingAppId := r.Header.Get("X-Appengine-Inbound-Appid")
    appId := appengine.AppID(ctx)

    if !dev && requestingAppId != appId {
      http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
    }

    fn(ctx, w, r)
  }

  return h
}
