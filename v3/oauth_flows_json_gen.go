package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:06:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type oAuthFlowsMarshalProxy struct {
	Implicit          OAuthFlow `json:"implicit" yaml:"implicit"`
	Password          OAuthFlow `json:"password" yaml:"password"`
	ClientCredentials OAuthFlow `json:"clientCredentials" yaml:"clientCredentials"`
	AuthorizationCode OAuthFlow `json:"authorizationCode" yaml:"authorizationCode"`
}

type oAuthFlowsUnmarshalProxy struct {
	Implicit          json.RawMessage `json:"implicit" yaml:"implicit"`
	Password          json.RawMessage `json:"password" yaml:"password"`
	ClientCredentials json.RawMessage `json:"clientCredentials" yaml:"clientCredentials"`
	AuthorizationCode json.RawMessage `json:"authorizationCode" yaml:"authorizationCode"`
}

func (v *oAuthFlows) MarshalJSON() ([]byte, error) {
	var proxy oAuthFlowsMarshalProxy
	proxy.Implicit = v.implicit
	proxy.Password = v.password
	proxy.ClientCredentials = v.clientCredentials
	proxy.AuthorizationCode = v.authorizationCode
	return json.Marshal(proxy)
}

func (v *oAuthFlows) UnmarshalJSON(data []byte) error {
	var proxy oAuthFlowsUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}

	if len(proxy.Implicit) > 0 {
		var decoded oAuthFlow
		if err := json.Unmarshal(proxy.Implicit, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Implicit`)
		}

		v.implicit = &decoded
	}

	if len(proxy.Password) > 0 {
		var decoded oAuthFlow
		if err := json.Unmarshal(proxy.Password, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Password`)
		}

		v.password = &decoded
	}

	if len(proxy.ClientCredentials) > 0 {
		var decoded oAuthFlow
		if err := json.Unmarshal(proxy.ClientCredentials, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ClientCredentials`)
		}

		v.clientCredentials = &decoded
	}

	if len(proxy.AuthorizationCode) > 0 {
		var decoded oAuthFlow
		if err := json.Unmarshal(proxy.AuthorizationCode, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field AuthorizationCode`)
		}

		v.authorizationCode = &decoded
	}
	return nil
}
