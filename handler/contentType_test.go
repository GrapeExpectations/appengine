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

func TestContentType(t *testing.T) {
	const contentType = "application/json"
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

	goodHandler := NewHandler("/", handler).ContentType(contentType)

	_, handleGood := goodHandler.Route()

	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	w := httptest.NewRecorder()
	handleGood(w, req)

	resp := w.Result()
	cType := resp.Header.Get("Content-Type")

	if resp.StatusCode != http.StatusOK {
		t.Errorf("bad status code, got: %d, want: %d", resp.StatusCode, http.StatusOK)
	}

	if cType != contentType {
		t.Errorf("bad content type, got: %s, want: %s", cType, contentType)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Failed to read response body")
	}
	if string(body) != responseBody {
		t.Errorf("incorrect response body, got: %v, want: %v", string(body), responseBody)
	}
}
