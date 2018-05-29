package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T20:40:48+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sort"
)

var _ = sort.Strings

func (v *response) Name() string {
	return v.name
}

func (v *response) Description() string {
	return v.description
}

func (v *response) Schema() Schema {
	return v.schema
}

func (v *response) Headers() *HeaderMapIterator {
	var keys []string
	for key := range v.headers {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.headers[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter HeaderMapIterator
	iter.list.items = items
	return &iter
}

func (v *response) Examples() *ExampleMapIterator {
	var keys []string
	for key := range v.examples {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.examples[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExampleMapIterator
	iter.list.items = items
	return &iter
}

func (v *response) Reference() string {
	return v.reference
}

func (v *response) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *response) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *response) Validate() error {
	return nil
}
