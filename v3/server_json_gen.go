package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:06:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type serverMarshalProxy struct {
	URL         string                    `json:"url" builder:"required"`
	Description string                    `json:"description,omitempty"`
	Variables   map[string]ServerVariable `json:"variables,omitempty"`
}

type serverUnmarshalProxy struct {
	URL         string                    `json:"url" builder:"required"`
	Description string                    `json:"description,omitempty"`
	Variables   map[string]ServerVariable `json:"variables,omitempty"`
}

func (v *server) MarshalJSON() ([]byte, error) {
	var proxy serverMarshalProxy
	proxy.URL = v.uRL
	proxy.Description = v.description
	proxy.Variables = v.variables
	return json.Marshal(proxy)
}

func (v *server) UnmarshalJSON(data []byte) error {
	var proxy serverUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	v.uRL = proxy.URL
	v.description = proxy.Description
	v.variables = proxy.Variables
	return nil
}
