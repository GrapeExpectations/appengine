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

func TestHandler(t *testing.T) {
	// setup instance
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer inst.Close()

	responseBody := "ok"
	path := "/"
	fn := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		fmt.Fprintf(w, responseBody)
		return nil
	}

	handler := NewHandler(path, fn)

	p, f := handler.Route()
	if p != path {
		t.Errorf("incorrect path, got: %v, want: %v", p, path)
	}

	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	w := httptest.NewRecorder()
	f(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("bad status code, got: %d, want: %d", resp.StatusCode, http.StatusOK)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Failed to read response body")
	}
	if string(body) != responseBody {
		t.Errorf("incorrect response, got: %v, want: %v", string(body), responseBody)
	}
}
