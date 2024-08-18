package service

import "github.com/Meduzz/dsl/api"

type (
	// ServiceKind - Quickapi, Gin or RPC
	ServiceKind string
	// ConfigKind - Argument or Environment
	ConfigKind string

	// Service - describes a service
	Service struct {
		Name        string      `json:"name"`
		Description string      `json:"description,omitempty"`
		Kind        ServiceKind `json:"kind"`
		Ports       []*Port     `json:"ports,omitempty"`
		Params      []*Config   `json:"params,omitempty"`
		Image       string      `json:"image"`
		Command     string      `json:"command"`
		Volumes     []string    `json:"volumes,omitempty"`
		Api         *api.Api    `json:"api,omitempty"`
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
)

var (
	Quickapi = ServiceKind("quickapi")
	Gin      = ServiceKind("gin")
	RPC      = ServiceKind("rpc")

	Argument    = ConfigKind("arg")
	Environment = ConfigKind("env")
)
