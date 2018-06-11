package codegen

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/lestrrat-go/openapi/internal/stringutil"
	"github.com/pkg/errors"
)

func DumpCode(dst io.Writer, src io.Reader) {
	scanner := bufio.NewScanner(src)
	lineno := 1
	for scanner.Scan() {
		fmt.Fprintf(dst, "%5d: %s\n", lineno, scanner.Text())
		lineno++
	}
}

func ExportedName(s string) string {
	switch s {
	case "defaultValue":
		return "Default"
	case "url":
		return "URL"
	case "xml":
		return "XML"
	case "typ":
		return "Type"
	}

	s = stringutil.Camel(stringutil.Snake(s))
	s = strings.Replace(s, "Id", "ID", -1)
	s = strings.Replace(s, "Url", "URL", -1)
	return s
}

func UnexportedName(s string) string {
	switch s {
	case "Default":
		return "defaultValue"
	case "URL":
		return "url"
	case "XML":
		return "xml"
	case "Type":
		return "typ"
	}

	s = stringutil.LcFirst(stringutil.Camel(stringutil.Snake(s)))
	s = strings.Replace(s, "Id", "ID", -1)
	s = strings.Replace(s, "Url", "URL", -1)
	return s
}

func WritePreamble(dst io.Writer, pkg string) error {
	fmt.Fprintf(dst, "\n\npackage %s", pkg)
	fmt.Fprintf(dst, "\n\n// This file was automatically generated.")
	fmt.Fprintf(dst, "\n// DO NOT EDIT MANUALLY. All changes will be lost\n")
	return nil
}

var importDummies = map[string]string{
	"github.com/pkg/errors": "errors.Cause",
	"encoding/json":         "json.Unmarshal",
	"log":                   "log.Printf",
	"sort":                  "sort.Strings",
}

func WriteImports(dst io.Writer, libs ...string) error {
	l := len(libs)
	if l == 0 {
		return nil
	}

	fmt.Fprintf(dst, "\n\nimport ")
	if l == 1 {
		fmt.Fprintf(dst, "%s", strconv.Quote(libs[0]))
		return nil
	}

	sort.Strings(libs)

	fmt.Fprintf(dst, "(")
	for _, lib := range libs {
		fmt.Fprintf(dst, "\n%s", strconv.Quote(lib))
	}
	fmt.Fprintf(dst, "\n)")

	// check to see if we need dummies
	var buf bytes.Buffer
	for _, lib := range libs {
		if v, ok := importDummies[lib]; ok {
			fmt.Fprintf(&buf, "\nvar _ = %s", v)
		}
	}

	if buf.Len() > 0 {
		fmt.Fprintf(dst, "\n")
		buf.WriteTo(dst)
	}
	return nil
}

func WriteFormattedToFile(fn string, code []byte) error {
	formatted, err := format.Source(code)
	if err != nil {
		return errors.Wrap(err, `failed to format source code`)
	}

	f, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, `failed to open file for writing`)
	}
	defer f.Close()

	if _, err := f.Write(formatted); err != nil {
		return errors.Wrap(err, `failed to write to file`)
	}
	return nil
}
