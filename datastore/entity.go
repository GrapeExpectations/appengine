package datastore

import (
	"google.golang.org/appengine/datastore"
)

type Entity interface {
	GetKey() *datastore.Key
	GetParentKey() *datastore.Key
	SetKey(*datastore.Key)
	SetParentKey(*datastore.Key)
}
