package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T20:40:48+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *SchemaMap) Clear() error {
	*v = make(SchemaMap)
	return nil
}

func (v SchemaMap) QueryJSON(path string) (ret interface{}, ok bool) {
	if path == `` {
		return v, true
	}

	var frag string
	frag, path = extractFragFromPath(path)
	target, ok := v[frag]
	if !ok {
		return nil, false
	}

	if qj, ok := target.(QueryJSONer); ok {
		return qj.QueryJSON(path)
	}

	if path == `` {
		return target, true
	}
	return nil, false
}

func (v *SchemaMap) UnmarshalJSON(data []byte) error {
	var proxy map[string]*schema
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}
	tmp := make(map[string]Schema)
	for name, value := range proxy {
		value.setName(name)
		tmp[name] = value
	}
	*v = tmp
	return nil
}
