package datastore

import (
	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
	"testing"
)

func TestGet_BadType(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	type BadType struct {
		Name string
	}

	var entity BadType
	err = Get(ctx, nil, &entity)
	if _, ok := err.(*errors.ErrorStatus); !ok {
		t.Errorf("expected error, got none")
	}
}

func TestGet_NilKey(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	var entity TestEntity
	err = Get(ctx, nil, &entity)
	if _, ok := err.(*errors.ErrorStatus); !ok {
		t.Errorf("expected error, got none")
	}
}

func TestGet_Success(t *testing.T) {
	const namespace = "local"
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	ctx, err = appengine.Namespace(ctx, namespace)
	if err != nil {
		t.Fatal(err)
	}

	parentKey := datastore.NewKey(ctx, TestEntityType, "ParentEntityID", 0, nil)
	e := TestEntity{
		Name: "test",
	}
	e.SetParentKey(parentKey)

	err = Put(ctx, &e)
	if err != nil {
		t.Errorf("err: [%v]", err)
	}

	var entity TestEntity
	err = Get(ctx, e.GetKey(), &entity)
	if err != nil {
		t.Errorf("error [%v]", err)
	}
	if entity.GetKey() == nil {
		t.Errorf("entity key is nil")
	}
	if entity.GetParentKey() == nil {
		t.Errorf("entity parent key is nil")
	}
}
