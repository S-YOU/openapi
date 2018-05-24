package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T19:11:33+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *ServerList) Clear() error {
	*v = ServerList(nil)
	return nil
}

func (v ServerList) Resolve(resolver *Resolver) error {
	if len(v) > 0 {
		for i, elem := range v {
			if err := elem.Resolve(resolver); err != nil {
				return errors.Wrapf(err, `failed to resolve ServerList (index = %d)`, i)
			}
		}
	}
	return nil
}

func (v *ServerList) UnmarshalJSON(data []byte) error {
	var proxy []*server
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}

	if len(proxy) == 0 {
		*v = ServerList(nil)
		return nil
	}

	tmp := make(ServerList, len(proxy))
	for i, value := range proxy {
		tmp[i] = value
	}
	*v = tmp
	return nil
}
