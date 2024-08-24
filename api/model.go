package api

type (
	// ParamKind - QueryKind, BodyKind, HeaderKind or PathKind
	ParamKind string

	Api struct {
		Endpoints []*Endpoint `json:"endpoints,omitempty"`
		Topics    []*Topic    `json:"topics,omitempty"`
	}

	// Endpoint - describes an endpoint
	Endpoint struct {
		Name        string   `json:"name,omitempty"`
		Description string   `json:"description,omitempty"`
		Method      string   `json:"method"` // http method
		Path        string   `json:"path"`   // http path
		Request     []*Param `json:"request"`
		Response    *Param   `json:"response"`
	}

	Topic struct {
		Name        string   `json:"name,omitempty"`
		Description string   `json:"description,omitempty"`
		Topic       string   `json:"topic"`
		Key         *Param   `json:"key,omitempty"`
		Value       []*Param `json:"value"`
	}

	Param struct {
		Name        string    `json:"name,omitempty"`
		Kind        ParamKind `json:"kind"`
		Description string    `json:"description,omitempty"`
		Type        string    `json:"type,omitempty"`
		Array       bool      `json:"array,omitempty"`
		Map         bool      `json:"map,omitempty"`
		Pointer     bool      `json:"pointer,omitempty"`
		Format      string    `json:"format,omitempty"` // contentType
		Payload     *Payload  `json:"payload,omitempty"`
	}

	Payload struct {
		Name   string   `json:"name"`
		Fields []*Field `json:"fields"`
	}

	Field struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
)

var (
	PathKind   = ParamKind("path")
	BodyKind   = ParamKind("body")
	QueryKind  = ParamKind("query")
	HeaderKind = ParamKind("header")
)
