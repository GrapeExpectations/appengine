package service

import (
	"context"
	"fmt"

	"github.com/GrapeExpectations/appengine/errors"
	"google.golang.org/appengine"
)

// GetModuleLocation returns the path to the specified module, if found
func GetModuleLocation(ctx context.Context, serviceName string) (string, error) {
	module, err := appengine.ModuleHostname(ctx, serviceName, "", "")
	if err != nil {
		return "", errors.Wrap(err,
			errors.Message{
				Pkg: "service",
				Fn:  "Module",
				Msg: fmt.Sprintf("error finding service [%v]", serviceName),
			})
	}

	return module, nil
}
