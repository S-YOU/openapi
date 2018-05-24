package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T19:11:33+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sync"
)

type mapIteratorItem struct {
	item interface{}
	key  interface{}
}

type mapIterator struct {
	list listIterator
}

func (iter *mapIterator) Next() bool {
	return iter.list.Next()
}

func (iter *mapIterator) Item() *mapIteratorItem {
	v := iter.list.Item()
	if v == nil {
		return nil
	}
	return v.(*mapIteratorItem)
}

type listIterator struct {
	mu    sync.RWMutex
	items []interface{}
}

// Item returns the next item in this iterator
func (iter *listIterator) Item() interface{} {
	iter.mu.Lock()
	defer iter.mu.Unlock()

	if !iter.nextNoLock() {
		return nil
	}

	item := iter.items[0]
	iter.items = iter.items[1:]
	return item
}

func (iter *listIterator) nextNoLock() bool {
	return len(iter.items) > 0
}

// Next returns true if there are more elements in this iterator
func (iter *listIterator) Next() bool {
	iter.mu.RLock()
	defer iter.mu.RUnlock()
	return iter.nextNoLock()
}

type InterfaceListIterator struct {
	listIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *InterfaceListIterator) Item() interface{} {
	return iter.listIterator.Item().(interface{})
}

type StringListIterator struct {
	listIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *StringListIterator) Item() string {
	return iter.listIterator.Item().(string)
}

type StringListListIterator struct {
	listIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *StringListListIterator) Item() []string {
	return iter.listIterator.Item().([]string)
}

type OperationListIterator struct {
	listIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *OperationListIterator) Item() Operation {
	return iter.listIterator.Item().(Operation)
}

type PathItemListIterator struct {
	listIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *PathItemListIterator) Item() PathItem {
	return iter.listIterator.Item().(PathItem)
}

type MediaTypeListIterator struct {
	listIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *MediaTypeListIterator) Item() MediaType {
	return iter.listIterator.Item().(MediaType)
}

type CallbackMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *CallbackMapIterator) Item() (string, Callback) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", Callback(nil)
	}
	return item.key.(string), item.item.(Callback)
}

type EncodingMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *EncodingMapIterator) Item() (string, Encoding) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", Encoding(nil)
	}
	return item.key.(string), item.item.(Encoding)
}

type ExampleMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *ExampleMapIterator) Item() (string, Example) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", Example(nil)
	}
	return item.key.(string), item.item.(Example)
}

type HeaderMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *HeaderMapIterator) Item() (string, Header) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", Header(nil)
	}
	return item.key.(string), item.item.(Header)
}

type InterfaceMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *InterfaceMapIterator) Item() (string, interface{}) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", interface{}(nil)
	}
	return item.key.(string), item.item.(interface{})
}

type LinkMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *LinkMapIterator) Item() (string, Link) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", Link(nil)
	}
	return item.key.(string), item.item.(Link)
}

type MediaTypeMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *MediaTypeMapIterator) Item() (string, MediaType) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", MediaType(nil)
	}
	return item.key.(string), item.item.(MediaType)
}

type PathItemMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *PathItemMapIterator) Item() (string, PathItem) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", PathItem(nil)
	}
	return item.key.(string), item.item.(PathItem)
}

type ParameterListIterator struct {
	listIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *ParameterListIterator) Item() Parameter {
	return iter.listIterator.Item().(Parameter)
}

type ParameterMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *ParameterMapIterator) Item() (string, Parameter) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", Parameter(nil)
	}
	return item.key.(string), item.item.(Parameter)
}

type RequestBodyMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *RequestBodyMapIterator) Item() (string, RequestBody) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", RequestBody(nil)
	}
	return item.key.(string), item.item.(RequestBody)
}

type ResponseMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *ResponseMapIterator) Item() (string, Response) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", Response(nil)
	}
	return item.key.(string), item.item.(Response)
}

type SchemaListIterator struct {
	listIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *SchemaListIterator) Item() Schema {
	return iter.listIterator.Item().(Schema)
}

type SchemaMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *SchemaMapIterator) Item() (string, Schema) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", Schema(nil)
	}
	return item.key.(string), item.item.(Schema)
}

type ScopeMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *ScopeMapIterator) Item() (string, string) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", ""
	}
	return item.key.(string), item.item.(string)
}

type SecurityRequirementListIterator struct {
	listIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *SecurityRequirementListIterator) Item() SecurityRequirement {
	return iter.listIterator.Item().(SecurityRequirement)
}

type SecuritySchemeMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *SecuritySchemeMapIterator) Item() (string, SecurityScheme) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", SecurityScheme(nil)
	}
	return item.key.(string), item.item.(SecurityScheme)
}

type ServerListIterator struct {
	listIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *ServerListIterator) Item() Server {
	return iter.listIterator.Item().(Server)
}

type ServerVariableMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *ServerVariableMapIterator) Item() (string, ServerVariable) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", ServerVariable(nil)
	}
	return item.key.(string), item.item.(ServerVariable)
}

type StringMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *StringMapIterator) Item() (string, string) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", ""
	}
	return item.key.(string), item.item.(string)
}

type StringListMapIterator struct {
	mapIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *StringListMapIterator) Item() (string, []string) {
	item := iter.mapIterator.Item()
	if item == nil {
		return "", []string(nil)
	}
	return item.key.(string), item.item.([]string)
}

type TagListIterator struct {
	listIterator
}

// Item returns the next item in this iterator. Make sure to call Next()
// before hand to check if the iterator has more items
func (iter *TagListIterator) Item() Tag {
	return iter.listIterator.Item().(Tag)
}
