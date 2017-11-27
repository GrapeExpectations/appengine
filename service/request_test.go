package service

import (
  "bytes"
  "net/http"
  "reflect"
  "strings"
  "testing"
)

func TestNewRequest(t *testing.T) {
  const path = "/test"
  emptyHeader := make(http.Header)

  req := NewRequest(path)

  if req.body != nil {
    t.Errorf("bad body, got: %v, want: %v", req.body, nil)
  }
  if !reflect.DeepEqual(req.Header, emptyHeader) {
    t.Errorf("bad Header, got: %v, want: %v", req.Header, emptyHeader)
  }
  if req.path != path {
    t.Errorf("bad path, got: %s, want: %s", req.path, path)
  }
}

func TestWith(t *testing.T) {
  const headerKey = "X-Foo"
  const headerVal = "Bar"

  fn := func(key, val string) func(r *Request) *Request {
    return func(r *Request) *Request {
      r.Header.Add(key, val)
      return r
    }
  }

  req := NewRequest("/").With(fn(headerKey, headerVal))

  header := req.Header.Get(headerKey)
  if header != headerVal {
    t.Errorf("bad Header, got: %v, want: %v", header, headerVal)
  }
}

func TestWithBody(t *testing.T) {
  const body = "ok"

  req := NewRequest("/").WithBody(strings.NewReader(body))

  if req.body == nil {
    t.Errorf("body is nil")
  }

  buf := new(bytes.Buffer)
  buf.ReadFrom(req.body)
  s := buf.String()
  if s != body {
    t.Errorf("bad body, got: %s, want: %s", s, body)
  }
}

func TestWithHeader(t *testing.T) {
  const headerKey = "X-Foo"
  const headerVal = "Bar"

  req := NewRequest("/").WithHeader(headerKey, headerVal)

  header := req.Header.Get(headerKey)
  if header != headerVal {
    t.Errorf("bad Header, got: %v, want: %v", header, headerVal)
  }
}
