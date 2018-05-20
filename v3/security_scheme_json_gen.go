package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:06:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type securitySchemeMarshalProxy struct {
	Type             string     `json:"type" yaml:"type"`
	Description      string     `json:"description" yaml:"description"`
	Name             string     `json:"name" yaml:"name"`
	In               string     `json:"in" yaml:"in"`
	Scheme           string     `json:"scheme" yaml:"scheme"`
	BearerFormat     string     `json:"bearerFormat" yaml:"bearerFormat"`
	Flows            OAuthFlows `json:"flows" yaml:"flows"`
	OpenIDConnectURL string     `json:"openIdConnectUrl" yaml:"openIdConnectUrl"`
}

type securitySchemeUnmarshalProxy struct {
	Type             string          `json:"type" yaml:"type"`
	Description      string          `json:"description" yaml:"description"`
	Name             string          `json:"name" yaml:"name"`
	In               string          `json:"in" yaml:"in"`
	Scheme           string          `json:"scheme" yaml:"scheme"`
	BearerFormat     string          `json:"bearerFormat" yaml:"bearerFormat"`
	Flows            json.RawMessage `json:"flows" yaml:"flows"`
	OpenIDConnectURL string          `json:"openIdConnectUrl" yaml:"openIdConnectUrl"`
}

func (v *securityScheme) MarshalJSON() ([]byte, error) {
	var proxy securitySchemeMarshalProxy
	proxy.Type = v.typ
	proxy.Description = v.description
	proxy.Name = v.name
	proxy.In = v.in
	proxy.Scheme = v.scheme
	proxy.BearerFormat = v.bearerFormat
	proxy.Flows = v.flows
	proxy.OpenIDConnectURL = v.openIDConnectURL
	return json.Marshal(proxy)
}

func (v *securityScheme) UnmarshalJSON(data []byte) error {
	var proxy securitySchemeUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	v.typ = proxy.Type
	v.description = proxy.Description
	v.name = proxy.Name
	v.in = proxy.In
	v.scheme = proxy.Scheme
	v.bearerFormat = proxy.BearerFormat

	if len(proxy.Flows) > 0 {
		var decoded oAuthFlows
		if err := json.Unmarshal(proxy.Flows, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Flows`)
		}

		v.flows = &decoded
	}
	v.openIDConnectURL = proxy.OpenIDConnectURL
	return nil
}
