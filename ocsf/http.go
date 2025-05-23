package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// HTTPRequestFields defines the Arrow fields for HTTPRequest.
var HTTPRequestFields = []arrow.Field{
	{Name: "args", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "body_length", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "http_method", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "length", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "referrer", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "user_agent", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "http_headers", Type: arrow.ListOf(HTTPHeaderStruct), Nullable: true},
	{Name: "url", Type: URLStruct, Nullable: true},
	{Name: "x_forwarded_for", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
}

var HTTPRequestStruct = arrow.StructOf(HTTPRequestFields...)
var HTTPRequestClassname = "http_request"

type HTTPRequest struct {
	Args          *string       `json:"args,omitempty" parquet:"args"`
	BodyLength    *int          `json:"body_length,omitempty" parquet:"body_length"`
	HTTPMethod    *string       `json:"http_method,omitempty" parquet:"http_method"`
	Length        *int          `json:"length,omitempty" parquet:"length"`
	Referrer      *string       `json:"referrer,omitempty" parquet:"referrer"`
	UID           *string       `json:"uid,omitempty" parquet:"uid"`
	UserAgent     *string       `json:"user_agent,omitempty" parquet:"user_agent"`
	Version       *string       `json:"version,omitempty" parquet:"version,optional"`
	HTTPHeaders   []*HTTPHeader `json:"http_headers,omitempty" parquet:"http_headers,list,optional"`
	URL           *URL          `json:"url,omitempty" parquet:"url"`
	XForwardedFor []string      `json:"x_forwarded_for,omitempty" parquet:"x_forwarded_for,list,optional"`
}

// HTTPResponseFields defines the Arrow fields for HTTPResponse.
var HTTPResponseFields = []arrow.Field{
	{Name: "body_length", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "code", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "content_type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "latency", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "length", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "message", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "http_headers", Type: arrow.ListOf(HTTPHeaderStruct), Nullable: true},
}

var HTTPResponseStruct = arrow.StructOf(HTTPResponseFields...)
var HTTPResponseClassname = "http_response"

type HTTPResponse struct {
	BodyLength  *int          `json:"body_length,omitempty" parquet:"body_length,optional"`
	Code        int           `json:"code" parquet:"code"`
	ContentType *string       `json:"content_type,omitempty" parquet:"content_type,optional"`
	Latency     *int          `json:"latency,omitempty" parquet:"latency,optional"`
	Length      *int          `json:"length,omitempty" parquet:"length,optional"`
	Message     *string       `json:"message,omitempty" parquet:"message,optional"`
	Status      *string       `json:"status,omitempty" parquet:"status,optional"`
	HTTPHeaders []*HTTPHeader `json:"http_headers,omitempty" parquet:"http_headers,list,optional"`
}

// HTTPHeaderFields defines the Arrow fields for HTTPHeader.
var HTTPHeaderFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "value", Type: arrow.BinaryTypes.String, Nullable: false},
}

var HTTPHeaderStruct = arrow.StructOf(HTTPHeaderFields...)
var HTTPHeaderClassname = "http_header"

type HTTPHeader struct {
	Name  string `json:"name" parquet:"name"`
	Value string `json:"value" parquet:"value"`
}
