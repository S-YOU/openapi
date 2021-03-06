package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// ResponsesVisitor is an interface for objects that knows
// how to process Responses elements while traversing the OpenAPI structure
type ResponsesVisitor interface {
	VisitResponses(context.Context, Responses) error
}

func visitResponses(ctx context.Context, elem Responses) error {
	if checker, ok := elem.(interface{ IsValid() bool }); ok {
		if !checker.IsValid() {
			return nil
		}
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(responsesVisitorCtxKey{}).(ResponsesVisitor); ok {
		if err := v.VisitResponses(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Responses element`)
		}
	}

	if child := elem.Default(); child != nil {
		if err := visitResponse(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Default element for Responses`)
		}
	}

	for iter := elem.Responses(); iter.Next(); {
		key, value := iter.Item()
		if err := visitResponse(context.WithValue(ctx, responseMapKeyVisitorCtxKey{}, key), value); err != nil {
			return errors.Wrap(err, `failed to visit Responses element for Responses`)
		}
	}
	return nil
}
