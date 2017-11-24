package request

import (
  "context"
  "fmt"
  "google.golang.org/appengine/aetest"
  "io/ioutil"
	"net/http"
	"net/http/httptest"
  "testing"
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
  fn := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, responseBody)
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

	if resp.StatusCode != 200 {
		t.Errorf("bad status code, got: %d, want: 200", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Failed to read response body")
	}
	if string(body) != responseBody {
		t.Errorf("incorrect response, got: %v, want: %v", string(body), responseBody)
	}
}
