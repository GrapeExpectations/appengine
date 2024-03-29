package datastore

import (
	"testing"

	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func TestPut_Error(t *testing.T) {
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

	err = Put(ctx, &e)
	if _, ok := err.(*errors.StatusError); !ok {
		t.Errorf("expected error, got none")
	}
}

func TestPut_Success(t *testing.T) {
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

	err = Put(ctx, &e)
	if err != nil {
		t.Errorf("err: [%v]", err)
	}
	if e.GetKey().Namespace() != namespace {
		t.Errorf("bad namespace, got: %s, want: %s", e.GetKey().Namespace(), namespace)
	}
}

func TestPutMulti_Error(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	key := datastore.NewIncompleteKey(ctx, "NoType", nil)
	e := TestEntity{
		Name: "Kenny",
	}
	e.SetKey(key)

	es := make([]Entity, 2)
	es[0] = &TestEntity{
		Name: "Jason",
	}
	es[1] = &e

	err = PutMulti(ctx, es)
	if _, ok := err.(*errors.StatusError); !ok {
		t.Errorf("expected error, got none")
	}
}

func TestPutMulti_Success(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	es := make([]Entity, 2)
	es[0] = &TestEntity{
		Name: "Jason",
	}

	es[1] = &TestEntity{
		Name: "Kenny",
	}

	err = PutMulti(ctx, es)
	if err != nil {
		t.Errorf("err: [%v]", err)
	}
}

func TestValidateKey_BadNamespace(t *testing.T) {
	const namespace = "local"
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	c, err := appengine.Namespace(ctx, namespace)
	if err != nil {
		t.Fatal(err)
	}

	key := datastore.NewIncompleteKey(c, TestEntityType, nil)
	e := TestEntity{
		Name: "test",
	}
	e.SetKey(key)

	_, err = validateKey(ctx, "", &e)
	if _, ok := err.(*errors.StatusError); !ok {
		t.Errorf("expected error, got none")
	}
}

func TestValidateKey_WrongKind(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	key := datastore.NewIncompleteKey(ctx, "Foo", nil)
	e := TestEntity{
		Name: "test",
	}
	e.SetKey(key)

	_, err = validateKey(ctx, "", &e)
	if _, ok := err.(*errors.StatusError); !ok {
		t.Errorf("expected error, got none")
	}
}

func TestValidateKey_WrongParent(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	key := datastore.NewIncompleteKey(ctx, TestEntityType, nil)
	parentKey := datastore.NewIncompleteKey(ctx, "Parent", nil)

	e := TestEntity{
		Name: "test",
	}
	e.SetKey(key)
	e.SetParentKey(parentKey)

	_, err = validateKey(ctx, "", &e)
	if _, ok := err.(*errors.StatusError); !ok {
		t.Errorf("expected error, got none")
	}
}

func TestValidateKey_Success(t *testing.T) {
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

	key, err := validateKey(ctx, namespace, &e)
	if err != nil {
		t.Errorf("error %v", err)
	}
	if key == nil {
		t.Errorf("expecting key, got none")
		t.Fatal("expecting key, got none")
	}
	if !key.Incomplete() {
		t.Errorf("expecting generated key, got: %v", key)
	}
}

func TestValidateKeys_Error(t *testing.T) {
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

	key := datastore.NewIncompleteKey(ctx, "NoType", nil)
	e := TestEntity{
		Name: "Kenny",
	}
	e.SetKey(key)

	es := make([]Entity, 2)
	es[0] = &TestEntity{
		Name: "Jason",
	}
	es[1] = &e

	_, err = validateKeys(ctx, es)
	if _, ok := err.(*errors.StatusError); !ok {
		t.Errorf("expected error, got none")
	}
}

func TestValidateKeys_Success(t *testing.T) {
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

	keys, err := validateKeys(ctx, es)
	if err != nil {
		t.Errorf("error %v", err)
	}
	if keys == nil {
		t.Errorf("expecting keys, got none")
		t.Fatal("expecting keys, got none")
	}
}
