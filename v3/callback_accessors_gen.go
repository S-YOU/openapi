package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T19:11:33+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *callback) Name() string {
	return v.name
}

func (v *callback) Urls() map[string]PathItem {
	return v.urls
}

func (v *callback) Reference() string {
	return v.reference
}

func (v *callback) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
