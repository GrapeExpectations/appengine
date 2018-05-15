package service

import (
	"context"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

type ServiceClient struct {
	client *http.Client
	module string
}

func NewServiceClient(ctx context.Context, module string) *ServiceClient {
	client := urlfetch.Client(ctx)

	return &ServiceClient{client, module}
}

func (c *ServiceClient) Delete(r *Request) (*http.Response, error) {
	r.Method("DELETE")
	return c.Do(r)
}

func (c *ServiceClient) Do(r *Request) (*http.Response, error) {
	r.Protocol(protocol())
	r.Module(c.module)

	req, err := r.HTTPRequest()
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c *ServiceClient) Get(r *Request) (*http.Response, error) {
	r.Method("GET")
	return c.Do(r)
}

func (c *ServiceClient) Post(r *Request) (*http.Response, error) {
	r.Method("POST")
	return c.Do(r)
}

func (c *ServiceClient) Put(r *Request) (*http.Response, error) {
	r.Method("PUT")
	return c.Do(r)
}

func protocol() string {
	if appengine.IsDevAppServer() {
		return "http://"
	}
	return "https://"
}
