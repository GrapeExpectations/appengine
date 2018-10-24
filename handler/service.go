package handler

import (
	"context"
	"net/http"

	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine"
)

func (h *Handler) ServiceRequest() *Handler {
	fn := h.handler
	h.handler = func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		dev := appengine.IsDevAppServer()

		requestingAppID := r.Header.Get("X-Appengine-Inbound-Appid")
		appID := appengine.AppID(ctx)

		if !dev && requestingAppID != appID {
			return errors.New(http.StatusForbidden, "invalid service request")
		}

		return fn(ctx, w, r)
	}

	return h
}
