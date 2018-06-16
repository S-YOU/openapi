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

type swaggerMarshalProxy struct {
	Reference           string                  `json:"$ref,omitempty"`
	Version             string                  `json:"swagger"`
	Info                Info                    `json:"info"`
	Host                string                  `json:"host,omitempty"`
	BasePath            string                  `json:"basePath,omitempty"`
	Schemes             SchemeList              `json:"schemes,omitempty"`
	Consumes            MIMETypeList            `json:"consumes,omitempty"`
	Produces            MIMETypeList            `json:"produces,omitempty"`
	Paths               Paths                   `json:"paths"`
	Definitions         InterfaceMap            `json:"definitions,omitempty"`
	Parameters          ParameterMap            `json:"parameters,omitempty"`
	Responses           ResponseMap             `json:"responses,omitempty"`
	SecurityDefinitions SecuritySchemeMap       `json:"securityDefinitions,omitempty"`
	Security            SecurityRequirementList `json:"security,omitempty"`
	Tags                TagList                 `json:"tags,omitempty"`
	ExternalDocs        ExternalDocumentation   `json:"externalDocs,omitempty"`
}

func (v *swagger) MarshalJSON() ([]byte, error) {
	var proxy swaggerMarshalProxy
	if s := v.reference; len(s) > 0 {
		return []byte(fmt.Sprintf(refOnlyTmpl, strconv.Quote(s))), nil
	}
	proxy.Version = v.version
	proxy.Info = v.info
	proxy.Host = v.host
	proxy.BasePath = v.basePath
	proxy.Schemes = v.schemes
	proxy.Consumes = v.consumes
	proxy.Produces = v.produces
	proxy.Paths = v.paths
	proxy.Definitions = v.definitions
	proxy.Parameters = v.parameters
	proxy.Responses = v.responses
	proxy.SecurityDefinitions = v.securityDefinitions
	proxy.Security = v.security
	proxy.Tags = v.tags
	proxy.ExternalDocs = v.externalDocs
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

// UnmarshalJSON defines how swagger is deserialized from JSON
func (v *swagger) UnmarshalJSON(data []byte) error {
	var proxy map[string]json.RawMessage
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if raw, ok := proxy["$ref"]; ok {
		if err := json.Unmarshal(raw, &v.reference); err != nil {
			return errors.Wrap(err, `failed to unmarshal $ref`)
		}
		return nil
	}

	mutator := MutateSwagger(v)

	const versionMapKey = "swagger"

	if raw, ok := proxy[versionMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field swagger`)
		}
		mutator.Version(decoded)
		delete(proxy, versionMapKey)
	}

	const infoMapKey = "info"

	if raw, ok := proxy[infoMapKey]; ok {
		var decoded info
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Info`)
		}

		mutator.Info(&decoded)
		delete(proxy, infoMapKey)
	}

	const hostMapKey = "host"

	if raw, ok := proxy[hostMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field host`)
		}
		mutator.Host(decoded)
		delete(proxy, hostMapKey)
	}

	const basePathMapKey = "basePath"

	if raw, ok := proxy[basePathMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field basePath`)
		}
		mutator.BasePath(decoded)
		delete(proxy, basePathMapKey)
	}

	const schemesMapKey = "schemes"
	if raw, ok := proxy[schemesMapKey]; ok {
		var decoded SchemeList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Schemes`)
		}
		for _, elem := range decoded {
			mutator.Scheme(elem)
		}
		delete(proxy, schemesMapKey)
	}

	const consumesMapKey = "consumes"
	if raw, ok := proxy[consumesMapKey]; ok {
		var decoded MIMETypeList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Consumes`)
		}
		for _, elem := range decoded {
			mutator.Consume(elem)
		}
		delete(proxy, consumesMapKey)
	}

	const producesMapKey = "produces"
	if raw, ok := proxy[producesMapKey]; ok {
		var decoded MIMETypeList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Produces`)
		}
		for _, elem := range decoded {
			mutator.Produce(elem)
		}
		delete(proxy, producesMapKey)
	}

	const pathsMapKey = "paths"

	if raw, ok := proxy[pathsMapKey]; ok {
		var decoded paths
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Paths`)
		}

		mutator.Paths(&decoded)
		delete(proxy, pathsMapKey)
	}

	const definitionsMapKey = "definitions"
	if raw, ok := proxy[definitionsMapKey]; ok {
		var decoded InterfaceMap
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Definitions`)
		}
		for key, elem := range decoded {
			mutator.Definition(key, elem)
		}
		delete(proxy, definitionsMapKey)
	}

	const parametersMapKey = "parameters"
	if raw, ok := proxy[parametersMapKey]; ok {
		var decoded ParameterMap
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Parameters`)
		}
		for key, elem := range decoded {
			mutator.Parameter(key, elem)
		}
		delete(proxy, parametersMapKey)
	}

	const responsesMapKey = "responses"
	if raw, ok := proxy[responsesMapKey]; ok {
		var decoded ResponseMap
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Responses`)
		}
		for key, elem := range decoded {
			mutator.Response(key, elem)
		}
		delete(proxy, responsesMapKey)
	}

	const securityDefinitionsMapKey = "securityDefinitions"
	if raw, ok := proxy[securityDefinitionsMapKey]; ok {
		var decoded SecuritySchemeMap
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field SecurityDefinitions`)
		}
		for key, elem := range decoded {
			mutator.SecurityDefinition(key, elem)
		}
		delete(proxy, securityDefinitionsMapKey)
	}

	const securityMapKey = "security"
	if raw, ok := proxy[securityMapKey]; ok {
		var decoded SecurityRequirementList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Security`)
		}
		for _, elem := range decoded {
			mutator.Security(elem)
		}
		delete(proxy, securityMapKey)
	}

	const tagsMapKey = "tags"
	if raw, ok := proxy[tagsMapKey]; ok {
		var decoded TagList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Tags`)
		}
		for _, elem := range decoded {
			mutator.Tag(elem)
		}
		delete(proxy, tagsMapKey)
	}

	const externalDocsMapKey = "externalDocs"

	if raw, ok := proxy[externalDocsMapKey]; ok {
		var decoded externalDocumentation
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ExternalDocs`)
		}

		mutator.ExternalDocs(&decoded)
		delete(proxy, externalDocsMapKey)
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

func (v *swagger) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "swagger":
		target = v.version
	case "info":
		target = v.info
	case "host":
		target = v.host
	case "basePath":
		target = v.basePath
	case "schemes":
		target = v.schemes
	case "consumes":
		target = v.consumes
	case "produces":
		target = v.produces
	case "paths":
		target = v.paths
	case "definitions":
		target = v.definitions
	case "parameters":
		target = v.parameters
	case "responses":
		target = v.responses
	case "securityDefinitions":
		target = v.securityDefinitions
	case "security":
		target = v.security
	case "tags":
		target = v.tags
	case "externalDocs":
		target = v.externalDocs
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

// SwaggerFromJSON constructs a Swagger from JSON buffer
func SwaggerFromJSON(buf []byte, v *Swagger) error {
	var tmp swagger
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}
	*v = &tmp
	return nil
}
