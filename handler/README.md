# Handler Library

## Handler

The Handler helper is intended to simplify adding common request wrappers around your handler function, in a composable manner.  The `Route()` function returns the handler path, and wrapped handler function.

```go
package api

import (
  "context"
  "appengine/handler"
  "github.com/gorilla/mux"
  "lib/ns"
)

func init() {
  list := handler.NewHandler("/", listHandler).
    NamespacedRequest(ns.GetNamespaceFromHost).
    ContentType("application/json").
    CORS()

  write := handler.NewHandler("/write", writeHandler).
    NamespacedRequest(ns.GetNamespaceFromHeader).
    ContentType("application/json").
    ServiceRequest()

  r := mux.NewRouter()
  r.HandleFunc(list.Route()).Methods("GET")
  r.HandleFunc(write.Route()).Methods("POST")
  http.Handle("/", r)
}

func listHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
  // ... list the things
}

func writeHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
  // ... write a thing
}
```

### Content Type

The `ContentType()` wrapper takes in a string, specifying the content type that the handler returns.  It sets the `Content-type` header in the response.

```go
list := request.NewHandler("/", listHandler).
  ContentType("application/json")
```
In the above example, the handler's `content-type` response header is set to `application/json`.

### CORS

The `CORS()` wrapper sets the `Access-Control-Allow-Headers`,
`Access-Control-Allow-Credentials`, `Access-Control-Allow-Origin`, and `Access-Control-Allow-Method` headers in the response.  These are currently not configurable, but will be in the future.

Additionally, when an `OPTIONS` request is received, the wrapper will respond without continuing down the chain.

```go
list := request.NewHandler("/", listHandler).
  CORS()
```
When an `OPTIONS` request is received in the above example, the response is returned (with cors headers) before the `listhandler` handler is called.

### Namespaced Request

The `NamespacedRequest()` wrapper takes in a function,
```go
  func(*http.Request) (string, error)
```
which accepts an `*http.Request` parameter, and returns a string and error.  That function will be called by the wrapping function to determine the namespace to use (returned by the string).  If an error is returned, the wrapping function logs the error, and returns `http.StatusBadRequest`.  The wrapping function applies this namespace to the context, and then calls the inner function.  If the wrapping function encounters an error while applying the namespace (for example, if the namespace string returned is not a legal namespace name), the wrapping function returns `http.StatusInternalServerError`.

```go
list := request.NewHandler("/", listHandler).
  NamespacedRequest(ns.GetNamespaceFromHost)
```
In the above example, `ns.GetNamespaceFromHost` is a function which takes `*http.Request` as a parameter, and returns `(string, error)`.

### Service Request

The `ServiceRequest()` wrapper takes no parameters.  It checks that either `appengine.IsDevAppServer()` is true, or that the requesting app (as determined by the `X-Appengine-Inbound-Appid` header... [see documentation](https://cloud.google.com/appengine/docs/standard/go/appidentity/)) matches the service's AppID.  If these conditions are not met, the wrapping function returns `http.StatusForbidden`.

```go
write := request.NewHandler("/write", writeHandler).
  ServiceRequest()
```
In the above example, the `/write` route can only be called by a service which has the same AppId

### Wrap

The `Wrap()` wrapper takes in a function,
```go
func(context.Context, http.ResponseWriter, *http.Request,
	func(context.Context, http.ResponseWriter, *http.Request))
```
which accepts a Context, ResponseWriter, and Request, as well as a function(taking those same parameters).  This allows you to provide a generic wrapper to modify the context, request, and/or response as you need.  It is your function's responsibility to call the next function in the chain (the provided function), as appropriate.

```go
wrapper := func(ctx context.Context, w http.ResponseWriter, r *http.Request, fn func(context.Context, http.ResponseWriter, *http.Request)) {
  w.Header().Set("Content-type", "application/json")
  fn(ctx, w, r)
}

list := request.NewHandler("/", listHandler).
  Wrap(wrapper)

```
In the above example, the handler's `content-type` response header is set to `application/json`, by the `wrapper` function.

*Note: This is only an example wrapper.  The same functionality can be achieved using the ContentType() wrapper instead
