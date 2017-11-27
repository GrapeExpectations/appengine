package handler

import (
	"context"
	"fmt"
	"google.golang.org/appengine/aetest"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORS(t *testing.T) {
	const responseBody = "ok"

	// setup instance
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer inst.Close()

	// setup handler
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, responseBody)
	}

	goodHandler := NewHandler("/", handler).CORS()

	_, handleGood := goodHandler.Route()

	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	w := httptest.NewRecorder()
	handleGood(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("bad status code, got: %d, want: %d", resp.StatusCode, http.StatusOK)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Failed to read response body")
	}
	if string(body) != responseBody {
		t.Errorf("incorrect response body, got: %v, want: %v", string(body), responseBody)
	}

	acahHeader := resp.Header.Get("Access-Control-Allow-Headers")
	acacHeader := resp.Header.Get("Access-Control-Allow-Credentials")
	acaoHeader := resp.Header.Get("Access-Control-Allow-Origin")
	acamHeader := resp.Header.Get("Access-Control-Allow-Methods")

	acahExpect := "Authorization,Content-Type"
	acacExpect := "true"
	acaoExpect := "*"
	acamExpect := "OPTIONS,GET,PUT,POST,DELETE"

	if acahHeader != acahExpect {
		t.Errorf("bad Access-Control-Allow-Headers, got: %s, want: %s", acahHeader, acahExpect)
	}
	if acacHeader != acacExpect {
		t.Errorf("bad Access-Control-Allow-Credentials, got: %s, want: %s", acahHeader, acahExpect)
	}
	if acaoHeader != acaoExpect {
		t.Errorf("bad Access-Control-Allow-Origin, got: %s, want: %s", acahHeader, acahExpect)
	}
	if acamHeader != acamExpect {
		t.Errorf("bad Access-Control-Allow-Methods, got: %s, want: %s", acahHeader, acahExpect)
	}
}

func TestCORS_OPTIONS(t *testing.T) {
	const responseBody = "ok"
	const emptyResponse = ""

	// setup instance
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer inst.Close()

	// setup handler
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, responseBody)
	}

	optionsHandler := NewHandler("/", handler).CORS()

	_, handleOptions := optionsHandler.Route()

	req, err := inst.NewRequest("OPTIONS", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	w := httptest.NewRecorder()
	handleOptions(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("bad status code, got: %d, want: %d", resp.StatusCode, http.StatusOK)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Failed to read response body")
	}
	if string(body) != emptyResponse {
		t.Errorf("incorrect response body, got: %v, want: %v", string(body), emptyResponse)
	}

	acahHeader := resp.Header.Get("Access-Control-Allow-Headers")
	acacHeader := resp.Header.Get("Access-Control-Allow-Credentials")
	acaoHeader := resp.Header.Get("Access-Control-Allow-Origin")
	acamHeader := resp.Header.Get("Access-Control-Allow-Methods")

	acahExpect := "Authorization,Content-Type"
	acacExpect := "true"
	acaoExpect := "*"
	acamExpect := "OPTIONS,GET,PUT,POST,DELETE"

	if acahHeader != acahExpect {
		t.Errorf("bad Access-Control-Allow-Headers, got: %s, want: %s", acahHeader, acahExpect)
	}
	if acacHeader != acacExpect {
		t.Errorf("bad Access-Control-Allow-Credentials, got: %s, want: %s", acahHeader, acahExpect)
	}
	if acaoHeader != acaoExpect {
		t.Errorf("bad Access-Control-Allow-Origin, got: %s, want: %s", acahHeader, acahExpect)
	}
	if acamHeader != acamExpect {
		t.Errorf("bad Access-Control-Allow-Methods, got: %s, want: %s", acahHeader, acahExpect)
	}
}
