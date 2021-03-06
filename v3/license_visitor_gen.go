package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause

// LicenseVisitor is an interface for objects that knows
// how to process License elements while traversing the OpenAPI structure
type LicenseVisitor interface {
	VisitLicense(context.Context, License) error
}

func visitLicense(ctx context.Context, elem License) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(licenseVisitorCtxKey{}).(LicenseVisitor); ok {
		if err := v.VisitLicense(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit License element`)
		}
	}
	return nil
}
