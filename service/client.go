package service

import (
  "context"
  "fmt"
  "google.golang.org/appengine"
  "google.golang.org/appengine/log"
  "google.golang.org/appengine/urlfetch"
  "net/http"
)

type ServiceClient struct {
  client *http.Client
  module string
}

func NewServiceClient(ctx context.Context, service string) (*ServiceClient, error) {
  client := urlfetch.Client(ctx)

  m, err := appengine.ModuleHostname(ctx, service, "", "")
	if err != nil {
    log.Debugf(ctx, "error finding service [%v]: %v", service, err)
    return nil, err
	}

  return &ServiceClient {client, m}, nil
}

func (c *ServiceClient) Delete(r *Request) (*http.Response, error) {
  return c.Request("DELETE", r)
}

func (c *ServiceClient) Get(r *Request) (*http.Response, error) {
  return c.Request("GET", r)
}

func (c *ServiceClient) Post(r *Request) (*http.Response, error) {
  return c.Request("POST", r)
}

func (c *ServiceClient) Put(r *Request) (*http.Response, error) {
  return c.Request("POST", r)
}

func (c *ServiceClient) Request(method string, r *Request) (*http.Response, error) {
  url := fmt.Sprintf("%s%s%s", protocol(), c.module, r.path)
  req, err := http.NewRequest(method, url, r.body)
	if err != nil {
    return nil, err
	}

  req.Header = r.Header

  return c.client.Do(req)
}

func protocol() string {
  if appengine.IsDevAppServer() {
    return "http://"
  }
  return "https://"
}
