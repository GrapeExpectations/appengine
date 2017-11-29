package datastore

import (
	"context"
	"errors"
	"github.com/GrapeExpectations/appengine/handler"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"reflect"
)

func Write(ctx context.Context, e Entity) error {
	ns := handler.NamespaceFromContext(ctx)
	entityType := reflect.TypeOf(e).Elem().Name()

	key := e.GetKey()
	parentKey := e.GetParentKey()

	if key == nil {
		key = datastore.NewIncompleteKey(ctx, entityType, e.GetParentKey())
		e.SetKey(key)
	}

	if key.Namespace() != ns {
		return errors.New("key does not belong to namespace")
	}
	if key.Kind() != entityType {
		return errors.New("key does not match kind")
	}
	if key.Parent() != parentKey {
		return errors.New("parent does not match key")
	}

	k, err := datastore.Put(ctx, key, e)
	if err != nil {
		log.Errorf(ctx, "could not write to datastore: %v", err)
		return err
	}
	e.SetKey(k)

	return nil
}
