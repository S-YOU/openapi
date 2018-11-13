package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sync"
)

// ContactMutator is used to build an instance of Contact. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Contact with new values
type ContactMutator struct {
	mu     sync.Mutex
	proxy  *contact
	target *contact
}

// Do finalizes the matuation process for Contact and returns the result
func (m *ContactMutator) Do() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	*m.target = *m.proxy
	return nil
}

// MutateContact creates a new mutator object for Contact
// Operations on the mutator are safe to be used concurrently, except for
// when calling `Do()`, where the user is responsible for restricting access
// to the target object to be mutated
func MutateContact(v Contact) *ContactMutator {
	return &ContactMutator{
		target: v.(*contact),
		proxy:  v.Clone().(*contact),
	}
}

// Name sets the Name field for object Contact.
func (m *ContactMutator) Name(v string) *ContactMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.name = v
	return m
}

// URL sets the URL field for object Contact.
func (m *ContactMutator) URL(v string) *ContactMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.url = v
	return m
}

// Email sets the Email field for object Contact.
func (m *ContactMutator) Email(v string) *ContactMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.email = v
	return m
}

// Extension sets an arbitrary extension field in Contact
func (m *ContactMutator) Extension(name string, value interface{}) *ContactMutator {
	if m.proxy.extensions == nil {
		m.proxy.extensions = Extensions{}
	}
	m.proxy.extensions[name] = value
	return m
}
