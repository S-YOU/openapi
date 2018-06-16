package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
	"sort"
)

var _ = sort.Strings
var _ = errors.Cause

func (v *contact) Name() string {
	return v.name
}

func (v *contact) URL() string {
	return v.url
}

func (v *contact) Email() string {
	return v.email
}

// Reference returns the value of `$ref` field
func (v *contact) Reference() string {
	return v.reference
}

func (v *contact) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

// Extension returns the value of an arbitrary extension
func (v *contact) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

// Extensions return an iterator to iterate over all extensions
func (v *contact) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *contact) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *contact) recurseValidate() error {
	return nil
}
