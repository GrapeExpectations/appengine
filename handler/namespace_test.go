package request

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNamespaceFromContext(t *testing.T) {
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

	ns := NamespaceFromContext(ctx)
	if ns != namespace {
		t.Errorf("bad namespace, got: %s, want: %s", ns, namespace)
	}
}

func TestNamespacedRequest(t *testing.T) {
	const namespace = "local"
	const invalidNamespace = "not a valid namespace!*"

	// setup instance
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer inst.Close()

	// setup handler
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		n := NamespaceFromContext(ctx)
		fmt.Fprintf(w, n)
	}

	goodHandler := NewHandler("/", handler).
		NamespacedRequest(func(r *http.Request) (string, error) {
			return namespace, nil
		})

	errorHandler := NewHandler("/", handler).
		NamespacedRequest(func(r *http.Request) (string, error) {
			return "", errors.New("no namespace")
		})

	invalidHandler := NewHandler("/", handler).
		NamespacedRequest(func(r *http.Request) (string, error) {
			return invalidNamespace, nil
		})

	_, handleGood := goodHandler.Route()
	_, handleError := errorHandler.Route()
	_, handleInvalid := invalidHandler.Route()

	// TEST 1: no namespace
	req1, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	w1 := httptest.NewRecorder()
	handleError(w1, req1)

	resp1 := w1.Result()

	if resp1.StatusCode != http.StatusBadRequest {
		t.Errorf("bad status code, got: %d, want: %d", resp1.StatusCode, http.StatusBadRequest)
	}

	// TEST 2: happy path
	req2, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	w2 := httptest.NewRecorder()
	handleGood(w2, req2)

	resp2 := w2.Result()

	if resp2.StatusCode != http.StatusOK {
		t.Errorf("bad status code, got: %d, want: %d", resp2.StatusCode, http.StatusOK)
	}

	body, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		t.Fatal("Failed to read response body")
	}
	if string(body) != namespace {
		t.Errorf("incorrect namespace, got: %v, want: %v", string(body), namespace)
	}

	// TEST 3: invalid namespace name
	req3, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	w3 := httptest.NewRecorder()
	handleInvalid(w3, req3)

	resp3 := w3.Result()

	if resp3.StatusCode != http.StatusInternalServerError {
		t.Errorf("bad status code, got: %d, want: %d", resp3.StatusCode, http.StatusInternalServerError)
	}
}