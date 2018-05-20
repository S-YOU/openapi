package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T21:56:43+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type responseMarshalProxy struct {
	Description string               `json:"description" builder:"required"`
	Headers     map[string]Header    `json:"headers,omitempty" builder:"-"`
	Content     map[string]MediaType `json:"content,omitempty" builder:"-"`
	Links       map[string]Link      `json:"links,omitempty" builder:"-"`
}

type responseUnmarshalProxy struct {
	Description string                     `json:"description" builder:"required"`
	Headers     map[string]json.RawMessage `json:"headers,omitempty" builder:"-"`
	Content     map[string]json.RawMessage `json:"content,omitempty" builder:"-"`
	Links       map[string]json.RawMessage `json:"links,omitempty" builder:"-"`
}

func (v *response) MarshalJSON() ([]byte, error) {
	var proxy responseMarshalProxy
	proxy.Description = v.description
	proxy.Headers = v.headers
	proxy.Content = v.content
	proxy.Links = v.links
	return json.Marshal(proxy)
}

func (v *response) UnmarshalJSON(data []byte) error {
	var proxy responseUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	v.description = proxy.Description

	if len(proxy.Headers) > 0 {
		m := make(map[string]Header)
		for key, pv := range proxy.Headers {
			var decoded header
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Headers`, key)
			}
			m[key] = &decoded
		}
		v.headers = m
	}

	if len(proxy.Content) > 0 {
		m := make(map[string]MediaType)
		for key, pv := range proxy.Content {
			var decoded mediaType
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Content`, key)
			}
			m[key] = &decoded
		}
		v.content = m
	}

	if len(proxy.Links) > 0 {
		m := make(map[string]Link)
		for key, pv := range proxy.Links {
			var decoded link
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Links`, key)
			}
			m[key] = &decoded
		}
		v.links = m
	}
	return nil
}
