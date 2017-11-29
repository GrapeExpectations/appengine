package helper

import (
  "google.golang.org/appengine"
  "google.golang.org/appengine/aetest"
  "testing"
)

func TestNamespaceFromContext(t *testing.T) {
	const namespace = "local"
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	ctx, err = appengine.Namespace(ctx, namespace)
	if err != nil {
		t.Fatal(err)
	}

	ns := NamespaceFromContext(ctx)
	if ns != namespace {
		t.Errorf("bad namespace, got: %s, want: %s", ns, namespace)
	}
}
