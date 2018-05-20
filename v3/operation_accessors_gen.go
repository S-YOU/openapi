package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:06:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *operation) Verb() string {
	return v.verb
}

func (v *operation) Tags() []string {
	return v.tags
}

func (v *operation) Summary() string {
	return v.summary
}

func (v *operation) Description() string {
	return v.description
}

func (v *operation) ExternalDocs() ExternalDocumentation {
	return v.externalDocs
}

func (v *operation) OperationID() string {
	return v.operationID
}

func (v *operation) Parameters() []Parameter {
	return v.parameters
}

func (v *operation) RequestBody() RequestBody {
	return v.requestBody
}

func (v *operation) Responses() Responses {
	return v.responses
}

func (v *operation) Callbacks() map[string]Callback {
	return v.callbacks
}

func (v *operation) Deprecated() bool {
	return v.deprecated
}

func (v *operation) Security() []SecurityRequirement {
	return v.security
}

func (v *operation) Servers() []Server {
	return v.servers
}
