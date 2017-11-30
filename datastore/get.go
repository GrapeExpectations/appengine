package datastore

import (
	"context"
	"google.golang.org/appengine/datastore"
)

func Get(ctx context.Context, key *datastore.Key, dst interface{}) error {
	return datastore.Get(ctx, key, dst)
}
