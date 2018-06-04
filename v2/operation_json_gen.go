package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
)

var _ = json.Unmarshal
var _ = fmt.Fprintf
var _ = log.Printf
var _ = strconv.ParseInt
var _ = errors.Cause

type operationMarshalProxy struct {
	Reference    string                  `json:"$ref,omitempty"`
	Tags         StringList              `json:"tags,omitempty"`
	Summary      string                  `json:"summary,omitempty"`
	Description  string                  `json:"description,omitempty"`
	ExternalDocs ExternalDocumentation   `json:"externalDocs,omitempty"`
	OperationID  string                  `json:"operationId,omitempty"`
	Consumes     MIMETypeList            `json:"consumes,omitempty"`
	Produces     MIMETypeList            `json:"produces,omitempty"`
	Parameters   ParameterList           `json:"parameters,omitempty"`
	Responses    Responses               `json:"responses"`
	Schemes      SchemeList              `json:"schemes,omitempty"`
	Deprecated   bool                    `json:"deprecated,omitempty"`
	Security     SecurityRequirementList `json:"security,omitempty"`
}

func (v *operation) MarshalJSON() ([]byte, error) {
	var proxy operationMarshalProxy
	if s := v.reference; len(s) > 0 {
		return []byte(fmt.Sprintf(refOnlyTmpl, strconv.Quote(s))), nil
	}
	proxy.Tags = v.tags
	proxy.Summary = v.summary
	proxy.Description = v.description
	proxy.ExternalDocs = v.externalDocs
	proxy.OperationID = v.operationID
	proxy.Consumes = v.consumes
	proxy.Produces = v.produces
	proxy.Parameters = v.parameters
	proxy.Responses = v.responses
	proxy.Schemes = v.schemes
	proxy.Deprecated = v.deprecated
	proxy.Security = v.security
	buf, err := json.Marshal(proxy)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal struct`)
	}
	if len(v.extensions) > 0 {
		extBuf, err := json.Marshal(v.extensions)
		if err != nil || len(extBuf) <= 2 {
			return nil, errors.Wrap(err, `failed to marshal struct (extensions)`)
		}
		buf = append(append(buf[:len(buf)-1], ','), extBuf[1:]...)
	}
	return buf, nil
}

func (v *operation) UnmarshalJSON(data []byte) error {
	var proxy map[string]json.RawMessage
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if raw := proxy["$ref"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.reference); err != nil {
			return errors.Wrap(err, `failed to unmarshal $ref`)
		}
		return nil
	}

	mutator := MutateOperation(v)

	if raw := proxy["tags"]; len(raw) > 0 {
		var decoded StringList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Tags`)
		}
		for _, elem := range decoded {
			mutator.Tag(elem)
		}
		delete(proxy, "tags")
	}

	if raw := proxy["summary"]; len(raw) > 0 {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field summary`)
		}
		mutator.Summary(decoded)
		delete(proxy, "summary")
	}

	if raw := proxy["description"]; len(raw) > 0 {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field description`)
		}
		mutator.Description(decoded)
		delete(proxy, "description")
	}

	if raw := proxy["externalDocs"]; len(raw) > 0 {
		var decoded externalDocumentation
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ExternalDocs`)
		}

		mutator.ExternalDocs(&decoded)
		delete(proxy, "externalDocs")
	}

	if raw := proxy["operationId"]; len(raw) > 0 {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field operationId`)
		}
		mutator.OperationID(decoded)
		delete(proxy, "operationId")
	}

	if raw := proxy["consumes"]; len(raw) > 0 {
		var decoded MIMETypeList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Consumes`)
		}
		for _, elem := range decoded {
			mutator.Consume(elem)
		}
		delete(proxy, "consumes")
	}

	if raw := proxy["produces"]; len(raw) > 0 {
		var decoded MIMETypeList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Produces`)
		}
		for _, elem := range decoded {
			mutator.Produce(elem)
		}
		delete(proxy, "produces")
	}

	if raw := proxy["parameters"]; len(raw) > 0 {
		var decoded ParameterList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Parameters`)
		}
		for _, elem := range decoded {
			mutator.Parameter(elem)
		}
		delete(proxy, "parameters")
	}

	if raw := proxy["responses"]; len(raw) > 0 {
		var decoded responses
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Responses`)
		}

		mutator.Responses(&decoded)
		delete(proxy, "responses")
	}

	if raw := proxy["schemes"]; len(raw) > 0 {
		var decoded SchemeList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Schemes`)
		}
		for _, elem := range decoded {
			mutator.Scheme(elem)
		}
		delete(proxy, "schemes")
	}

	if raw := proxy["deprecated"]; len(raw) > 0 {
		var decoded bool
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field deprecated`)
		}
		mutator.Deprecated(decoded)
		delete(proxy, "deprecated")
	}

	if raw := proxy["security"]; len(raw) > 0 {
		var decoded SecurityRequirementList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Security`)
		}
		for _, elem := range decoded {
			mutator.Security(elem)
		}
		delete(proxy, "security")
	}

	for name, raw := range proxy {
		if strings.HasPrefix(name, `x-`) {
			var ext interface{}
			if err := json.Unmarshal(raw, &ext); err != nil {
				return errors.Wrapf(err, `failed to unmarshal field %s`, name)
			}
			mutator.Extension(name, ext)
		}
	}

	if err := mutator.Do(); err != nil {
		return errors.Wrap(err, `failed to  unmarshal JSON`)
	}
	return nil
}

func (v *operation) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "tags":
		target = v.tags
	case "summary":
		target = v.summary
	case "description":
		target = v.description
	case "externalDocs":
		target = v.externalDocs
	case "operationId":
		target = v.operationID
	case "consumes":
		target = v.consumes
	case "produces":
		target = v.produces
	case "parameters":
		target = v.parameters
	case "responses":
		target = v.responses
	case "schemes":
		target = v.schemes
	case "deprecated":
		target = v.deprecated
	case "security":
		target = v.security
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

func OperationFromJSON(buf []byte, v *Operation) error {
	var tmp operation
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}
	*v = &tmp
	return nil
}
