package datastore

import (
	"google.golang.org/appengine/datastore"
)

type Keys struct {
	Key       *datastore.Key `datastore:"-" json:"keyString"`
	ParentKey *datastore.Key `datastore:"-" json:"parentKeyString,omitempty"`
}
