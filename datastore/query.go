package datastore

import (
	"github.com/GrapeExpectations/appengine/errors"
	"context"
	"google.golang.org/appengine/datastore"
	"net/http"
	"reflect"
)

func Query(ctx context.Context, q *datastore.Query, entities interface{}, entity Entity) (*datastore.Cursor, error) {
	slice := reflect.ValueOf(entities).Elem()
	if slice.Kind() != reflect.Slice {
		return nil, errors.New(http.StatusBadRequest, "List requires slice")
	}

	t := q.Run(ctx)
	for {
		k, err := t.Next(entity)
		if err == datastore.Done {
			break // No further entities match the query.
		}
		if err != nil {
			return nil, err
		}
		entity.SetKey(k)
		entity.SetParentKey(k.Parent())
		slice.Set(reflect.Append(slice, reflect.Indirect(reflect.ValueOf(entity))))
	}

	cursor, err := t.Cursor()
	return &cursor, err
}
