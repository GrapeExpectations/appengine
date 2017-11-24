package request

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNamespacedRequest(t *testing.T) {
	const namespace = "local"

	// setup instance
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer inst.Close()

	// setup handler
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		n := datastore.NewKey(ctx, "E", "e", 0, nil).Namespace()
		fmt.Fprintf(w, n)
	}

	handleGood := NamespacedRequest(func(r *http.Request) (string, error) {
		return namespace, nil
	}, handler)
	handleError := NamespacedRequest(func(r *http.Request) (string, error) {
		return "", errors.New("no namespace")
	}, handler)

	// TEST 1: no namespace
	req1, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	w1 := httptest.NewRecorder()
	handleError(w1, req1)

	resp1 := w1.Result()

	if resp1.StatusCode != 400 {
		t.Errorf("bad status code, got: %d, want: 400", resp1.StatusCode)
	}

	// TEST 2: happy path
	req2, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	w2 := httptest.NewRecorder()
	handleGood(w2, req2)

	resp2 := w2.Result()

	if resp2.StatusCode != 200 {
		t.Errorf("bad status code, got: %d, want: 200", resp2.StatusCode)
	}

	body, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		t.Fatal("Failed to read response body")
	}
	if string(body) != namespace {
		t.Errorf("incorrect namespace, got: %v, want: %v", string(body), namespace)
	}
}
