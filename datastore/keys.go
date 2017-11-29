package datastore

import (
	"google.golang.org/appengine/datastore"
)

type Keys struct {
	Key       *datastore.Key `datastore:"-"`
	ParentKey *datastore.Key `datastore:"-"`
}
