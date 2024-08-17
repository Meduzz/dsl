package service

type (
	// ServiceKind - Quickapi, Gin or RPC
	ServiceKind string
	// ConfigKind - Argument or Environment
	ConfigKind string
	// ParamKind - QueryKind, BodyKind, HeaderKind or PathKind
	ParamKind string

	// Service - describes a service
	Service struct {
		Name        string      `json:"name"`
		Description string      `json:"description,omitempty"`
		Kind        ServiceKind `json:"kind"`
		Endpoints   []*Endpoint `json:"endpoints"`
		Ports       []*Port     `json:"ports,omitempty"`
		Params      []*Config   `json:"params,omitempty"`
		Deploy      *Deploy     `json:"deploy,omitempty"`
	}

	// Endpoint - describes an endpoint, mostly but also public methods..ish
	Endpoint struct {
		Name        string   `json:"name,omitempty"`
		Description string   `json:"description,omitempty"`
		Method      string   `json:"method"` // http method
		Path        string   `json:"path"`   // http path
		Arguments   []*Param `json:"arguments,omitempty"`
		Returns     *Param   `json:"returns,omitempty"`
	}

	Param struct {
		Name        string    `json:"name"`
		Kind        ParamKind `json:"kind"`
		Description string    `json:"description,omitempty"`
		Type        string    `json:"type,omitempty"`
		Array       bool      `json:"array,omitempty"`
		Map         bool      `json:"map,omitempty"`
		Pointer     bool      `json:"pointer,omitempty"`
		Format      string    `json:"format,omitempty"` // contentType
	}

	// Port - describes the ports the app is listening to
	Port struct {
		Protocol    string `json:"protocol"` // tcp/udp
		Port        int    `json:"port"`
		Description string `json:"description,omitempty"`
	}

	// Config - any type of runtime param (except config) that the service expects (args or envs)
	Config struct {
		Name        string     `json:"name"`
		Kind        ConfigKind `json:"kind"` // arg|env
		Description string     `json:"description,omitempty"`
		Default     string     `json:"default,omitempty"`
	}

	// Event - event mapping for dapr style app
	Event struct {
		Topic string `json:"topic"`
		Path  string `json:"path"`
	}

	Deploy struct {
		Image    string     `json:"image"`
		Command  string     `json:"command"`
		PortMaps []*PortMap `json:"portMap,omitempty"`
		Volumes  []*Volume  `json:"volumes,omitempty"`
		Networks []string   `json:"networks,omitempty"`
	}

	PortMap struct {
		Protocol  string `json:"protocol"`
		Container int    `json:"container"`
		Host      int    `json:"host"`
	}

	Volume struct {
		Container string `json:"container"`
		Host      string `json:"host"`
	}
)

var (
	Quickapi = ServiceKind("quickapi")
	Gin      = ServiceKind("gin")
	RPC      = ServiceKind("rpc")

	Argument    = ConfigKind("arg")
	Environment = ConfigKind("env")

	PathKind   = ParamKind("path")
	BodyKind   = ParamKind("body")
	QueryKind  = ParamKind("query")
	HeaderKind = ParamKind("header")
)
