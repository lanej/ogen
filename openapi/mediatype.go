package openapi

import (
	"encoding/json"

	"github.com/lanej/ogen/jsonschema"
	"github.com/lanej/ogen/location"
)

// MediaType is Media Type Object.
type MediaType struct {
	Schema   *jsonschema.Schema
	Example  json.RawMessage
	Examples map[string]*Example
	Encoding map[string]*Encoding

	XOgenJSONStreaming bool
	XOgenRawResponse   bool

	location.Pointer `json:"-" yaml:"-"`
}

// Encoding is Encoding Type Object.
type Encoding struct {
	ContentType   string
	Headers       map[string]*Header
	Style         ParameterStyle
	Explode       bool
	AllowReserved bool

	location.Pointer `json:"-" yaml:"-"`
}
