package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:06:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *pathItem) Path() string {
	return v.path
}

func (v *pathItem) Reference() string {
	return v.reference
}

func (v *pathItem) Summary() string {
	return v.summary
}

func (v *pathItem) Description() string {
	return v.description
}

func (v *pathItem) Get() Operation {
	return v.get
}

func (v *pathItem) Put() Operation {
	return v.put
}

func (v *pathItem) Post() Operation {
	return v.post
}

func (v *pathItem) Delete() Operation {
	return v.delete
}

func (v *pathItem) Options() Operation {
	return v.options
}

func (v *pathItem) Head() Operation {
	return v.head
}

func (v *pathItem) Patch() Operation {
	return v.patch
}

func (v *pathItem) Trace() Operation {
	return v.trace
}

func (v *pathItem) Servers() []Server {
	return v.servers
}

func (v *pathItem) Parameters() []Parameter {
	return v.parameters
}