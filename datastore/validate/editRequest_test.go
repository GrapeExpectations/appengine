package validate

import (
	"net/http"
	"testing"

	ds "github.com/GrapeExpectations/appengine/datastore"
	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func TestEditRequest(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	testKey := datastore.NewKey(ctx, TestEntityType, "EntityID", 0, nil)
	testEntity := &TestEntity{
		Name: "test",
	}
	testEntity.SetKey(testKey)

	t.Run("Invalid Method", func(t *testing.T) {
		err := EditRequest("GET", testKey.Encode(), TestEntityType, testEntity)
		if err == nil {
			t.Fatal("expected error, got none")
		}

		switch err := err.(type) {
		case *errors.StatusError:
			if err.Code != http.StatusMethodNotAllowed {
				t.Errorf("bad error code, got: %d, want: %d", err.Code, http.StatusMethodNotAllowed)
			}
		default:
			t.Errorf("bad error type, got: %T, want: StatusError", err)
		}
	})

	t.Run("POST Has Key", func(t *testing.T) {
		err := EditRequest("POST", testKey.Encode(), TestEntityType, testEntity)
		if err == nil {
			t.Fatal("expected error, got none")
		}

		switch err := err.(type) {
		case *errors.StatusError:
			if err.Code != http.StatusBadRequest {
				t.Errorf("bad error code, got: %d, want: %d", err.Code, http.StatusBadRequest)
			}
		default:
			t.Errorf("bad error type, got: %T, want: StatusError", err)
		}
	})

	t.Run("POST Success", func(t *testing.T) {
		testEntity := &TestEntity{
			Name: "test",
		}

		if err := EditRequest("POST", testKey.Encode(), TestEntityType, testEntity); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("PUT Key Mismatch", func(t *testing.T) {
		err := EditRequest("PUT", "keystring", TestEntityType, testEntity)
		if err == nil {
			t.Fatal("expected error, got none")
		}

		switch err := err.(type) {
		case *errors.StatusError:
			if err.Code != http.StatusBadRequest {
				t.Errorf("bad error code, got: %d, want: %d", err.Code, http.StatusBadRequest)
			}
		default:
			t.Errorf("bad error type, got: %T, want: StatusError", err)
		}
	})

	t.Run("PUT Wrong Key Kind", func(t *testing.T) {
		err := EditRequest("PUT", testKey.Encode(), "WrongKind", testEntity)
		if err == nil {
			t.Fatal("expected error, got none")
		}

		switch err := err.(type) {
		case *errors.StatusError:
			if err.Code != http.StatusBadRequest {
				t.Errorf("bad error code, got: %d, want: %d", err.Code, http.StatusBadRequest)
			}
		default:
			t.Errorf("bad error type, got: %T, want: StatusError", err)
		}
	})

	t.Run("PUT Success", func(t *testing.T) {
		if err := EditRequest("PUT", testKey.Encode(), TestEntityType, testEntity); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("PUT (no key) Wrong Key Kind", func(t *testing.T) {
		testEntity := &TestEntity{
			Name: "test",
		}

		err := EditRequest("PUT", testKey.Encode(), "WrongKind", testEntity)
		if err == nil {
			t.Fatal("expected error, got none")
		}

		switch err := err.(type) {
		case *errors.StatusError:
			code := err.GetCode()
			if code != http.StatusBadRequest {
				t.Errorf("bad error code, got: %d, want: %d", code, http.StatusBadRequest)
			}
		default:
			t.Errorf("bad error type, got: %T, want: StatusError", err)
		}
	})

	t.Run("PUT (no key) Success", func(t *testing.T) {
		testEntity := &TestEntity{
			Name: "test",
		}

		if err := EditRequest("PUT", testKey.Encode(), TestEntityType, testEntity); err != nil {
			t.Fatal(err)
		}
	})
}

const TestEntityType = "Test"

type TestEntity struct {
	ds.Keys
	Name string
}

func (e *TestEntity) GetKey() *datastore.Key {
	return e.Key
}

func (e *TestEntity) SetKey(k *datastore.Key) {
	e.Key = k
}

func (e *TestEntity) GetParentKey() *datastore.Key {
	return e.ParentKey
}

func (e *TestEntity) SetParentKey(k *datastore.Key) {
	e.ParentKey = k
}
