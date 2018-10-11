package datastore

import (
	"context"

	"google.golang.org/appengine/datastore"
)

type Store interface {
	Delete(context.Context, *datastore.Key, string) error
	Get(context.Context, *datastore.Key, interface{}) error
	List(context.Context, *datastore.Query, interface{}) error
	Put(context.Context, Entity) error
	PutMulti(context.Context, []Entity) error
	Query(context.Context, *datastore.Query, interface{}) (*datastore.Cursor, error)
}

type store struct{}

func NewStore() Store {
	return store{}
}

func (s store) Delete(ctx context.Context, key *datastore.Key, kind string) error {
	return Delete(ctx, key, kind)
}

func (s store) Get(ctx context.Context, key *datastore.Key, dst interface{}) error {
	return Get(ctx, key, dst)
}

func (s store) List(ctx context.Context, q *datastore.Query, entities interface{}) error {
	return List(ctx, q, entities)
}

func (s store) Put(ctx context.Context, e Entity) error {
	return Put(ctx, e)
}

func (s store) PutMulti(ctx context.Context, entities []Entity) error {
	return PutMulti(ctx, entities)
}

func (s store) Query(ctx context.Context, q *datastore.Query, entities interface{}) (*datastore.Cursor, error) {
	return Query(ctx, q, entities)
}
