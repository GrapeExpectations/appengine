package datastore

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"testing"
)

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

	e := TestEntity{
		Name: "test",
	}

	err = Put(ctx, &e)
	if err != nil {
		t.Errorf("err: [%v]", err)
	}

	var entity TestEntity
	err = Get(ctx, e.GetKey(), &entity)
	if err != nil {
		t.Errorf("error [%v]", err)
	}
}
