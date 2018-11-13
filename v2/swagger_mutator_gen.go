package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sync"
)

// SwaggerMutator is used to build an instance of Swagger. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Swagger with new values
type SwaggerMutator struct {
	mu     sync.Mutex
	proxy  *swagger
	target *swagger
}

// Do finalizes the matuation process for Swagger and returns the result
func (m *SwaggerMutator) Do() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	*m.target = *m.proxy
	return nil
}

// MutateSwagger creates a new mutator object for Swagger
// Operations on the mutator are safe to be used concurrently, except for
// when calling `Do()`, where the user is responsible for restricting access
// to the target object to be mutated
func MutateSwagger(v Swagger) *SwaggerMutator {
	return &SwaggerMutator{
		target: v.(*swagger),
		proxy:  v.Clone().(*swagger),
	}
}

// Version sets the Version field for object Swagger.
func (m *SwaggerMutator) Version(v string) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.version = v
	return m
}

// Info sets the Info field for object Swagger.
func (m *SwaggerMutator) Info(v Info) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.info = v
	return m
}

// Host sets the Host field for object Swagger.
func (m *SwaggerMutator) Host(v string) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.host = v
	return m
}

// BasePath sets the BasePath field for object Swagger.
func (m *SwaggerMutator) BasePath(v string) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.basePath = v
	return m
}

// ClearSchemes clears all elements in schemes
func (m *SwaggerMutator) ClearSchemes() *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	_ = m.proxy.schemes.Clear()
	return m
}

// Scheme appends a value to schemes
func (m *SwaggerMutator) Scheme(value string) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.schemes = append(m.proxy.schemes, value)
	return m
}

// ClearConsumes clears all elements in consumes
func (m *SwaggerMutator) ClearConsumes() *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	_ = m.proxy.consumes.Clear()
	return m
}

// Consume appends a value to consumes
func (m *SwaggerMutator) Consume(value string) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.consumes = append(m.proxy.consumes, value)
	return m
}

// ClearProduces clears all elements in produces
func (m *SwaggerMutator) ClearProduces() *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	_ = m.proxy.produces.Clear()
	return m
}

// Produce appends a value to produces
func (m *SwaggerMutator) Produce(value string) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.produces = append(m.proxy.produces, value)
	return m
}

// Paths sets the Paths field for object Swagger.
func (m *SwaggerMutator) Paths(v Paths) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.paths = v
	return m
}

// ClearDefinitions removes all values in definitions field
func (m *SwaggerMutator) ClearDefinitions() *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	_ = m.proxy.definitions.Clear()
	return m
}

// Definition sets the value of definitions
func (m *SwaggerMutator) Definition(key InterfaceMapKey, value interface{}) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.proxy.definitions == nil {
		m.proxy.definitions = InterfaceMap{}
	}

	m.proxy.definitions[key] = value
	return m
}

// ClearParameters removes all values in parameters field
func (m *SwaggerMutator) ClearParameters() *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	_ = m.proxy.parameters.Clear()
	return m
}

// Parameter sets the value of parameters
func (m *SwaggerMutator) Parameter(key ParameterMapKey, value Parameter) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.proxy.parameters == nil {
		m.proxy.parameters = ParameterMap{}
	}

	m.proxy.parameters[key] = value
	return m
}

// ClearResponses removes all values in responses field
func (m *SwaggerMutator) ClearResponses() *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	_ = m.proxy.responses.Clear()
	return m
}

// Response sets the value of responses
func (m *SwaggerMutator) Response(key ResponseMapKey, value Response) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.proxy.responses == nil {
		m.proxy.responses = ResponseMap{}
	}

	m.proxy.responses[key] = value
	return m
}

// ClearSecurityDefinitions removes all values in securityDefinitions field
func (m *SwaggerMutator) ClearSecurityDefinitions() *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	_ = m.proxy.securityDefinitions.Clear()
	return m
}

// SecurityDefinition sets the value of securityDefinitions
func (m *SwaggerMutator) SecurityDefinition(key SecuritySchemeMapKey, value SecurityScheme) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.proxy.securityDefinitions == nil {
		m.proxy.securityDefinitions = SecuritySchemeMap{}
	}

	m.proxy.securityDefinitions[key] = value
	return m
}

// ClearSecurity clears all elements in security
func (m *SwaggerMutator) ClearSecurity() *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	_ = m.proxy.security.Clear()
	return m
}

// Security appends a value to security
func (m *SwaggerMutator) Security(value SecurityRequirement) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.security = append(m.proxy.security, value)
	return m
}

// ClearTags clears all elements in tags
func (m *SwaggerMutator) ClearTags() *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	_ = m.proxy.tags.Clear()
	return m
}

// Tag appends a value to tags
func (m *SwaggerMutator) Tag(value Tag) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.tags = append(m.proxy.tags, value)
	return m
}

// ExternalDocs sets the ExternalDocs field for object Swagger.
func (m *SwaggerMutator) ExternalDocs(v ExternalDocumentation) *SwaggerMutator {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.proxy.externalDocs = v
	return m
}

// Extension sets an arbitrary extension field in Swagger
func (m *SwaggerMutator) Extension(name string, value interface{}) *SwaggerMutator {
	if m.proxy.extensions == nil {
		m.proxy.extensions = Extensions{}
	}
	m.proxy.extensions[name] = value
	return m
}
