package openapi_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	openapi "github.com/lestrrat-go/openapi/v3"
	"github.com/stretchr/testify/assert"
)

func TestOpenAPI(t *testing.T) {
	root := filepath.Join("..", "spec", "examples", "v3.0")
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if !assert.NoError(t, err, `reading from %s should succeed`, path) {
			return err
		}
		//		t.Logf("%s", data)
		buf := bytes.NewBuffer(data)
		rdr := bytes.NewReader(buf.Bytes())

		t.Run(path, func(t *testing.T) {
			spec, err := openapi.ParseYAML(rdr)
			if !assert.NoError(t, err, `Decoding spec should succeed`) {
				rdr.Seek(0, 0)
				scanner := bufio.NewScanner(rdr)
				lineno := 1
				for scanner.Scan() {
					t.Logf("%4d: %s", lineno, scanner.Text())
					lineno++
				}
				return
			}
			_ = spec
		})

		return nil
	})
}
