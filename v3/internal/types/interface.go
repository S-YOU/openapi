// This file serves as the base template of openapi/v3/interface_gen.go.
// This is separated from the main openapi/v3 files because if we
// rely on code that resides in the same directory as where we are
// generating code, we risk messing up the code and not being able to
// run code generation tools again because of compile problems
package types

import "sync"

const (
	DefaultSpecVersion   = "0.0.1"
	DefaultVersion       = "3.0.1"
	defaultHeaderInValue = "header"
)

type Location string

const (
	InHeader Location = "header"
	InQuery  Location = "query"
	InPath   Location = "path"
	InCookie Location = "cookie"
)

type PrimitiveType string

const (
	Integer PrimitiveType = "integer"
	Number  PrimitiveType = "number"
	String  PrimitiveType = "string"
	Boolean PrimitiveType = "boolean"
	Object  PrimitiveType = "object"
	Array   PrimitiveType = "array"
	Null    PrimitiveType = "null"
)

type OpenAPI interface {
}

type openAPI struct {
	version      string                `json:"openapi" builder:"required" default:"DefaultVersion"`
	info         Info                  `json:"info" builder:"required"`
	servers      []Server              `json:"servers,omitempty"`
	paths        Paths                 `json:"paths" builder:"required"`
	components   Components            `json:"components,omitempty"`
	security     SecurityRequirement   `json:"security,omitempty"`
	tags         []Tag                 `json:"tags,omitempty"`
	externalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type Info interface {
}

type info struct {
	title          string  `json:"title" builder:"required"`
	description    string  `json:"description,omitempty"`
	termsOfService string  `json:"termsOfService,omitempty"`
	contact        Contact `json:"contact,omitempty"`
	license        License `json:"license,omitempty"`
	version        string  `json:"version" builder:"required" default:"DefaultSpecVersion"`
}

// Contact represents the contact object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.1.md#contactObject
type Contact interface {
}

type contact struct {
	name  string `json:"name,omitempty" yaml:"name,omitempty"`
	uRL   string `json:"url,omitempty" yaml:"url,omitempty"`
	email string `json:"email,omitempty" yaml:"email,omitempty"`
}

type License interface {
}

type license struct {
	name string `json:"name" builder:"required"`
	uRL  string `json:"url,omitempty"`
}

type Server interface {
}

type server struct {
	uRL         string                    `json:"url" builder:"required"`
	description string                    `json:"description,omitempty"`
	variables   map[string]ServerVariable `json:"variables,omitempty"`
}

type ServerVariable interface {
}

type serverVariable struct {
	enum         []string `json:"enum"`
	defaultValue string   `json:"default" builder:"required"`
	description  string   `json:"description"`
}

type Components interface {
}

type components struct {
	schemas         map[string]Schema         `json:"schemas,omitempty"`         // or Reference
	responses       map[string]Response       `json:"responses,omitempty"`       // or Reference
	parameters      map[string]Parameter      `json:"parameters,omitempty"`      // or Reference
	examples        map[string]Example        `json:"examples,omitempty"`        // or Reference
	requestBodies   map[string]RequestBody    `json:"requestBodies,omitempty"`   // or Reference
	headers         map[string]Header         `json:"headers,omitempty"`         // or Reference
	securitySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty"` // or Reference
	links           map[string]Link           `json:"links,omitempty"`           // or Reference
	callbacks       map[string]Callback       `json:"callbacks,omitempty"`       // or Reference
}

type PathItemIterator struct {
	mu    sync.RWMutex
	items []PathItem
}

type Paths interface {
	Items() *PathItemIterator
}

type paths struct {
	paths map[string]PathItem `json:"-"`
}

type OperationIterator struct {
	mu    sync.RWMutex
	items []Operation
}

type PathItem interface {
	Operations() *OperationIterator
	setPath(string)
}

type pathItem struct {
	path        string      `json:"-"` // This is a secret variable that gets reset when the item is added to a path
	reference   string      `json:"$ref,omitempty"`
	summary     string      `json:"summary,omitempty"`
	description string      `json:"description,omitempty"`
	get         Operation   `json:"get,omitempty"`
	put         Operation   `json:"put,omitempty"`
	post        Operation   `json:"post,omitempty"`
	delete      Operation   `json:"delete,omitempty"`
	options     Operation   `json:"options,omitempty"`
	head        Operation   `json:"head,omitempty"`
	patch       Operation   `json:"patch,omitempty"`
	trace       Operation   `json:"trace,omitempty"`
	servers     []Server    `json:"servers,omitempty"`
	parameters  []Parameter `json:"parameters,omitempty"` // or Reference
}

type Operation interface {
	setVerb(string)
}

type operation struct {
	verb         string                `json:"-"` // This is a secreate variable that gets reset when the operation is added to a pathItem
	tags         []string              `json:"tags,omitempty"`
	summary      string                `json:"summary,omitempty"`
	description  string                `json:"description,omitempty"`
	externalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
	operationID  string                `json:"operationId,omitempty"`
	parameters   []Parameter           `json:"parameters,omitempty"`  // or Reference
	requestBody  RequestBody           `json:"requestBody,omitempty"` // or Reference
	responses    Responses             `json:"responses" builder:"required"`
	callbacks    map[string]Callback   `json:"callbacks,omitempty"` // or Reference
	deprecated   bool                  `json:"deprecated,omitempty"`
	security     []SecurityRequirement `json:"security,omitempty"`
	servers      []Server              `json:"servers,omitempty"`
}

type ExternalDocumentation interface {
}

type externalDocumentation struct {
	description string `json:"description" yaml:"description"`
	uRL         string `json:"url" yaml:"url"` // REQUIRED
}

type RequestBody interface {
}

type requestBody struct {
	description string               `json:"description,omitempty"`
	content     map[string]MediaType `json:"content,omitempty"`
	required    bool                 `json:"required,omitempty"`
}

type MediaType interface {
}

type mediaType struct {
	schema   Schema              `json:"schema,omitempty"` // or Reference
	example  interface{}         `json:"example,omitempty"`
	examples map[string]Example  `json:"examples,omitempty"`
	encoding map[string]Encoding `json:"encoding,omitempty"`
}

type Encoding interface {
}

type encoding struct {
	contentType   string            `json:"contentType" yaml:"contentType"`
	headers       map[string]Header `json:"headers" yaml:"headers"`
	explode       bool              `json:"explode" yaml:"explode"`
	allowReserved bool              `json:"allowReserved" yaml:"allowReserved"`
}

type Responses interface {
}

type responses struct {
	defaultValue Response            `json:"default,omitempty" yaml:"default,omitempty"` // or Reference
	responses    map[string]Response `json:"-" yaml:"inline" builder:"-"`                // or Reference
}

type Response interface {
}

type response struct {
	description string               `json:"description" builder:"required"`
	headers     map[string]Header    `json:"headers,omitempty" builder:"-"` // or Reference
	content     map[string]MediaType `json:"content,omitempty" builder:"-"`
	links       map[string]Link      `json:"links,omitempty" builder:"-"` // or Reference
}

type Callback interface {
}

type callback struct {
	uRLs map[string]PathItem
}

type Example interface {
}

type example struct {
	description   string      `json:"description" yaml:"description"`
	value         interface{} `json:"value" yaml:"value"`
	externalValue string      `json:"external_value" yaml:"external_value"`
}

type Link interface {
}

type link struct {
	operationRef string                 `json:"operationRef" yaml:"operationRef"`
	operationID  string                 `json:"operationId" yaml:"operationId"`
	parameters   map[string]interface{} `json:"parameters" yaml:"parameters"`
	requestBody  interface{}            `json:"requestBody" yaml:"requestBody"`
	description  string                 `json:"description" yaml:"description"`
	server       *Server                `json:"server" yaml:"server"`
}

type Tag interface {
}

type tag struct {
	name         string                `json:"name" builder:"required"`
	description  string                `json:"description,omitempty"`
	externalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type Reference interface {
}

type reference struct {
	uRL string `json:"-"` // REQUIRED
}

type Schema interface {
}

type schema struct {
	reference        string                `json:"$ref,omitempty"`
	title            string                `json:"title,omitempty"`
	multipleOf       float64               `json:"multipleOf,omitempty"`
	maximum          float64               `json:"maximum,omitempty"`
	exclusiveMaximum float64               `json:"exclusiveMaximum,omitempty"`
	minimum          float64               `json:"minimum,omitempty"`
	exclusiveMinimum float64               `json:"exclusiveMinimum,omitempty"`
	maxLength        int                   `json:"maxLength,omitempty"`
	minLength        int                   `json:"minLength,omitempty"`
	pattern          string                `json:"pattern,omitempty"`
	maxItems         int                   `json:"maxItems,omitempty"`
	minItems         int                   `json:"minItems,omitempty"`
	uniqueItems      bool                  `json:"uniqueItems,omitempty"`
	maxProperties    int                   `json:"maxProperties,omitempty"`
	minProperties    int                   `json:"minProperties,omitempty"`
	required         []string              `json:"required,omitempty"`
	enum             []interface{}         `json:"enum,omitempty"`
	typ              PrimitiveType         `json:"type,omitempty"`
	allOf            []Schema              `json:"allOf,omitempty"`
	oneOf            []Schema              `json:"oneOf,omitempty"`
	anyOf            []Schema              `json:"anyOf,omitempty"`
	not              Schema                `json:"not,omitempty"`
	items            Schema                `json:"items,omitempty"`
	properties       map[string]Schema     `json:"properties,omitempty"`
	format           string                `json:"format,omitempty"`
	defaultValue     interface{}           `json:"default,omitempty"`
	nullable         bool                  `json:"nullable,omitempty"`
	discriminator    Discriminator         `json:"discriminator,omitempty"`
	readOnly         bool                  `json:"read_only,omitempty"`
	writeOnly        bool                  `json:"write_only,omitempty"`
	externalDocs     ExternalDocumentation `json:"externalDocs,omitempty"`
	example          interface{}           `json:"example,omitempty"`
	deprecated       bool                  `json:"deprecated,omitempty"`
}

type Discriminator interface {
}

type discriminator struct {
	propertyName string            `json:"property_name" yaml:"property_name"` // REQUIRED
	mapping      map[string]string `json:"mapping" yaml:"mapping"`
}

type SecurityScheme interface {
}

type securityScheme struct {
	typ              string     `json:"type" yaml:"type"` // REQUIRED
	description      string     `json:"description" yaml:"description"`
	name             string     `json:"name" yaml:"name"`     // REQUIRED
	in               string     `json:"in" yaml:"in"`         // REQUIRED
	scheme           string     `json:"scheme" yaml:"scheme"` // REQUIRED
	bearerFormat     string     `json:"bearerFormat" yaml:"bearerFormat"`
	flows            OAuthFlows `json:"flows" yaml:"flows"`                       // REQUIRED
	openIDConnectURL string     `json:"openIdConnectUrl" yaml:"openIdConnectUrl"` // REQUIRED
}

type OAuthFlows interface {
}

type oAuthFlows struct {
	implicit          OAuthFlow `json:"implicit" yaml:"implicit"`
	password          OAuthFlow `json:"password" yaml:"password"`
	clientCredentials OAuthFlow `json:"clientCredentials" yaml:"clientCredentials"`
	authorizationCode OAuthFlow `json:"authorizationCode" yaml:"authorizationCode"`
}

type OAuthFlow interface {
}

type oAuthFlow struct {
	authorizationURL string            `json:"authorizationUrl" yaml:"authorizationUrl"`
	tokenURL         string            `json:"tokenUrl" yaml:"tokenUrl"`
	refreshURL       string            `json:"refreshUrl" yaml:"refreshUrl"`
	scopes           map[string]string `json:"scopes" yaml:"scopes"`
}

type SecurityRequirement interface {
}

type securityRequirement struct {
	schemes map[string][]string
}

type Header interface {
}

type header struct {
	in              Location             `json:"-" builder:"required" default:"InHeader"`
	required        bool                 `json:"required,omitempty"`
	description     string               `json:"description,omitempty"`
	deprecated      bool                 `json:"deprecated,omitempty"`
	allowEmptyValue bool                 `json:"allowEmptyValue,omitempty"`
	explode         bool                 `json:"explode,omitempty"`
	allowReserved   bool                 `json:"allowReserved,omitempty"`
	schema          Schema               `json:"schema,omitempty"`
	example         interface{}          `json:"example,omitempty"`
	examples        map[string]Example   `json:"examples,omitempty"`
	content         map[string]MediaType `json:"content,omitempty"`
}

type Parameter interface {
}

type parameter struct {
	name            string               `json:"name,omitempty" builder:"required"`
	in              Location             `json:"in" builder:"required"`
	required        bool                 `json:"required,omitempty" default:"defaultParameterRequiredFromLocation(in)"`
	description     string               `json:"description,omitempty"`
	deprecated      bool                 `json:"deprecated,omitempty"`
	allowEmptyValue bool                 `json:"allowEmptyValue,omitempty"`
	explode         bool                 `json:"explode,omitempty"`
	allowReserved   bool                 `json:"allowReserved,omitempty"`
	schema          Schema               `json:"schema,omitempty"`
	example         interface{}          `json:"example,omitempty"`
	examples        map[string]Example   `json:"examples,omitempty"`
	content         map[string]MediaType `json:"content,omitempty"`
}