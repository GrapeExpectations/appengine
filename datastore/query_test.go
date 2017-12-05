package datastore

import (
	"appengine/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
	"testing"
	"time"
)

func TestQuery_NotSlice(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	var entity TestEntity
	var entities TestEntity
	q := datastore.NewQuery(TestEntityType)

	_, err = Query(ctx, q, &entities, &entity)
	if _, ok := err.(*errors.ErrorStatus); !ok {
		t.Error("expected error, got none")
	}
}

func TestQuery_Success(t *testing.T) {
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

	e := TestEntity{
		Name: "Jason",
	}

	f := TestEntity{
		Name: "Kenny",
	}

	err = Put(ctx, &e)
	if err != nil {
		t.Fatal(err)
	}

	err = Put(ctx, &f)
	if err != nil {
		t.Fatal(err)
	}

	// wait for datastore
	time.Sleep(100 * time.Millisecond)

	var entity TestEntity
	entities := make([]TestEntity, 0)
	q := datastore.NewQuery(TestEntityType)

	cursor, err := Query(ctx, q, &entities, &entity)
	if err != nil {
		t.Errorf("query error: %v", err)
	}
	if cursor == nil {
		t.Errorf("expected cursor, got none")
	}
	if len(entities) != 2 {
		t.Errorf("expected two entities, got: %d [%v]", len(entities), entities)
	}
}
