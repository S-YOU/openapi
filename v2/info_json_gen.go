package openapi

// This file was automatically generated by gentypes.go
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
var _ = fmt.Fprintf
var _ = log.Printf
var _ = strconv.ParseInt
var _ = errors.Cause

type infoMarshalProxy struct {
	Reference      string  `json:"$ref,omitempty"`
	Title          string  `json:"title"`
	Version        string  `json:"version"`
	Description    string  `json:"description"`
	TermsOfService string  `json:"termsOfService"`
	Contact        Contact `json:"contact"`
	License        License `json:"license"`
}

func (v *info) MarshalJSON() ([]byte, error) {
	var proxy infoMarshalProxy
	if s := v.reference; len(s) > 0 {
		return []byte(fmt.Sprintf(refOnlyTmpl, strconv.Quote(s))), nil
	}
	proxy.Title = v.title
	proxy.Version = v.version
	proxy.Description = v.description
	proxy.TermsOfService = v.termsOfService
	proxy.Contact = v.contact
	proxy.License = v.license
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

func (v *info) UnmarshalJSON(data []byte) error {
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

	mutator := MutateInfo(v)

	if raw := proxy["title"]; len(raw) > 0 {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field title`)
		}
		mutator.Title(decoded)
		delete(proxy, "title")
	}

	if raw := proxy["version"]; len(raw) > 0 {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field version`)
		}
		mutator.Version(decoded)
		delete(proxy, "version")
	}

	if raw := proxy["description"]; len(raw) > 0 {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field description`)
		}
		mutator.Description(decoded)
		delete(proxy, "description")
	}

	if raw := proxy["termsOfService"]; len(raw) > 0 {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field termsOfService`)
		}
		mutator.TermsOfService(decoded)
		delete(proxy, "termsOfService")
	}

	if raw := proxy["contact"]; len(raw) > 0 {
		var decoded contact
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Contact`)
		}

		mutator.Contact(&decoded)
		delete(proxy, "contact")
	}

	if raw := proxy["license"]; len(raw) > 0 {
		var decoded license
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field License`)
		}

		mutator.License(&decoded)
		delete(proxy, "license")
	}

	for name, raw := range proxy {
		if strings.HasPrefix(name, `x-`) {
			var ext interface{}
			if err := json.Unmarshal(raw, &ext); err != nil {
				return errors.Wrapf(err, `failed to unmarshal field %s`, name)
			}
			mutator.Extension(name, ext)
		}
	}

	if err := mutator.Do(); err != nil {
		return errors.Wrap(err, `failed to  unmarshal JSON`)
	}
	return nil
}

func (v *info) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "title":
		target = v.title
	case "version":
		target = v.version
	case "description":
		target = v.description
	case "termsOfService":
		target = v.termsOfService
	case "contact":
		target = v.contact
	case "license":
		target = v.license
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

func InfoFromJSON(buf []byte, v *Info) error {
	var tmp info
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}
	*v = &tmp
	return nil
}
