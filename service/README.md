# Service Library

## Service Client

The `ServiceClient` helper is intended to simplify sending requests to services in your proejct.

```go
package api

import (
  "appengine/service"
)

func memberList(ctx context.Context) error {
  // ... call the member service to list the members
  client, err := service.NewServiceClient(ctx, "member")
  if err != nil {
    return err
  }

  // ...
}
```

### New Service Client

The `NewServiceClient()` function takes the context, and the service name as parameters.  It returns a configured `ServiceClient`, and an error.  An error is returned when the specified service cannot be found (using [`appengine.ModuleHostname()`](https://cloud.google.com/appengine/docs/standard/go/reference#ModuleHostname)).

```go
client, err := service.NewServiceClient(ctx, "member")
if err != nil {
  // ... handle error
}
```

### Delete, Get, Post, Put... Request

`Delete()`, `Get()`, `Post()`, and `Put()` are all convenience methods that call [`Request()`](#request).  Each function takes in only a [`Request`](#request-1) pointer, calling `Request()` with the appropriate method.

```go
req := service.NewRequest("/list")
resp, err := client.Get(req)
if err != nil {
  // ... handle error
}
```

### Request

The `Request()` function takes in a method name (`DELETE, GET, POST, PUT`), and a [`Request`](#request-1) pointer.  It makes a "`method`" request to the configured service endpoint, using the `Request` pointer's `path`, applying the pointer's `Header`s, and sending the pointer's `body`.

```go
req := service.NewRequest("/user")
resp, err := client.Request("POST", req)
if err != nil {
  // ... handle error
}
```

## Request

The `Request` helper is intended to simplify creating requests to send to your services, in a composable manner.

```go
package api

import (
  "appengine/service"
  "lib/ns"
)

func memberWrite(ctx context.Context, encodedMember string) error {
  reader := strings.NewReader(encodedMember)
  trackingGUID := ctx.Value("Ctx-Tracking-Id")

  // ... configure request
  req := service.NewRequest("/write").
    With(ns.AddNamespace(ctx)).
    WithBody(reader).
    WithHeader("X-Tracking-Id", trackingGUID)

  // ... configure a client for the "member" service
  client, err := service.NewServiceClient(ctx, "member")
  if err != nil {
    return err
  }

  // ... call the member service to write the member
  resp, err := client.Post(req)
  if err != nil {
    return err
  }
}
```

### New Request

`NewRequest()` takes in the `path` for the request, and returns a configured `Request`.

```go
  req := service.NewRequest("/write")
```

### With Header

`WithHeader()` takes in a header key, and value (strings).  The `value` is added to the `Request` header, using the provided `key`.

```go
  req := service.NewRequest("/write").
    WithHeader("X-Tracking-Id", trackingGUID)
```

### With Body

`WithBody()` uses the provided `io.Reader` as the request body to send.

```go
  reader := strings.NewReader(body)
  req := service.NewRequest("/write").
    WithBody(reader)
```

### With

`With()` provides a generic way to modify the `Request`, by taking a function,

```go
  func(*Request) *Request
```

which accepts and returns a `Request` pointer.

```go
func SendUser(user *SystemUser) func(*Request) *Request {
  encoded := user.Encode()
  return func(r *Request) *Request {
    return r.WithHeader("X-System-User", encoded)
  }
}

func memberWrite(ctx context.Context, user *SystemUser) error {
  req := service.NewRequest("/list").
    With(SendUser(user))

  // ...
}
```
In the above example, the `SendUser` function is used to encode a system user, and add that as a header value.  An alternative to the `SendUser` function would be,

```go
func memberWrite(ctx context.Context, user *SystemUser) error {
  req := service.NewRequest("/list").
    WithHeader("X-System-User", user.Encode())

  // ...
}
```
The `With(SendUser(user))` approach can be much more valuable if additional processing is required (done in the `SendUser` function), rather than continually repeating that code.
