package validate

import (
	"net/http"

	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine/datastore"
)

// Key validates that the provided datastore key is not nil, and is of the specified Kind.
// It returns nil if succcessfully validated, or a StatusError if not.  When Kind is an
// empty string, any key kind is acceptable
func Key(key *datastore.Key, kind string) error {
	if key == nil {
		return errors.New(http.StatusBadRequest, "key not specified")
	}

	if kind != "" && key.Kind() != kind {
		return errors.New(http.StatusBadRequest, "wrong key kind")
	}

	return nil
}

// KeyString validates that the provided keystring is a valid datastore key, and is of the specified Kind.
// It returns the key if succcessfully validated, or a StatusError if not
func KeyString(keyString, kind string) (*datastore.Key, error) {
	key, err := datastore.DecodeKey(keyString)

	if err != nil {
		return nil, errors.Wrap(err, "invalid key").
			SetCode(http.StatusBadRequest)
	}

	if kind != "" && key.Kind() != kind {
		return nil, errors.New(http.StatusBadRequest, "key is wrong kind")
	}

	return key, nil
}
