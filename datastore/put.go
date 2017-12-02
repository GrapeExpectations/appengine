package datastore

import (
	"context"
	"errors"
	"github.com/GrapeExpectations/appengine/helper"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"reflect"
)

func Put(ctx context.Context, e Entity) error {
	ns := helper.NamespaceFromContext(ctx)
	key, err := validateKey(ctx, ns, e)
	if err != nil {
		return err
	}

	k, err := datastore.Put(ctx, key, e)
	if err != nil {
		log.Debugf(ctx, "could not write to datastore: %v", err)
		return err
	}
	e.SetKey(k)

	return nil
}

func PutMulti(ctx context.Context, entities []Entity) error {
	keys, err := validateKeys(ctx, entities)
	if err != nil {
		return err
	}

	ks, err := datastore.PutMulti(ctx, keys, entities)
	if err != nil {
		return err
	}

	for i, e := range entities {
		e.SetKey(ks[i])
	}

	return nil
}

func validateKey(ctx context.Context, ns string, e Entity) (*datastore.Key, error) {
	entityType := reflect.TypeOf(e).Elem().Name()

	key := e.GetKey()
	parentKey := e.GetParentKey()

	if key == nil {
		key = datastore.NewIncompleteKey(ctx, entityType, e.GetParentKey())
		e.SetKey(key)
	}

	if key.Namespace() != ns {
		return nil, errors.New("key does not belong to namespace")
	}
	if key.Kind() != entityType {
		return nil, errors.New("key does not match kind")
	}
	if key.Parent() != parentKey {
		return nil, errors.New("parent does not match key")
	}

	return key, nil
}

func validateKeys(ctx context.Context, entities []Entity) ([]*datastore.Key, error) {
	ns := helper.NamespaceFromContext(ctx)
	keys := make([]*datastore.Key, len(entities))
	for i, e := range entities {
		key, err := validateKey(ctx, ns, e)
		if err != nil {
			return nil, err
		}
		keys[i] = key
	}
	return keys, nil
}
