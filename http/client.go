package http

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
    log.Errorf(ctx, "error finding service [%v]: %v", service, err)
    return nil, err
	}

  return &ServiceClient {client, m}, nil
}

func (c *ServiceClient) Get(r *Request) (*http.Response, error) {
  url := fmt.Sprintf("%s%s%s", protocol(), c.module, r.path)
  req, err := http.NewRequest("GET", url, nil)
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
