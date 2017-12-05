package datastore

import (
	"github.com/GrapeExpectations/appengine/errors"
	"context"
	"google.golang.org/appengine/datastore"
	"net/http"
	"reflect"
)

func Get(ctx context.Context, key *datastore.Key, dst interface{}) error {
	if key == nil {
		return errors.New(http.StatusBadRequest, "key not provided")
	}

	iEntity := reflect.TypeOf((*Entity)(nil)).Elem()
	elemType := reflect.TypeOf(dst).Elem()

	if !reflect.PtrTo(elemType).Implements(iEntity) {
		return errors.New(http.StatusBadRequest, "type does not implement Entity interface")
	}

	if err := datastore.Get(ctx, key, dst); err != nil {
		return err
	}

	e := reflect.ValueOf(dst).Interface().(Entity)
	e.SetKey(key)
	e.SetParentKey(key.Parent())

	return nil
}
