package datastore

import (
  "google.golang.org/appengine"
  "google.golang.org/appengine/aetest"
  "google.golang.org/appengine/datastore"
  "testing"
)

func TestWrite_BadNamespace(t *testing.T) {
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

  c, err := appengine.Namespace(ctx, "Foo")
  if err != nil {
    t.Fatal(err)
  }

  key := datastore.NewIncompleteKey(c, "TestEntity", nil)
  e := TestEntity{
    Name: "test",
  }
  e.SetKey(key)

  err = Write(ctx, &e)
  if err == nil {
    t.Errorf("expected error, got none")
  }
}

func TestWrite_WrongKind(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

  key := datastore.NewIncompleteKey(ctx, "Test", nil)
  e := TestEntity{
    Name: "test",
  }
  e.SetKey(key)

  err = Write(ctx, &e)
  if err == nil {
    t.Errorf("expected error, got none")
  }
}

func TestWrite_WrongParent(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

  key := datastore.NewIncompleteKey(ctx, "TestEntity", nil)
  parentKey := datastore.NewIncompleteKey(ctx, "Parent", nil)

  e := TestEntity{
    Name: "test",
  }
  e.SetKey(key)
  e.SetParentKey(parentKey)

  err = Write(ctx, &e)
  if err == nil {
    t.Errorf("expected error, got none")
  }
}

func TestWrite_Success(t *testing.T) {
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

  parent := datastore.NewKey(ctx, "Parent", "parentId", 0, nil)
  e := TestEntity{
    Name: "test",
  }
  e.SetParentKey(parent)

  err = Write(ctx, &e)
  if err != nil {
    t.Errorf("err: [%v]", err)
  }
  if e.GetKey().Namespace() != namespace {
    t.Errorf("bad namespace, got: %s, want: %s", e.GetKey().Namespace(), namespace)
  }
}
