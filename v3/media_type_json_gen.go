package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:06:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type mediaTypeMarshalProxy struct {
	Schema   Schema              `json:"schema,omitempty"`
	Example  interface{}         `json:"example,omitempty"`
	Examples map[string]Example  `json:"examples,omitempty"`
	Encoding map[string]Encoding `json:"encoding,omitempty"`
}

type mediaTypeUnmarshalProxy struct {
	Schema   json.RawMessage     `json:"schema,omitempty"`
	Example  interface{}         `json:"example,omitempty"`
	Examples map[string]Example  `json:"examples,omitempty"`
	Encoding map[string]Encoding `json:"encoding,omitempty"`
}

func (v *mediaType) MarshalJSON() ([]byte, error) {
	var proxy mediaTypeMarshalProxy
	proxy.Schema = v.schema
	proxy.Example = v.example
	proxy.Examples = v.examples
	proxy.Encoding = v.encoding
	return json.Marshal(proxy)
}

func (v *mediaType) UnmarshalJSON(data []byte) error {
	var proxy mediaTypeUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}

	if len(proxy.Schema) > 0 {
		var decoded schema
		if err := json.Unmarshal(proxy.Schema, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Schema`)
		}

		v.schema = &decoded
	}
	v.example = proxy.Example
	v.examples = proxy.Examples
	v.encoding = proxy.Encoding
	return nil
}