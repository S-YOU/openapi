package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T21:56:43+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type exampleMarshalProxy struct {
	Description   string      `json:"description" yaml:"description"`
	Value         interface{} `json:"value" yaml:"value"`
	ExternalValue string      `json:"external_value" yaml:"external_value"`
}

type exampleUnmarshalProxy struct {
	Description   string      `json:"description" yaml:"description"`
	Value         interface{} `json:"value" yaml:"value"`
	ExternalValue string      `json:"external_value" yaml:"external_value"`
}

func (v *example) MarshalJSON() ([]byte, error) {
	var proxy exampleMarshalProxy
	proxy.Description = v.description
	proxy.Value = v.value
	proxy.ExternalValue = v.externalValue
	return json.Marshal(proxy)
}

func (v *example) UnmarshalJSON(data []byte) error {
	var proxy exampleUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	v.description = proxy.Description
	v.value = proxy.Value
	v.externalValue = proxy.ExternalValue
	return nil
}
