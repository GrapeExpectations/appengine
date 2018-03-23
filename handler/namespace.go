package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GrapeExpectations/appengine/errors"

	"google.golang.org/appengine"
)

func (h *Handler) NamespacedRequest(ns func(*http.Request) (string, error)) *Handler {

	fn := h.handler
	h.handler = func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		namespace, err := ns(r)
		if err != nil {
			return errors.New(http.StatusBadRequest, fmt.Sprintf("error getting namespace: %v", err))
		}

		namespacedCtx, err := appengine.Namespace(ctx, namespace)
		if err != nil {
			return errors.New(http.StatusInternalServerError, fmt.Sprintf("error using namespace: %v", err))
		}

		return fn(namespacedCtx, w, r)
	}

	return h

}
