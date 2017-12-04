package datastore

import (
	"context"
	"errors"
	"google.golang.org/appengine/datastore"
)

func Delete(ctx context.Context, key *datastore.Key, kind string) error {
	if key == nil {
		return errors.New("no key specified")
	}
	if key.Kind() != kind {
		return errors.New("key is not the specified type")
	}

	return datastore.Delete(ctx, key)
}
