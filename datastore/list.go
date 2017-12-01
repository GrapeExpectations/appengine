package datastore

import (
	"context"
	"errors"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"reflect"
)

func List(ctx context.Context, q *datastore.Query, entities interface{}) error {
	keys, err := q.GetAll(ctx, entities)
	if err != nil {
		log.Debugf(ctx, "error in List: %v", err)
		return err
	}

	return setKeys(keys, entities)
}

func setKeys(keys []*datastore.Key, entities interface{}) error {
	slice := reflect.ValueOf(entities).Elem()
	if slice.Kind() != reflect.Slice {
		return errors.New("List requires slice")
	}

	elemType := slice.Type().Elem()
	iEntity := reflect.TypeOf((*Entity)(nil)).Elem()

	if !reflect.PtrTo(elemType).Implements(iEntity) {
		return errors.New("type does not implement Entity interface")
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
