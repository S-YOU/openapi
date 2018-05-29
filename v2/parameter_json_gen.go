package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T10:54:12+09:00
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
var _ = log.Printf
var _ = errors.Cause

type parameterMarshalProxy struct {
	Reference        string           `json:"$ref,omitempty"`
	Name             string           `json:"name"`
	Description      string           `json:"description,omitempty"`
	Required         bool             `json:"required,omitempty"`
	In               Location         `json:"in"`
	Schema           Schema           `json:"schema,omitempty"`
	Type             PrimitiveType    `json:"type,omitempty"`
	Format           string           `json:"format,omitempty"`
	Title            string           `json:"title,omitempty"`
	AllowEmptyValue  bool             `json:"allowEmptyValue,omitempty"`
	Items            Items            `json:"items,omitempty"`
	CollectionFormat CollectionFormat `json:"collectionFormat,omitempty"`
	DefaultValue     interface{}      `json:"default,omitempty"`
	Maximum          float64          `json:"maximum,omitempty"`
	ExclusiveMaximum float64          `json:"exclusiveMaximum,omitempty"`
	Minimum          float64          `json:"minimum,omitempty"`
	ExclusiveMinimum float64          `json:"exclusiveMinimum,omitempty"`
	MaxLength        int              `json:"maxLength,omitempty"`
	MinLength        int              `json:"minLength,omitempty"`
	Pattern          string           `json:"pattern,omitempty"`
	MaxItems         int              `json:"maxItems,omitempty"`
	MinItems         int              `json:"minItems,omitempty"`
	UniqueItems      bool             `json:"uniqueItems,omitempty"`
	Enum             InterfaceList    `json:"enum,omitempty"`
	MultipleOf       float64          `json:"multipleOf,omitempty"`
}

func (v *parameter) MarshalJSON() ([]byte, error) {
	var proxy parameterMarshalProxy
	if s := v.reference; len(s) > 0 {
		return []byte(fmt.Sprintf(refOnlyTmpl, strconv.Quote(s))), nil
	}
	proxy.Name = v.name
	proxy.Description = v.description
	proxy.Required = v.required
	proxy.In = v.in
	proxy.Schema = v.schema
	proxy.Type = v.typ
	proxy.Format = v.format
	proxy.Title = v.title
	proxy.AllowEmptyValue = v.allowEmptyValue
	proxy.Items = v.items
	proxy.CollectionFormat = v.collectionFormat
	proxy.DefaultValue = v.defaultValue
	proxy.Maximum = v.maximum
	proxy.ExclusiveMaximum = v.exclusiveMaximum
	proxy.Minimum = v.minimum
	proxy.ExclusiveMinimum = v.exclusiveMinimum
	proxy.MaxLength = v.maxLength
	proxy.MinLength = v.minLength
	proxy.Pattern = v.pattern
	proxy.MaxItems = v.maxItems
	proxy.MinItems = v.minItems
	proxy.UniqueItems = v.uniqueItems
	proxy.Enum = v.enum
	proxy.MultipleOf = v.multipleOf
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

func (v *parameter) UnmarshalJSON(data []byte) error {
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

	mutator := MutateParameter(v)

	if raw := proxy["name"]; len(raw) > 0 {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field name`)
		}
		mutator.Name(decoded)
		delete(proxy, "name")
	}

	if raw := proxy["description"]; len(raw) > 0 {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field description`)
		}
		mutator.Description(decoded)
		delete(proxy, "description")
	}

	if raw := proxy["required"]; len(raw) > 0 {
		var decoded bool
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field required`)
		}
		mutator.Required(decoded)
		delete(proxy, "required")
	}

	if raw := proxy["in"]; len(raw) > 0 {
		var decoded Location
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field in`)
		}
		mutator.In(decoded)
		delete(proxy, "in")
	}

	if raw := proxy["schema"]; len(raw) > 0 {
		var decoded schema
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Schema`)
		}

		mutator.Schema(&decoded)
		delete(proxy, "schema")
	}

	if raw := proxy["type"]; len(raw) > 0 {
		var decoded PrimitiveType
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field type`)
		}
		mutator.Type(decoded)
		delete(proxy, "type")
	}

	if raw := proxy["format"]; len(raw) > 0 {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field format`)
		}
		mutator.Format(decoded)
		delete(proxy, "format")
	}

	if raw := proxy["title"]; len(raw) > 0 {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field title`)
		}
		mutator.Title(decoded)
		delete(proxy, "title")
	}

	if raw := proxy["allowEmptyValue"]; len(raw) > 0 {
		var decoded bool
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field allowEmptyValue`)
		}
		mutator.AllowEmptyValue(decoded)
		delete(proxy, "allowEmptyValue")
	}

	if raw := proxy["items"]; len(raw) > 0 {
		var decoded items
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Items`)
		}

		mutator.Items(&decoded)
		delete(proxy, "items")
	}

	if raw := proxy["collectionFormat"]; len(raw) > 0 {
		var decoded CollectionFormat
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field collectionFormat`)
		}
		mutator.CollectionFormat(decoded)
		delete(proxy, "collectionFormat")
	}

	if raw := proxy["default"]; len(raw) > 0 {
		var decoded interface{}
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field default`)
		}
		mutator.DefaultValue(decoded)
		delete(proxy, "default")
	}

	if raw := proxy["maximum"]; len(raw) > 0 {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field maximum`)
		}
		mutator.Maximum(decoded)
		delete(proxy, "maximum")
	}

	if raw := proxy["exclusiveMaximum"]; len(raw) > 0 {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field exclusiveMaximum`)
		}
		mutator.ExclusiveMaximum(decoded)
		delete(proxy, "exclusiveMaximum")
	}

	if raw := proxy["minimum"]; len(raw) > 0 {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field minimum`)
		}
		mutator.Minimum(decoded)
		delete(proxy, "minimum")
	}

	if raw := proxy["exclusiveMinimum"]; len(raw) > 0 {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field exclusiveMinimum`)
		}
		mutator.ExclusiveMinimum(decoded)
		delete(proxy, "exclusiveMinimum")
	}

	if raw := proxy["maxLength"]; len(raw) > 0 {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field maxLength`)
		}
		mutator.MaxLength(decoded)
		delete(proxy, "maxLength")
	}

	if raw := proxy["minLength"]; len(raw) > 0 {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field minLength`)
		}
		mutator.MinLength(decoded)
		delete(proxy, "minLength")
	}

	if raw := proxy["pattern"]; len(raw) > 0 {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field pattern`)
		}
		mutator.Pattern(decoded)
		delete(proxy, "pattern")
	}

	if raw := proxy["maxItems"]; len(raw) > 0 {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field maxItems`)
		}
		mutator.MaxItems(decoded)
		delete(proxy, "maxItems")
	}

	if raw := proxy["minItems"]; len(raw) > 0 {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field minItems`)
		}
		mutator.MinItems(decoded)
		delete(proxy, "minItems")
	}

	if raw := proxy["uniqueItems"]; len(raw) > 0 {
		var decoded bool
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field uniqueItems`)
		}
		mutator.UniqueItems(decoded)
		delete(proxy, "uniqueItems")
	}

	if raw := proxy["enum"]; len(raw) > 0 {
		var decoded InterfaceList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Enum`)
		}
		for _, elem := range decoded {
			mutator.Enum(elem)
		}
		delete(proxy, "enum")
	}

	if raw := proxy["multipleOf"]; len(raw) > 0 {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field multipleOf`)
		}
		mutator.MultipleOf(decoded)
		delete(proxy, "multipleOf")
	}

	for name, raw := range proxy {
		if strings.HasPrefix(name, `x-`) {
			mutator.Extension(name, raw)
		}
	}

	if err := mutator.Do(); err != nil {
		return errors.Wrap(err, `failed to  unmarshal JSON`)
	}
	return nil
}

