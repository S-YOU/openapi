package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"
	"sort"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = sort.Strings
var _ = errors.Cause

func (v *securityRequirement) Name() string {
	return v.name
}

func (v *securityRequirement) Scopes() *ScopesMapIterator {
	var keys []string
	for key := range v.scopes {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.scopes[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ScopesMapIterator
	iter.list.items = items
	return &iter
}

// Reference returns the value of `$ref` field
func (v *securityRequirement) Reference() string {
	return v.reference
}

func (v *securityRequirement) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

// Extension returns the value of an arbitrary extension
func (v *securityRequirement) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

// Extensions return an iterator to iterate over all extensions
func (v *securityRequirement) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *securityRequirement) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
