package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestServiceRequest(t *testing.T) {
	const namespace = "local"
	const responseBody = "ok"

	// setup instance
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer inst.Close()

	// setup handler
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		fmt.Fprintf(w, responseBody)
		return nil
	}

	goodHandler := NewHandler("/", handler).ServiceRequest()
	errorHandler := NewHandler("/", handler).ServiceRequest()

	_, handleGood := goodHandler.Route(nil)
	_, handleError := errorHandler.Route(nil)

	// TEST 1: no namespace
	req1, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	w1 := httptest.NewRecorder()
	handleError(w1, req1)

	resp1 := w1.Result()

	if resp1.StatusCode != http.StatusForbidden {
		t.Errorf("bad status code, got: %d, want: %d", resp1.StatusCode, http.StatusForbidden)
	}

	// TEST 2: happy path
	req2, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	req2.Header.Add("X-Appengine-Inbound-Appid", "testapp")

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
	if string(body) != responseBody {
		t.Errorf("incorrect response body, got: %v, want: %v", string(body), responseBody)
	}
}
