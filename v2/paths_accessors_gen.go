package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T16:03:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *paths) Paths() *PathItemMapIterator {
	var items []interface{}
	for key, item := range v.paths {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter PathItemMapIterator
	iter.list.items = items
	return &iter
}

func (v *paths) Reference() string {
	return v.reference
}

func (v *paths) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *paths) Validate() error {
	return nil
}