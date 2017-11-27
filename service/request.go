package http

import (
  "io"
  "net/http"
)

type Request struct {
  body io.Reader
  Header http.Header
  path string
}

func NewRequest(path string) *Request {
  return &Request{
    Header: make(http.Header),
    path: path,
  }
}

func (r *Request) With(fn func(*Request) *Request) *Request {
  return fn(r)
}

func (r *Request) WithBody(body io.Reader) *Request {
  r.body = body
  return r
}

func (r *Request) WithHeader(key, value string) *Request {
  r.Header.Add(key, value)
  return r
}
