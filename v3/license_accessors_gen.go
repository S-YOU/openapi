package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T19:11:33+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *license) Name() string {
	return v.name
}

func (v *license) URL() string {
	return v.uRL
}

func (v *license) Reference() string {
	return v.reference
}

func (v *license) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
