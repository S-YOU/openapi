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

type requestBodyMarshalProxy struct {
	Reference   string       `json:"$ref,omitempty"`
	Description string       `json:"description,omitempty"`
	Content     MediaTypeMap `json:"content,omitempty"`
	Required    bool         `json:"required,omitempty"`
}

type requestBodyUnmarshalProxy struct {
	Reference   string       `json:"$ref,omitempty"`
	Description string       `json:"description,omitempty"`
	Content     MediaTypeMap `json:"content,omitempty"`
	Required    bool         `json:"required,omitempty"`
}

func (v *requestBody) MarshalJSON() ([]byte, error) {
	var proxy requestBodyMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Description = v.description
	proxy.Content = v.content
	proxy.Required = v.required
	return json.Marshal(proxy)
}

func (v *requestBody) UnmarshalJSON(data []byte) error {
	var proxy requestBodyUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.description = proxy.Description
	v.content = proxy.Content
	v.required = proxy.Required

	v.postUnmarshalJSON()
	return nil
}

func (v *requestBody) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*requestBody)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type RequestBody, but got %T`, resolved)
		}
		mutator := MutateRequestBody(v)
		mutator.Name(asserted.Name())
		mutator.Description(asserted.Description())
		for iter := asserted.Content(); iter.Next(); {
			key, item := iter.Item()
			mutator.Content(key, item)
		}
		mutator.Required(asserted.Required())
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.content != nil {
		if err := v.content.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Content`)
		}
	}
	return nil
}

func (v *requestBody) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "description":
		target = v.description
	case "content":
		target = v.content
	case "required":
		target = v.required
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
