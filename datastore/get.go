package datastore

import (
	"context"
	"errors"
	"google.golang.org/appengine/datastore"
	"reflect"
)

func Get(ctx context.Context, key *datastore.Key, dst interface{}) error {
	iEntity := reflect.TypeOf((*Entity)(nil)).Elem()
	elemType := reflect.TypeOf(dst).Elem()

	if !reflect.PtrTo(elemType).Implements(iEntity) {
		return errors.New("type does not implement Entity interface")
	}

	if err := datastore.Get(ctx, key, dst); err != nil {
		return err
	}

	e := reflect.ValueOf(dst).Interface().(Entity)
	e.SetKey(key)
	e.SetParentKey(key.Parent())

	return nil
}
