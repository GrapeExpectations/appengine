package datastore

import (
	"context"
	"errors"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"reflect"
)

func Query(ctx context.Context, q *datastore.Query, entities interface{}, entity Entity) (*datastore.Cursor, error) {
	slice := reflect.ValueOf(entities).Elem()
	if slice.Kind() != reflect.Slice {
		return nil, errors.New("List requires slice")
	}

	t := q.Run(ctx)
	for {
		k, err := t.Next(entity)
		if err == datastore.Done {
			break // No further entities match the query.
		}
		if err != nil {
			log.Debugf(ctx, "error fetching next entity: %v", err)
			return nil, err
		}
		entity.SetKey(k)
		entity.SetParentKey(k.Parent())
		slice.Set(reflect.Append(slice, reflect.Indirect(reflect.ValueOf(entity))))
	}

	cursor, err := t.Cursor()
	if err != nil {
		return nil, err
	}

	return &cursor, nil
}
