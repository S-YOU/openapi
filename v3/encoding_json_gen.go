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

type encodingMarshalProxy struct {
	Reference     string    `json:"$ref,omitempty"`
	ContentType   string    `json:"contentType"`
	Headers       HeaderMap `json:"headers"`
	Explode       bool      `json:"explode"`
	AllowReserved bool      `json:"allowReserved"`
}

type encodingUnmarshalProxy struct {
	Reference     string    `json:"$ref,omitempty"`
	ContentType   string    `json:"contentType"`
	Headers       HeaderMap `json:"headers"`
	Explode       bool      `json:"explode"`
	AllowReserved bool      `json:"allowReserved"`
}

func (v *encoding) MarshalJSON() ([]byte, error) {
	var proxy encodingMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.ContentType = v.contentType
	proxy.Headers = v.headers
	proxy.Explode = v.explode
	proxy.AllowReserved = v.allowReserved
	return json.Marshal(proxy)
}

func (v *encoding) UnmarshalJSON(data []byte) error {
	var proxy encodingUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.contentType = proxy.ContentType
	v.headers = proxy.Headers
	v.explode = proxy.Explode
	v.allowReserved = proxy.AllowReserved
	return nil
}

func (v *encoding) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*encoding)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type Encoding, but got %T`, resolved)
		}
		mutator := MutateEncoding(v)
		mutator.Name(asserted.Name())
		mutator.ContentType(asserted.ContentType())
		for iter := asserted.Headers(); iter.Next(); {
			key, item := iter.Item()
			mutator.Header(key, item)
		}
		mutator.Explode(asserted.Explode())
		mutator.AllowReserved(asserted.AllowReserved())
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.headers != nil {
		if err := v.headers.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Headers`)
		}
	}
	return nil
}

func (v *encoding) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "contentType":
		target = v.contentType
	case "headers":
		target = v.headers
	case "explode":
		target = v.explode
	case "allowReserved":
		target = v.allowReserved
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
