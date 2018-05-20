package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T21:56:43+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type discriminatorMarshalProxy struct {
	PropertyName string            `json:"property_name" yaml:"property_name"`
	Mapping      map[string]string `json:"mapping" yaml:"mapping"`
}

type discriminatorUnmarshalProxy struct {
	PropertyName string `json:"property_name" yaml:"property_name"`
}

func (v *discriminator) MarshalJSON() ([]byte, error) {
	var proxy discriminatorMarshalProxy
	proxy.PropertyName = v.propertyName
	proxy.Mapping = v.mapping
	return json.Marshal(proxy)
}

func (v *discriminator) UnmarshalJSON(data []byte) error {
	var proxy discriminatorUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	v.propertyName = proxy.PropertyName
	return nil
}
