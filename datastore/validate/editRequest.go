package validate

import (
	"fmt"
	"net/http"

	ds "github.com/GrapeExpectations/appengine/datastore"
	"github.com/GrapeExpectations/appengine/errors"
)

// EditRequest validates that specified method is either "POST" or "PUT".  If it
// is another method, a StatusError is returned, with StatusMethodNotAllowed.
// If the method is "POST" the entity is validated to not have an existing Key,
// and returns a StatusError with StatusBadRequest if the entity has an existing Key.
// If the method is "PUT" the key, keystring, and kind are validated to match.
func EditRequest(method, keyString, kind string, e ds.Entity) error {
	if method == "POST" {
		if e.GetKey() != nil {
			return errors.New(http.StatusBadRequest, "new entity already has key")
		}
	} else if method == "PUT" {
		if e.GetKey() == nil {
			key, err := KeyString(keyString, kind)
			if err != nil {
				return errors.Wrap(err, "error validating put key")
			}
			e.SetKey(key)
		} else if e.GetKey().Encode() != keyString {
			return errors.New(http.StatusBadRequest, "entity key doesn't match request")
		} else if e.GetKey().Kind() != kind {
			return errors.New(http.StatusBadRequest, "incorrect entity key kind")
		}
	} else {
		return errors.New(http.StatusMethodNotAllowed, fmt.Sprintf("invalid request Method <%s>", method))
	}

	return nil
}
