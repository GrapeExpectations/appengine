package datastore

import (
	"context"
	"net/http"
	"reflect"

	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine/datastore"
)

func Query(ctx context.Context, q *datastore.Query, entities interface{}) (*datastore.Cursor, error) {
	slice := reflect.ValueOf(entities).Elem()
	if slice.Kind() != reflect.Slice {
		return nil, errors.New(http.StatusBadRequest, errors.Message{Pkg: "datastore", Fn: "Query", Msg: "List requires slice"})
	}

	entity := reflect.TypeOf((*Entity)(nil)).Elem()
	elemType := slice.Type().Elem()
	if !reflect.PtrTo(elemType).Implements(entity) {
		return nil, errors.New(http.StatusBadRequest, errors.Message{Pkg: "datastore", Fn: "Query", Msg: "Entity type required"})
	}

	t := q.Run(ctx)
	for {
		ev := reflect.New(elemType)
		k, err := t.Next(ev.Interface())
		if err == datastore.Done {
			break // No further entities match the query.
		}
		if err != nil {
			return nil, err
		}
		ev.MethodByName("SetKey").Call([]reflect.Value{reflect.ValueOf(k)})
		ev.MethodByName("SetParentKey").Call([]reflect.Value{reflect.ValueOf(k.Parent())})
		slice.Set(reflect.Append(slice, ev.Elem()))
	}

	cursor, err := t.Cursor()
	return &cursor, err
}
