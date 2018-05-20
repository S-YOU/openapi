package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T21:56:43+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// OpenAPIMutator is used to build an instance of OpenAPI. The user must
// call `Do()` after providing all the necessary information to
// the new instance of OpenAPI with new values
type OpenAPIMutator struct {
	proxy  *openAPI
	target *openAPI
}

// Do finalizes the matuation process for OpenAPI and returns the result
func (b *OpenAPIMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateOpenAPI creates a new mutator object for OpenAPI
func MutateOpenAPI(v OpenAPI) *OpenAPIMutator {
	return &OpenAPIMutator{
		target: v.(*openAPI),
		proxy:  v.Clone().(*openAPI),
	}
}

// Version sets the Version field for object OpenAPI.
func (b *OpenAPIMutator) Version(v string) *OpenAPIMutator {
	b.proxy.version = v
	return b
}

// Info sets the Info field for object OpenAPI.
func (b *OpenAPIMutator) Info(v Info) *OpenAPIMutator {
	b.proxy.info = v
	return b
}

// Servers sets the Servers field for object OpenAPI.
func (b *OpenAPIMutator) Servers(v []Server) *OpenAPIMutator {
	b.proxy.servers = v
	return b
}

// Paths sets the Paths field for object OpenAPI.
func (b *OpenAPIMutator) Paths(v Paths) *OpenAPIMutator {
	b.proxy.paths = v
	return b
}

// Components sets the Components field for object OpenAPI.
func (b *OpenAPIMutator) Components(v Components) *OpenAPIMutator {
	b.proxy.components = v
	return b
}

// Security sets the Security field for object OpenAPI.
func (b *OpenAPIMutator) Security(v SecurityRequirement) *OpenAPIMutator {
	b.proxy.security = v
	return b
}

// Tags sets the Tags field for object OpenAPI.
func (b *OpenAPIMutator) Tags(v []Tag) *OpenAPIMutator {
	b.proxy.tags = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object OpenAPI.
func (b *OpenAPIMutator) ExternalDocs(v ExternalDocumentation) *OpenAPIMutator {
	b.proxy.externalDocs = v
	return b
}
