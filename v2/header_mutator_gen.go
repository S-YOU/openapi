package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sync"
)

// HeaderMutator is used to build an instance of Header. The user must
// call `Apply()` after providing all the necessary information to
// the new instance of Header with new values
type HeaderMutator struct {
	lock   sync.Locker
	proxy  *header
	target *header
}

// Apply finalizes the matuation process for Header and returns the result
func (m *HeaderMutator) Apply() error {
	m.lock.Lock()
	defer m.lock.Unlock()
	*m.target = *m.proxy
	return nil
}

// MutateHeader creates a new mutator object for Header
// Operations on the mutator are safe to be used concurrently, except for
// when calling `Apply()`, where the user is responsible for restricting access
// to the target object to be mutated
func MutateHeader(v Header, options ...Option) *HeaderMutator {
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
	return &HeaderMutator{
		lock:   lock,
		target: v.(*header),
		proxy:  v.Clone().(*header),
	}
}

// Name sets the Name field for object Header.
func (m *HeaderMutator) Name(v string) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.name = v
	return m
}

// Description sets the Description field for object Header.
func (m *HeaderMutator) Description(v string) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.description = v
	return m
}

// Type sets the Type field for object Header.
func (m *HeaderMutator) Type(v string) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.typ = v
	return m
}

// Format sets the Format field for object Header.
func (m *HeaderMutator) Format(v string) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.format = v
	return m
}

// Items sets the Items field for object Header.
func (m *HeaderMutator) Items(v Items) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.items = v
	return m
}

// CollectionFormat sets the CollectionFormat field for object Header.
func (m *HeaderMutator) CollectionFormat(v CollectionFormat) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.collectionFormat = v
	return m
}

// Default sets the Default field for object Header.
func (m *HeaderMutator) Default(v interface{}) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.defaultValue = v
	return m
}

// Maximum sets the Maximum field for object Header.
func (m *HeaderMutator) Maximum(v float64) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.maximum = v
	return m
}

// ExclusiveMaximum sets the ExclusiveMaximum field for object Header.
func (m *HeaderMutator) ExclusiveMaximum(v float64) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.exclusiveMaximum = v
	return m
}

// Minimum sets the Minimum field for object Header.
func (m *HeaderMutator) Minimum(v float64) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.minimum = v
	return m
}

// ExclusiveMinimum sets the ExclusiveMinimum field for object Header.
func (m *HeaderMutator) ExclusiveMinimum(v float64) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.exclusiveMinimum = v
	return m
}

// MaxLength sets the MaxLength field for object Header.
func (m *HeaderMutator) MaxLength(v int) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.maxLength = v
	return m
}

// MinLength sets the MinLength field for object Header.
func (m *HeaderMutator) MinLength(v int) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.minLength = v
	return m
}

// Pattern sets the Pattern field for object Header.
func (m *HeaderMutator) Pattern(v string) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.pattern = v
	return m
}

// MaxItems sets the MaxItems field for object Header.
func (m *HeaderMutator) MaxItems(v int) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.maxItems = v
	return m
}

// MinItems sets the MinItems field for object Header.
func (m *HeaderMutator) MinItems(v int) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.minItems = v
	return m
}

// UniqueItems sets the UniqueItems field for object Header.
func (m *HeaderMutator) UniqueItems(v bool) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.uniqueItems = v
	return m
}

// ClearEnum clears all elements in enum
func (m *HeaderMutator) ClearEnum() *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	_ = m.proxy.enum.Clear()
	return m
}

// Enum appends a value to enum
func (m *HeaderMutator) Enum(value interface{}) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.enum = append(m.proxy.enum, value)
	return m
}

// MultipleOf sets the MultipleOf field for object Header.
func (m *HeaderMutator) MultipleOf(v float64) *HeaderMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.multipleOf = v
	return m
}

// Extension sets an arbitrary extension field in Header
func (m *HeaderMutator) Extension(name string, value interface{}) *HeaderMutator {
	if m.proxy.extensions == nil {
		m.proxy.extensions = Extensions{}
	}
	m.proxy.extensions[name] = value
	return m
}
