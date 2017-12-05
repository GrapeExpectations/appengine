package datastore

import (
	"context"
	"appengine/errors"
	"google.golang.org/appengine/datastore"
	"net/http"
)

func Delete(ctx context.Context, key *datastore.Key, kind string) error {
	if key == nil {
		return errors.New(http.StatusBadRequest, "error deleting: no key specified")
	}
	if key.Kind() != kind {
		return errors.New(http.StatusBadRequest, "key is not the specified type")
	}

	return datastore.Delete(ctx, key)
}
