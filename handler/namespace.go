package handler

import (
	"context"
	"net/http"

	"github.com/GrapeExpectations/appengine/errors"

	"google.golang.org/appengine"
)

func (h *Handler) NamespacedRequest(ns func(*http.Request) (string, error)) *Handler {

	fn := h.handler
	h.handler = func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		namespace, err := ns(r)
		if err != nil {
			return errors.Wrap(err, "error getting namespace").
				SetCode(http.StatusBadRequest)
		}

		namespacedCtx, err := appengine.Namespace(ctx, namespace)
		if err != nil {
			return errors.Wrap(err, "error using namespace").
				SetCode(http.StatusInternalServerError)
		}

		return fn(namespacedCtx, w, r)
	}

	return h

}
