# Request Library

## Handler

The Handler helper is intended to simplify adding common request wrappers around your handler function, in a composable manner.  The `Route()` function returns the handler path, and wrapped handler function.

```go
package api

import (
  "context"
  "appengine/request"
  "github.com/gorilla/mux"
  "lib/ns"
)

func init() {
	list := request.NewHandler("/", listHandler).
      NamespacedRequest(ns.GetNamespaceFromHost)

	write := request.NewHandler("/write", writeHandler).
      NamespacedRequest(ns.GetNamespaceFromHeader).
      ServiceRequest()

	r := mux.NewRouter()
	r.HandleFunc(list.Route()).Methods("GET")
	r.HandleFunc(write.Route()).Methods("POST")
	http.Handle("/", r)
}

func listHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
  // ... list the things
}

func writeHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
  // ... write a thing
}
```

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
