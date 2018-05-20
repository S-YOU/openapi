package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:06:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type licenseMarshalProxy struct {
	Name string `json:"name" builder:"required"`
	URL  string `json:"url,omitempty"`
}

type licenseUnmarshalProxy struct {
	Name string `json:"name" builder:"required"`
	URL  string `json:"url,omitempty"`
}

func (v *license) MarshalJSON() ([]byte, error) {
	var proxy licenseMarshalProxy
	proxy.Name = v.name
	proxy.URL = v.uRL
	return json.Marshal(proxy)
}

func (v *license) UnmarshalJSON(data []byte) error {
	var proxy licenseUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	v.name = proxy.Name
	v.uRL = proxy.URL
	return nil
}
