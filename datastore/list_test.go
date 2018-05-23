package datastore

import (
	"testing"
	"time"

	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func TestList_NotSlice(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	var entities TestEntity
	q := datastore.NewQuery(TestEntityType)

	err = List(ctx, q, &entities)
	if _, ok := err.(*errors.StatusError); !ok {
		t.Error("expected error, got none")
	}
}

func TestList_Success(t *testing.T) {
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

	es := make([]Entity, 2)
	es[0] = &TestEntity{
		Name: "Jason",
	}

	es[1] = &TestEntity{
		Name: "Kenny",
	}

	err = PutMulti(ctx, es)
	if err != nil {
		t.Fatal(err)
	}

	// wait for datastore
	time.Sleep(100 * time.Millisecond)

	entities := make([]TestEntity, 0)
	q := datastore.NewQuery(TestEntityType)

	err = List(ctx, q, &entities)
	if err != nil {
		t.Errorf("query error: %v", err)
	}
	if len(entities) != 2 {
		t.Errorf("expected two entities, got: %d [%v]", len(entities), entities)
	}
}

func TestSetKeys_NotSlice(t *testing.T) {
	var entities TestEntity
	err := setKeys(nil, &entities)
	if _, ok := err.(*errors.StatusError); !ok {
		t.Error("expected error, got none")
	}
}

func TestSetKeys_NotInterface(t *testing.T) {
	entities := make([]string, 0)
	err := setKeys(nil, &entities)
	if _, ok := err.(*errors.StatusError); !ok {
		t.Error("expected error, got none")
	}
}
