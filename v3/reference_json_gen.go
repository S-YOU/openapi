package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:06:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type referenceMarshalProxy struct {
	URL string `json:"-"`
}

type referenceUnmarshalProxy struct {
	URL string `json:"-"`
}

func (v *reference) MarshalJSON() ([]byte, error) {
	var proxy referenceMarshalProxy
	proxy.URL = v.uRL
	return json.Marshal(proxy)
}

func (v *reference) UnmarshalJSON(data []byte) error {
	var proxy referenceUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	v.uRL = proxy.URL
	return nil
}
