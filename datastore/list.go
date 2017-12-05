package datastore

import (
	"github.com/GrapeExpectations/appengine/errors"
	"context"
	"google.golang.org/appengine/datastore"
	"net/http"
	"reflect"
)

func List(ctx context.Context, q *datastore.Query, entities interface{}) error {
	keys, err := q.GetAll(ctx, entities)
	if err != nil {
		return err
	}

	return setKeys(keys, entities)
}

func setKeys(keys []*datastore.Key, entities interface{}) error {
	slice := reflect.ValueOf(entities).Elem()
	if slice.Kind() != reflect.Slice {
		return errors.New(http.StatusBadRequest, "List requires slice")
	}

	elemType := slice.Type().Elem()
	iEntity := reflect.TypeOf((*Entity)(nil)).Elem()

	if !reflect.PtrTo(elemType).Implements(iEntity) {
		return errors.New(http.StatusBadRequest, "type does not implement Entity interface")
	}

	for i := 0; i < slice.Len(); i++ {
		if keys[i] != nil {
			e := slice.Index(i).Addr().Interface().(Entity)
			e.SetKey(keys[i])
			e.SetParentKey(keys[i].Parent())
		}
	}

	return nil
}
