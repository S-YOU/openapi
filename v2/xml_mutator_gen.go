package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sync"
)

// XMLMutator is used to build an instance of XML. The user must
// call `Apply()` after providing all the necessary information to
// the new instance of XML with new values
type XMLMutator struct {
	lock   sync.Locker
	proxy  *xml
	target *xml
}

// Apply finalizes the matuation process for XML and returns the result
func (m *XMLMutator) Apply() error {
	m.lock.Lock()
	defer m.lock.Unlock()
	*m.target = *m.proxy
	return nil
}

// MutateXML creates a new mutator object for XML
// Operations on the mutator are safe to be used concurrently, except for
// when calling `Apply()`, where the user is responsible for restricting access
// to the target object to be mutated
func MutateXML(v XML, options ...Option) *XMLMutator {
	var lock sync.Locker = &sync.Mutex{}
	for _, option := range options {
		switch option.Name() {
		case optkeyLocker:
			lock = option.Value().(sync.Locker)
		}
	}
	if lock == nil {
		lock = nilLock{}
	}
	return &XMLMutator{
		lock:   lock,
		target: v.(*xml),
		proxy:  v.Clone().(*xml),
	}
}

// Name sets the Name field for object XML.
func (m *XMLMutator) Name(v string) *XMLMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.name = v
	return m
}

// Namespace sets the Namespace field for object XML.
func (m *XMLMutator) Namespace(v string) *XMLMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.namespace = v
	return m
}

// Prefix sets the Prefix field for object XML.
func (m *XMLMutator) Prefix(v string) *XMLMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.prefix = v
	return m
}

// Attribute sets the Attribute field for object XML.
func (m *XMLMutator) Attribute(v bool) *XMLMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.attribute = v
	return m
}

// Wrapped sets the Wrapped field for object XML.
func (m *XMLMutator) Wrapped(v bool) *XMLMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.wrapped = v
	return m
}

// Extension sets an arbitrary extension field in XML
func (m *XMLMutator) Extension(name string, value interface{}) *XMLMutator {
	if m.proxy.extensions == nil {
		m.proxy.extensions = Extensions{}
	}
	m.proxy.extensions[name] = value
	return m
}
