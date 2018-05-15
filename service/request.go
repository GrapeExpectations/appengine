package service

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	body     io.Reader
	Header   http.Header
	method   string
	module   string
	path     string
	protocol string
}

func NewRequest(path string) *Request {
	return &Request{
		Header: make(http.Header),
		path:   path,
	}
}

func (r *Request) HTTPRequest() (*http.Request, error) {
	u, err := url.Parse(fmt.Sprintf("%s%s%s", r.protocol, r.module, r.path))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(r.method, u.String(), r.body)
	if err != nil {
		return nil, err
	}

	req.Header = r.Header
	return req, nil
}

func (r *Request) Method(m string) *Request {
	r.method = m
	return r
}

func (r *Request) Module(m string) *Request {
	r.module = m
	return r
}

func (r *Request) Protocol(p string) *Request {
	r.protocol = p
	return r
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
