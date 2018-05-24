package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T19:11:33+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *HeaderMap) Clear() error {
	*v = make(HeaderMap)
	return nil
}

func (v HeaderMap) Resolve(resolver *Resolver) error {
	if len(v) > 0 {
		for name, elem := range v {
			if err := elem.Resolve(resolver); err != nil {
				return errors.Wrapf(err, `failed to resolve HeaderMap (key = %s)`, name)
			}
		}
	}
	return nil
}

func (v HeaderMap) QueryJSON(path string) (ret interface{}, ok bool) {
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

func (v *HeaderMap) UnmarshalJSON(data []byte) error {
	var proxy map[string]*header
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}
	tmp := make(map[string]Header)
	for name, value := range proxy {
		value.setName(name)
		tmp[name] = value
	}
	*v = tmp
	return nil
}
