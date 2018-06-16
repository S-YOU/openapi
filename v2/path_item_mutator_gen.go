package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"log"
)

var _ = log.Printf

// PathItemMutator is used to build an instance of PathItem. The user must
// call `Do()` after providing all the necessary information to
// the new instance of PathItem with new values
type PathItemMutator struct {
	proxy  *pathItem
	target *pathItem
}

// Do finalizes the matuation process for PathItem and returns the result
func (m *PathItemMutator) Do() error {
	*m.target = *m.proxy
	return nil
}

// MutatePathItem creates a new mutator object for PathItem
func MutatePathItem(v PathItem) *PathItemMutator {
	return &PathItemMutator{
		target: v.(*pathItem),
		proxy:  v.Clone().(*pathItem),
	}
}

// ClearParameters clears all elements in parameters
func (m *PathItemMutator) ClearParameters() *PathItemMutator {
	_ = m.proxy.parameters.Clear()
	return m
}

// Parameter appends a value to parameters
func (m *PathItemMutator) Parameter(value Parameter) *PathItemMutator {
	m.proxy.parameters = append(m.proxy.parameters, value)
	return m
}

// Extension sets an arbitrary extension field in PathItem
func (m *PathItemMutator) Extension(name string, value interface{}) *PathItemMutator {
	if m.proxy.extensions == nil {
		m.proxy.extensions = Extensions{}
	}
	m.proxy.extensions[name] = value
	return m
}
