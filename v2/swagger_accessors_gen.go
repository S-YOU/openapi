package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
	"sort"
)

var _ = sort.Strings
var _ = errors.Cause

func (v *swagger) Version() string {
	return v.version
}

func (v *swagger) Info() Info {
	return v.info
}

func (v *swagger) Host() string {
	return v.host
}

func (v *swagger) BasePath() string {
	return v.basePath
}

func (v *swagger) Schemes() *SchemeListIterator {
	var items []interface{}
	for _, item := range v.schemes {
		items = append(items, item)
	}
	var iter SchemeListIterator
	iter.items = items
	return &iter
}

func (v *swagger) Consumes() *MIMETypeListIterator {
	var items []interface{}
	for _, item := range v.consumes {
		items = append(items, item)
	}
	var iter MIMETypeListIterator
	iter.items = items
	return &iter
}

func (v *swagger) Produces() *MIMETypeListIterator {
	var items []interface{}
	for _, item := range v.produces {
		items = append(items, item)
	}
	var iter MIMETypeListIterator
	iter.items = items
	return &iter
}

func (v *swagger) Paths() Paths {
	return v.paths
}

func (v *swagger) Definitions() *InterfaceMapIterator {
	var keys []string
	for key := range v.definitions {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.definitions[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter InterfaceMapIterator
	iter.list.items = items
	return &iter
}

func (v *swagger) Parameters() *ParameterMapIterator {
	var keys []string
	for key := range v.parameters {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.parameters[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ParameterMapIterator
	iter.list.items = items
	return &iter
}

func (v *swagger) Responses() *ResponseMapIterator {
	var keys []string
	for key := range v.responses {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.responses[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ResponseMapIterator
	iter.list.items = items
	return &iter
}

func (v *swagger) SecurityDefinitions() *SecuritySchemeMapIterator {
	var keys []string
	for key := range v.securityDefinitions {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.securityDefinitions[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter SecuritySchemeMapIterator
	iter.list.items = items
	return &iter
}

func (v *swagger) Security() *SecurityRequirementListIterator {
	var items []interface{}
	for _, item := range v.security {
		items = append(items, item)
	}
	var iter SecurityRequirementListIterator
	iter.items = items
	return &iter
}

func (v *swagger) Tags() *TagListIterator {
	var items []interface{}
	for _, item := range v.tags {
		items = append(items, item)
	}
	var iter TagListIterator
	iter.items = items
	return &iter
}

func (v *swagger) ExternalDocs() ExternalDocumentation {
	return v.externalDocs
}

// Reference returns the value of `$ref` field
func (v *swagger) Reference() string {
	return v.reference
}

func (v *swagger) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

// Extension returns the value of an arbitrary extension
func (v *swagger) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

// Extensions return an iterator to iterate over all extensions
func (v *swagger) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *swagger) recurseValidate() error {
	if elem := v.info; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "info"`)
		}
	}
	if elem := v.schemes; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "schemes"`)
		}
	}
	if elem := v.consumes; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "consumes"`)
		}
	}
	if elem := v.produces; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "produces"`)
		}
	}
	if elem := v.paths; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "paths"`)
		}
	}
	if elem := v.definitions; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "definitions"`)
		}
	}
	if elem := v.parameters; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "parameters"`)
		}
	}
	if elem := v.responses; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "responses"`)
		}
	}
	if elem := v.securityDefinitions; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "securityDefinitions"`)
		}
	}
	if elem := v.security; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "security"`)
		}
	}
	if elem := v.tags; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "tags"`)
		}
	}
	if elem := v.externalDocs; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "externalDocs"`)
		}
	}
	return nil
}
