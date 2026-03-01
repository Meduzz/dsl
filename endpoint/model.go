package endpoint

import "github.com/Meduzz/helper/meta/schema"

type (
	// ParamKind - QueryKind, BodyKind, HeaderKind or PathKind
	ParamKind string

	// RouteKind - GinKind, RpcKind
	RouteKind string

	// Endpoint - describes an endpoint
	Endpoint struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		Route       *Route `json:"route"`
		Permission  string `json:"permission,omitempty"` // TODO is one permission enough?
		// TODO some kind of description of input and output
	}

	Route struct {
		Kind          RouteKind `json:"kind"`
		Method        string    `json:"method,omitempty"` // http method
		Path          string    `json:"path,omitempty"`   // http path
		Topic         string    `json:"topic,omitempty"`
		ConsumerGroup string    `json:"group,omitempty"`
		In            []*Param  `json:"in,omitempty"`
		Out           []*Param  `json:"out,omitempty"`
	}

	Param struct {
		Kind        ParamKind      `json:"kind"`
		Name        string         `json:"name,omitempty"`
		Schema      *schema.Schema `json:"schema,omitempty"`
		ContentType string         `json:"contentType,omitempty"`
	}

	EndpointBuilder interface {
		SetDescription(string)
		SetPermission(string)
		Topic(string, string)
		GET(string)
		POST(string)
		PUT(string)
		DELETE(string)
		PATCH(string)
		RequestHeader(name string)
		Query(name string)
		QueryMap(name string)
		Path(name string)
		RequestBody(contentType string, payload any)
		ResponseHeader(name string)
		ResponseBody(contentType string, payload any)
	}

	endpointBuilder struct {
		endpoint *Endpoint
	}
)

var (
	_ EndpointBuilder = &endpointBuilder{}

	HttpKind = RouteKind("http")
	RpcKind  = RouteKind("rpc")

	HeaderKind = ParamKind("header")
	QueryKind  = ParamKind("query")
	PathKind   = ParamKind("path")
	BodyKind   = ParamKind("body")

	HTML = "text/html"
	JSON = "application/json"
	TEXT = "text/plain"
	MD   = "text/markdown"
)
