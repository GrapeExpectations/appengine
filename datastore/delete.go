package datastore

import (
	"context"
	"net/http"

	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine/datastore"
)

func Delete(ctx context.Context, key *datastore.Key, kind string) error {
	if key == nil {
		return errors.New(http.StatusBadRequest, errors.Message{Pkg: "datastore", Fn: "Delete", Msg: "error deleting, no key specified"})
	}
	if key.Kind() != kind {
		return errors.New(http.StatusBadRequest, errors.Message{Pkg: "datastore", Fn: "Delete", Msg: "key is not the specified type"})
	}

	return datastore.Delete(ctx, key)
}
