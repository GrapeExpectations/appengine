package request

import (
	"context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"net/http"
)

func NamespaceFromContext(ctx context.Context) string {
	n := datastore.NewKey(ctx, "E", "e", 0, nil).Namespace()
	return n
}

func (h *Handler) NamespacedRequest(ns func(*http.Request) (string, error)) *Handler {

  fn := h.handler
  h.handler = func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
    namespace, err := ns(r)
		if err != nil {
			log.Errorf(ctx, "error getting namespace: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		ctx, err = appengine.Namespace(ctx, namespace)
		if err != nil {
			log.Errorf(ctx, "error using namespace: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		fn(ctx, w, r)
  }

  return h

}
