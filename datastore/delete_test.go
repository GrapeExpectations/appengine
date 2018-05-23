package datastore

import (
	"testing"

	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func TestDelete_WrongKind(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	key := datastore.NewKey(ctx, "EntityType", "EntityID", 0, nil)
	err = Delete(ctx, key, TestEntityType)
	if _, ok := err.(*errors.StatusError); !ok {
		t.Errorf("error expected, got none")
	}
}

func TestDelete_NilKey(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	err = Delete(ctx, nil, TestEntityType)
	if _, ok := err.(*errors.StatusError); !ok {
		t.Errorf("error expected, got none")
	}
}

func TestDelete_Success(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	key := datastore.NewKey(ctx, TestEntityType, "EntityID", 0, nil)
	err = Delete(ctx, key, TestEntityType)
	if err != nil {
		t.Errorf("error [%v]", err)
	}
}
