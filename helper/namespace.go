package helper

import (
  "context"
  "google.golang.org/appengine/datastore"
)

func NamespaceFromContext(ctx context.Context) string {
	n := datastore.NewKey(ctx, "E", "e", 0, nil).Namespace()
	return n
}
