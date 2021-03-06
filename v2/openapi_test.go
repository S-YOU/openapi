package openapi_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ghodss/yaml"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

// encoding/json sorts the keys of map[string]*** data, but it doesn't
// do the same for struct-based data (which is understandable). Here,
// we force-sort the keys by marshaling twice -- the second time we
// shove the data into a map, and marshal that, allowing
// encoding/json to work its sorting magic.
//
// We format the value while we're at it
func sortMarshal(t *testing.T, v interface{}) ([]byte, error) {
	t.Helper()

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		return nil, errors.Wrap(err, `failed to marshal data`)
	}

	var m map[string]interface{}
	if err := json.NewDecoder(&buf).Decode(&m); err != nil {
		return nil, errors.Wrap(err, `failed to unmarshal data`)
	}

	return json.MarshalIndent(m, "", "  ")
}

func withLineno(t *testing.T, src io.Reader) {
	scanner := bufio.NewScanner(src)
	lineno := 1
	for scanner.Scan() {
		t.Logf("%4d: %s", lineno, scanner.Text())
		lineno++
	}
}

// These objects exist so that we can reuse them in tests later
var apiSupport openapi.Contact
var petsInfo openapi.Info
var petsLicense openapi.License
var petSchema openapi.Schema
var petListSchema openapi.Schema
var petsGetResponse openapi.Response
var petsResponses openapi.Responses
var petsGetOper openapi.Operation
var petsPathItem openapi.PathItem
var petsPaths openapi.Paths

func init() {
	apiSupport = openapi.NewContact().
		Name("API Support").
		URL("http://www.swagger.io/support").
		Email("support@swagger.io").
		MustBuild()

	petsLicense = openapi.NewLicense("Apache 2.0").
		URL("http://www.apache.org/licenses/LICENSE-2.0.html").
		MustBuild()

	petsInfo = openapi.NewInfo("Swagger Sample App", "1.0.1").
		Description("This is a sample server Petstore server.").
		TermsOfService("http://swagger.io/terms/").
		Contact(apiSupport).
		License(petsLicense).
		MustBuild()

	petSchema = openapi.NewSchema().
		Reference("#/definitions/pet").
		MustBuild()

	petListSchema = openapi.NewSchema().
		Type(openapi.Array).
		Items(petSchema).
		MustBuild()

	petsGetResponse = openapi.NewResponse("A list of pets.").
		Schema(petListSchema).
		MustBuild()

	petsResponses = openapi.NewResponses().
		Response("200", petsGetResponse).
		MustBuild()

	petsGetOper = openapi.NewOperation(petsResponses).
		Description("Returns all pets from the system that the user has access to").
		Produces("application/json").
		MustBuild()

	petsPathItem = openapi.NewPathItem().
		Get(petsGetOper).
		MustBuild()

	petsPaths = openapi.NewPaths().
		Path("/pets", petsPathItem).
		MustBuild()
}

func TestOpenAPI(t *testing.T) {
	root := filepath.Join("..", "spec", "examples", "v2.0")
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

			{
				encoded, _ := json.MarshalIndent(spec, "", "  ")
				withLineno(t, bytes.NewReader(encoded))
				return
			}
		})

		return nil
	})
}

func TestParseExtensions(t *testing.T) {
	srcFile := filepath.Join("..", "spec", "examples", "v2.0", "petstore-expanded.yaml")
	data, err := ioutil.ReadFile(srcFile)
	if !assert.NoError(t, err, `reading from %s should succeed`, srcFile) {
		return
	}
	var buf bytes.Buffer
	buf.Write(data)
	buf.WriteString("\nx-hello-world: Hello, World")

	spec, err := openapi.ParseYAML(&buf)
	if !assert.NoError(t, err, `parse YAML should succeed`) {
		return
	}

	encoded, err := yaml.Marshal(spec)
	if !assert.NoError(t, err, `yaml.Marshal should succeed`) {
		return
	}

	withLineno(t, bytes.NewReader(encoded))
	if !assert.True(t, bytes.Contains(encoded, []byte("x-hello-world: Hello, World")), "exntesion should exist") {
		return
	}
}

func TestValidate(t *testing.T) {
	_, err := openapi.NewSwagger(nil, nil).Build()
	if !assert.Error(t, err, "expected to see an error") {
		return
	}
	t.Logf("%s", err)
}

func TestAllOf(t *testing.T) {
	const src = `
swagger: "2.0"
info:
  version: 1.0.0
  title: Swagger Petstore
definitions:
  Pet:
    type: object
    discriminator: petType
    properties:
      name:
        type: string
      petType:
        type: string
    required:
    - name
    - petType
  Cat:
    description: A representation of a cat
    allOf:
    - $ref: '#/definitions/Pet'
    - type: object
      properties:
        huntingSkill:
          type: string
          description: The measured skill for hunting
          default: lazy
          enum:
          - clueless
          - lazy
          - adventurous
          - aggressive
      required:
      - huntingSkill
  Dog:
    description: A representation of a dog
    allOf:
    - $ref: '#/definitions/Pet'
    - type: object
      properties:
        packSize:
          type: integer
          format: int32
          description: the size of the pack the dog is from
          default: 0
          minimum: 0
      required:
      - packSize
paths: {}
`
	rdr := strings.NewReader(src)
	spec, err := openapi.ParseYAML(rdr)

	if !assert.NoError(t, err, `failed to parse source`) {
		return
	}

	t.Run("Check data", func(t *testing.T) {
		const expected = `{
  "definitions": {
    "Cat": {
      "allOf": [
        {
          "$ref": "#/definitions/Pet"
        },
        {
          "properties": {
            "huntingSkill": {
              "default": "lazy",
              "description": "The measured skill for hunting",
              "enum": [
                "clueless",
                "lazy",
                "adventurous",
                "aggressive"
              ],
              "type": "string"
            }
          },
          "required": [
            "huntingSkill"
          ],
          "type": "object"
        }
      ],
      "description": "A representation of a cat"
    },
    "Dog": {
      "allOf": [
        {
          "$ref": "#/definitions/Pet"
        },
        {
          "properties": {
            "packSize": {
              "default": 0,
              "description": "the size of the pack the dog is from",
              "format": "int32",
              "minimum": 0,
              "type": "integer"
            }
          },
          "required": [
            "packSize"
          ],
          "type": "object"
        }
      ],
      "description": "A representation of a dog"
    },
    "Pet": {
      "discriminator": "petType",
      "properties": {
        "name": {
          "type": "string"
        },
        "petType": {
          "type": "string"
        }
      },
      "required": [
        "name",
        "petType"
      ],
      "type": "object"
    }
  },
  "info": {
    "title": "Swagger Petstore",
    "version": "1.0.0"
  },
  "paths": {},
  "swagger": "2.0"
}`
		buf, err := sortMarshal(t, spec)
		if !assert.NoError(t, err, `marshaling should succeed`) {
			return
		}
		if !assert.Equal(t, []byte(expected), buf, "marshaled data should match") {
			return
		}
	})
}
