package validate

import (
	"net/http"
	"testing"

	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

const kind = "KeyKind"

func TestKey(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	t.Run("Nil Key", func(t *testing.T) {
		err := Key(nil, kind)
		if err == nil {
			t.Fatal("expected error, got none")
		}

		switch e := err.(type) {
		case *errors.StatusError:
			if e.Code != http.StatusBadRequest {
				t.Errorf("bad error code, got: %d, want: %d", e.Code, http.StatusBadRequest)
			}
		default:
			t.Errorf("bad error type, got: %T, want: StatusError", e)
		}
	})

	t.Run("Wrong Kind", func(t *testing.T) {
		testKey := datastore.NewKey(ctx, kind, "KeyID", 0, nil)

		err := Key(testKey, "WrongKind")
		if err == nil {
			t.Fatal("expected error, got none")
		}

		switch e := err.(type) {
		case *errors.StatusError:
			if e.Code != http.StatusBadRequest {
				t.Errorf("bad error code, got: %d, want: %d", e.Code, http.StatusBadRequest)
			}
		default:
			t.Errorf("bad error type, got: %T, want: StatusError", e)
		}
	})

	t.Run("Any Kind", func(t *testing.T) {
		testKey := datastore.NewKey(ctx, kind, "KeyID", 0, nil)
		if err := Key(testKey, ""); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Success", func(t *testing.T) {
		testKey := datastore.NewKey(ctx, kind, "KeyID", 0, nil)
		if err := Key(testKey, kind); err != nil {
			t.Fatal(err)
		}
	})
}

func TestKeyString(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	t.Run("Invalid Key", func(t *testing.T) {
		_, err := KeyString("notakey", kind)
		if err == nil {
			t.Error("expected error, got none")
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

	t.Run("Wrong Kind", func(t *testing.T) {
		testKey := datastore.NewKey(ctx, kind, "KeyID", 0, nil)

		_, err := KeyString(testKey.Encode(), "WrongKind")
		if err == nil {
			t.Error("expected error, got none")
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

	t.Run("Any Kind", func(t *testing.T) {
		testKey := datastore.NewKey(ctx, kind, "KeyID", 0, nil)
		key, err := KeyString(testKey.Encode(), "")
		if err != nil {
			t.Fatal(err)
		}

		if !key.Equal(testKey) {
			t.Error("expecting keys to match, got mismatch")
		}
	})

	t.Run("Success", func(t *testing.T) {
		testKey := datastore.NewKey(ctx, kind, "KeyID", 0, nil)
		key, err := KeyString(testKey.Encode(), kind)
		if err != nil {
			t.Fatal(err)
		}

		if !key.Equal(testKey) {
			t.Error("expecting keys to match, got mismatch")
		}
	})
}
