package datastore

import (
  "google.golang.org/appengine/datastore"
)

type TestEntity struct {
  Keys
  Name string
}

func (e *TestEntity) GetKey() *datastore.Key {
  return e.Key
}

func (e *TestEntity) GetParentKey() *datastore.Key {
  return e.ParentKey
}

func (e *TestEntity) SetKey(k *datastore.Key) {
  e.Key = k
}

func (e *TestEntity) SetParentKey(k *datastore.Key) {
  e.ParentKey = k
}
