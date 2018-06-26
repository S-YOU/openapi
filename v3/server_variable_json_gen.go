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

type serverVariableMarshalProxy struct {
	Reference   string     `json:"$ref,omitempty"`
	Enum        StringList `json:"enum"`
	Default     string     `json:"default"`
	Description string     `json:"description"`
}

type serverVariableUnmarshalProxy struct {
	Reference   string     `json:"$ref,omitempty"`
	Enum        StringList `json:"enum"`
	Default     string     `json:"default"`
	Description string     `json:"description"`
}

func (v *serverVariable) MarshalJSON() ([]byte, error) {
	var proxy serverVariableMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Enum = v.enum
	proxy.Default = v.defaultValue
	proxy.Description = v.description
	return json.Marshal(proxy)
}

func (v *serverVariable) UnmarshalJSON(data []byte) error {
	var proxy serverVariableUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.enum = proxy.Enum
	v.defaultValue = proxy.Default
	v.description = proxy.Description
	return nil
}

func (v *serverVariable) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "enum":
		target = v.enum
	case "default":
		target = v.defaultValue
	case "description":
		target = v.description
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
