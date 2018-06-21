package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *SchemaList) Clear() error {
	*v = SchemaList(nil)
	return nil
}

func (v SchemaList) Resolve(resolver *Resolver) error {
	if len(v) > 0 {
		for i, elem := range v {
			if err := elem.Resolve(resolver); err != nil {
				return errors.Wrapf(err, `failed to resolve SchemaList (index = %d)`, i)
			}
		}
	}
	return nil
}

func (v *SchemaList) UnmarshalJSON(data []byte) error {
	var proxy []*schema
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}

	if len(proxy) == 0 {
		*v = SchemaList(nil)
		return nil
	}

	tmp := make(SchemaList, len(proxy))
	for i, value := range proxy {
		tmp[i] = value
	}
	*v = tmp
	return nil
}
