package datastore

import (
	"context"
	"net/http"

	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine/datastore"
)

type mockStore struct {
	delete   func(context.Context, *datastore.Key, string) error
	get      func(context.Context, *datastore.Key, interface{}) error
	list     func(context.Context, *datastore.Query, interface{}) error
	put      func(context.Context, Entity) error
	putMulti func(context.Context, []Entity) error
	query    func(context.Context, *datastore.Query, interface{}) (*datastore.Cursor, error)
}

func MockStore() *mockStore {
	s := mockStore{}

	s.MockDelete(func(ctx context.Context, key *datastore.Key, kind string) error {
		return errors.New(http.StatusNotImplemented, errors.Message{
			Pkg: "datastore",
			Fn:  "MockDelete",
			Msg: "test error",
		})
	})

	s.MockGet(func(ctx context.Context, key *datastore.Key, dst interface{}) error {
		return errors.New(http.StatusNotImplemented, errors.Message{
			Pkg: "datastore",
			Fn:  "MockGet",
			Msg: "test error",
		})
	})

	s.MockList(func(ctx context.Context, q *datastore.Query, entities interface{}) error {
		return errors.New(http.StatusNotImplemented, errors.Message{
			Pkg: "datastore",
			Fn:  "MockList",
			Msg: "test error",
		})
	})

	s.MockPut(func(ctx context.Context, e Entity) error {
		return errors.New(http.StatusNotImplemented, errors.Message{
			Pkg: "datastore",
			Fn:  "MockPut",
			Msg: "test error",
		})
	})

	s.MockPutMulti(func(ctx context.Context, entities []Entity) error {
		return errors.New(http.StatusNotImplemented, errors.Message{
			Pkg: "datastore",
			Fn:  "MockPutMulti",
			Msg: "test error",
		})
	})

	s.MockQuery(func(ctx context.Context, q *datastore.Query, entities interface{}) (*datastore.Cursor, error) {
		return nil, errors.New(http.StatusNotImplemented, errors.Message{
			Pkg: "datastore",
			Fn:  "MockQuery",
			Msg: "test error",
		})
	})

	return &s
}

func (s *mockStore) Delete(ctx context.Context, key *datastore.Key, kind string) error {
	return s.delete(ctx, key, kind)
}

func (s *mockStore) Get(ctx context.Context, key *datastore.Key, dst interface{}) error {
	return s.get(ctx, key, dst)
}

func (s *mockStore) List(ctx context.Context, q *datastore.Query, entities interface{}) error {
	return s.list(ctx, q, entities)
}

func (s *mockStore) Put(ctx context.Context, e Entity) error {
	return s.put(ctx, e)
}

func (s *mockStore) PutMulti(ctx context.Context, entities []Entity) error {
	return s.putMulti(ctx, entities)
}

func (s *mockStore) Query(ctx context.Context, q *datastore.Query, entities interface{}) (*datastore.Cursor, error) {
	return s.query(ctx, q, entities)
}

func (s *mockStore) MockDelete(fn func(context.Context, *datastore.Key, string) error) {
	s.delete = fn
}

func (s *mockStore) MockGet(fn func(context.Context, *datastore.Key, interface{}) error) {
	s.get = fn
}

func (s *mockStore) MockList(fn func(context.Context, *datastore.Query, interface{}) error) {
	s.list = fn
}

func (s *mockStore) MockPut(fn func(context.Context, Entity) error) {
	s.put = fn
}

func (s *mockStore) MockPutMulti(fn func(context.Context, []Entity) error) {
	s.putMulti = fn
}

func (s *mockStore) MockQuery(fn func(context.Context, *datastore.Query, interface{}) (*datastore.Cursor, error)) {
	s.query = fn
}
