package request

import (
	"context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func NamespacedRequest(ns func(*http.Request) (string, error), fn func(context.Context, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

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

}
