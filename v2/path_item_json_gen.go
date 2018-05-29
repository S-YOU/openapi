package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T10:54:12+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
)

var _ = json.Unmarshal
var _ = log.Printf
var _ = errors.Cause

type pathItemMarshalProxy struct {
	Reference  string        `json:"$ref,omitempty"`
	Get        Operation     `json:"get,omitempty"`
	Put        Operation     `json:"put,omitempty"`
	Post       Operation     `json:"post,omitempty"`
	Delete     Operation     `json:"delete,omitempty"`
	Options    Operation     `json:"options,omitempty"`
	Head       Operation     `json:"head,omitempty"`
	Patch      Operation     `json:"patch,omitempty"`
	Parameters ParameterList `json:"parameters,omitempty"`
}

func (v *pathItem) MarshalJSON() ([]byte, error) {
	var proxy pathItemMarshalProxy
	if s := v.reference; len(s) > 0 {
		return []byte(fmt.Sprintf(refOnlyTmpl, strconv.Quote(s))), nil
	}
	proxy.Get = v.get
	proxy.Put = v.put
	proxy.Post = v.post
	proxy.Delete = v.delete
	proxy.Options = v.options
	proxy.Head = v.head
	proxy.Patch = v.patch
	proxy.Parameters = v.parameters
	buf, err := json.Marshal(proxy)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal struct`)
	}
	if len(v.extensions) > 0 {
		extBuf, err := json.Marshal(v.extensions)
		if err != nil || len(extBuf) <= 2 {
			return nil, errors.Wrap(err, `failed to marshal struct (extensions)`)
		}
		buf = append(append(buf[:len(buf)-1], ','), extBuf[1:]...)
	}
	return buf, nil
}

func (v *pathItem) UnmarshalJSON(data []byte) error {
	var proxy map[string]json.RawMessage
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if raw := proxy["$ref"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.reference); err != nil {
			return errors.Wrap(err, `failed to unmarshal $ref`)
		}
		return nil
	}

	mutator := MutatePathItem(v)

	if raw := proxy["get"]; len(raw) > 0 {
		var decoded operation
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Get`)
		}

		mutator.Get(&decoded)
		delete(proxy, "get")
	}

	if raw := proxy["put"]; len(raw) > 0 {
		var decoded operation
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Put`)
		}

		mutator.Put(&decoded)
		delete(proxy, "put")
	}

	if raw := proxy["post"]; len(raw) > 0 {
		var decoded operation
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Post`)
		}

		mutator.Post(&decoded)
		delete(proxy, "post")
	}

	if raw := proxy["delete"]; len(raw) > 0 {
		var decoded operation
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Delete`)
		}

		mutator.Delete(&decoded)
		delete(proxy, "delete")
	}

	if raw := proxy["options"]; len(raw) > 0 {
		var decoded operation
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Options`)
		}

		mutator.Options(&decoded)
		delete(proxy, "options")
	}

	if raw := proxy["head"]; len(raw) > 0 {
		var decoded operation
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Head`)
		}

		mutator.Head(&decoded)
		delete(proxy, "head")
	}

	if raw := proxy["patch"]; len(raw) > 0 {
		var decoded operation
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Patch`)
		}

		mutator.Patch(&decoded)
		delete(proxy, "patch")
	}

	if raw := proxy["parameters"]; len(raw) > 0 {
		var decoded ParameterList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Parameters`)
		}
		for _, elem := range decoded {
			mutator.Parameter(elem)
		}
		delete(proxy, "parameters")
	}

	for name, raw := range proxy {
		if strings.HasPrefix(name, `x-`) {
			mutator.Extension(name, raw)
		}
	}

	if err := mutator.Do(); err != nil {
		return errors.Wrap(err, `failed to  unmarshal JSON`)
	}
	return nil
}

func (v *pathItem) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*pathItem)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type PathItem, but got %T`, resolved)
		}
		mutator := MutatePathItem(v)
		mutator.Get(asserted.Get())
		mutator.Put(asserted.Put())
		mutator.Post(asserted.Post())
		mutator.Delete(asserted.Delete())
		mutator.Options(asserted.Options())
		mutator.Head(asserted.Head())
		mutator.Patch(asserted.Patch())
		for iter := asserted.Parameters(); iter.Next(); {
			item := iter.Item()
			mutator.Parameter(item)
		}
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.get != nil {
		if err := v.get.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Get`)
		}
	}
	if v.put != nil {
		if err := v.put.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Put`)
		}
	}
	if v.post != nil {
		if err := v.post.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Post`)
		}
	}
	if v.delete != nil {
		if err := v.delete.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Delete`)
		}
	}
	if v.options != nil {
		if err := v.options.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Options`)
		}
	}
	if v.head != nil {
		if err := v.head.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Head`)
		}
	}
	if v.patch != nil {
		if err := v.patch.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Patch`)
		}
	}
	if v.parameters != nil {
		if err := v.parameters.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Parameters`)
		}
	}
	return nil
}

func (v *pathItem) QueryJSON(path string) (ret interface{}, ok bool) {
	path = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })
	if path == "" {
		return v, true
	}

	var frag string
	if i := strings.Index(path, "/"); i > -1 {
		frag = path[:i]
		path = path[i+1:]
	} else {
		frag = path
		path = ""
	}

	var target interface{}

	switch frag {
	case "get":
		target = v.get
	case "put":
		target = v.put
	case "post":
		target = v.post
	case "delete":
		target = v.delete
	case "options":
		target = v.options
	case "head":
		target = v.head
	case "patch":
		target = v.patch
	case "parameters":
		target = v.parameters
	default:
		return nil, false
	}

	if qj, ok := target.(QueryJSONer); ok {
		return qj.QueryJSON(path)
	}
	if path == "" {
		return target, true
	}
	return nil, false
}
