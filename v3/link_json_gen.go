package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/pkg/errors"
)

var _ = log.Printf
var _ = json.Unmarshal
var _ = errors.Cause

type linkMarshalProxy struct {
	Reference    string       `json:"$ref,omitempty"`
	OperationRef string       `json:"operationRef"`
	OperationID  string       `json:"operationId"`
	Parameters   InterfaceMap `json:"parameters"`
	RequestBody  interface{}  `json:"requestBody"`
	Description  string       `json:"description"`
	Server       Server       `json:"server"`
}

type linkUnmarshalProxy struct {
	Reference    string          `json:"$ref,omitempty"`
	OperationRef string          `json:"operationRef"`
	OperationID  string          `json:"operationId"`
	Parameters   InterfaceMap    `json:"parameters"`
	RequestBody  interface{}     `json:"requestBody"`
	Description  string          `json:"description"`
	Server       json.RawMessage `json:"server"`
}

func (v *link) MarshalJSON() ([]byte, error) {
	var proxy linkMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.OperationRef = v.operationRef
	proxy.OperationID = v.operationID
	proxy.Parameters = v.parameters
	proxy.RequestBody = v.requestBody
	proxy.Description = v.description
	proxy.Server = v.server
	return json.Marshal(proxy)
}

func (v *link) UnmarshalJSON(data []byte) error {
	var proxy linkUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.operationRef = proxy.OperationRef
	v.operationID = proxy.OperationID
	v.parameters = proxy.Parameters
	v.requestBody = proxy.RequestBody
	v.description = proxy.Description

	if len(proxy.Server) > 0 {
		var decoded server
		if err := json.Unmarshal(proxy.Server, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Server`)
		}

		v.server = &decoded
	}
	return nil
}

func (v *link) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*link)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type Link, but got %T`, resolved)
		}
		mutator := MutateLink(v)
		mutator.Name(asserted.Name())
		mutator.OperationRef(asserted.OperationRef())
		mutator.OperationID(asserted.OperationID())
		for iter := asserted.Parameters(); iter.Next(); {
			key, item := iter.Item()
			mutator.Parameter(key, item)
		}
		mutator.RequestBody(asserted.RequestBody())
		mutator.Description(asserted.Description())
		mutator.Server(asserted.Server())
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.parameters != nil {
		if err := v.parameters.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Parameters`)
		}
	}
	if v.server != nil {
		if err := v.server.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Server`)
		}
	}
	return nil
}

func (v *link) QueryJSON(path string) (ret interface{}, ok bool) {
	path = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })
	if path == "" {
		return v, true
	}

	var frag string
	if i := strings.Index(path, "/"); i > -1 {
		frag = path[:i]
		path = path[i+1:]
	} else {
		frag = path
		path = ""
	}

	var target interface{}

	switch frag {
	case "operationRef":
		target = v.operationRef
	case "operationId":
		target = v.operationID
	case "parameters":
		target = v.parameters
	case "requestBody":
		target = v.requestBody
	case "description":
		target = v.description
	case "server":
		target = v.server
	default:
		return nil, false
	}

	if qj, ok := target.(QueryJSONer); ok {
		return qj.QueryJSON(path)
	}
	if path == "" {
		return target, true
	}
	return nil, false
}