func (v *parameter) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*parameter)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type Parameter, but got %T`, resolved)
		}
		mutator := MutateParameter(v)
		mutator.Name(asserted.Name())
		mutator.Description(asserted.Description())
		mutator.Required(asserted.Required())
		mutator.In(asserted.In())
		mutator.Schema(asserted.Schema())
		mutator.Type(asserted.Type())
		mutator.Format(asserted.Format())
		mutator.Title(asserted.Title())
		mutator.AllowEmptyValue(asserted.AllowEmptyValue())
		mutator.Items(asserted.Items())
		mutator.CollectionFormat(asserted.CollectionFormat())
		mutator.DefaultValue(asserted.DefaultValue())
		mutator.Maximum(asserted.Maximum())
		mutator.ExclusiveMaximum(asserted.ExclusiveMaximum())
		mutator.Minimum(asserted.Minimum())
		mutator.ExclusiveMinimum(asserted.ExclusiveMinimum())
		mutator.MaxLength(asserted.MaxLength())
		mutator.MinLength(asserted.MinLength())
		mutator.Pattern(asserted.Pattern())
		mutator.MaxItems(asserted.MaxItems())
		mutator.MinItems(asserted.MinItems())
		mutator.UniqueItems(asserted.UniqueItems())
		for iter := asserted.Enum(); iter.Next(); {
			item := iter.Item()
			mutator.Enum(item)
		}
		mutator.MultipleOf(asserted.MultipleOf())
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.schema != nil {
		if err := v.schema.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Schema`)
		}
	}
	if v.items != nil {
		if err := v.items.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Items`)
		}
	}
	if v.enum != nil {
		if err := v.enum.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Enum`)
		}
	}
	return nil
}

func (v *parameter) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "name":
		target = v.name
	case "description":
		target = v.description
	case "required":
		target = v.required
	case "in":
		target = v.in
	case "schema":
		target = v.schema
	case "type":
		target = v.typ
	case "format":
		target = v.format
	case "title":
		target = v.title
	case "allowEmptyValue":
		target = v.allowEmptyValue
	case "items":
		target = v.items
	case "collectionFormat":
		target = v.collectionFormat
	case "default":
		target = v.defaultValue
	case "maximum":
		target = v.maximum
	case "exclusiveMaximum":
		target = v.exclusiveMaximum
	case "minimum":
		target = v.minimum
	case "exclusiveMinimum":
		target = v.exclusiveMinimum
	case "maxLength":
		target = v.maxLength
	case "minLength":
		target = v.minLength
	case "pattern":
		target = v.pattern
	case "maxItems":
		target = v.maxItems
	case "minItems":
		target = v.minItems
	case "uniqueItems":
		target = v.uniqueItems
	case "enum":
		target = v.enum
	case "multipleOf":
		target = v.multipleOf
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
