package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T20:40:48+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sort"
)

var _ = sort.Strings

func (v *xml) Name() string {
	return v.name
}

func (v *xml) Namespace() string {
	return v.namespace
}

func (v *xml) Prefix() string {
	return v.prefix
}

func (v *xml) Attribute() bool {
	return v.attribute
}

func (v *xml) Wrapped() bool {
	return v.wrapped
}

func (v *xml) Reference() string {
	return v.reference
}

func (v *xml) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *xml) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *xml) Validate() error {
	return nil
}
