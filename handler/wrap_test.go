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

func TestWrap(t *testing.T) {
	const responseBody = "ok"
	const contentType = "application/json; charset=utf-8"

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

	wrapper := func(ctx context.Context, w http.ResponseWriter, r *http.Request, fn HandlerFunc) error {
		w.Header().Set("Content-type", contentType)
		return fn(ctx, w, r)
	}

	wrappedHandler := NewHandler("/", handler).Wrap(wrapper)

	_, handleWrap := wrappedHandler.Route()

	// TEST 1: no namespace
	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	w := httptest.NewRecorder()
	handleWrap(w, req)

	resp := w.Result()
	respContentType := resp.Header.Get("Content-Type")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Failed to read response body")
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("bad status code, got: %d, want: %d", resp.StatusCode, http.StatusOK)
	}
	if respContentType != contentType {
		t.Errorf("bad content type, got: %v, want: %v", respContentType, contentType)
	}
	if string(body) != responseBody {
		t.Errorf("incorrect response body, got: %v, want: %v", string(body), responseBody)
	}
}
